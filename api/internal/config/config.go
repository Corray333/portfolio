package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func MustInit(envPath string) {
	if err := godotenv.Load(envPath); err != nil {
		panic("error while loading .env file: " + err.Error())
	}
	SetupLogger()
}

func SetupLogger() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(log)
}
