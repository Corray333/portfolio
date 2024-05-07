package auth

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	AccessTokenLifeTime  = time.Minute * 15
	RefreshTokenLifeTime = time.Hour * 24 * 365
)

var secretKey []byte

// init initializes the secret key from the environment variable
func init() {
	secretKey = []byte(os.Getenv("SECRET_KEY"))
}

func NewMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		slog.Info("auth middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			if err := VerifyToken(r.Header.Get("Authorization")); err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				slog.Error("Unauthorized: " + err.Error())
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

// Hash hashes the password using bcrypt package
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Verify checks if the hashed password is equal to the password using bcrypt package
func Verify(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

// CreateToken creates a new JWT token by the email
func CreateToken(lifeTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": time.Now().Add(lifeTime).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken checks if the JWT is valid
func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

type Credentials struct {
	Email string `json:"email"`
	ID    int    `json:"id,omitempty" db:"user_id"`
}

func ExtractCredentials(tokenString string) (Credentials, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return Credentials{}, err
	}
	credentials := Credentials{
		ID: int(claims["id"].(float64)),
	}
	return credentials, nil
}

type Storage interface {
	CheckAndUpdateRefresh(id int, refresh string) (string, error)
}

func RefreshAccessToken(store Storage, refresh string) (string, string, error) {

	token, err := jwt.Parse(refresh, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", fmt.Errorf("invalid refresh token")
	}

	creds, err := ExtractCredentials(refresh)
	if err != nil {
		return "", "", err
	}

	newRefresh, err := store.CheckAndUpdateRefresh(creds.ID, refresh)
	if err != nil {
		return "", "", err
	}
	newAccess, err := CreateToken(AccessTokenLifeTime)
	if err != nil {
		return "", "", err
	}
	return newAccess, newRefresh, nil

}
