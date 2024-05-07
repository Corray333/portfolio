package config

import "github.com/joho/godotenv"

func MustInit(envPath string) {
	if err := godotenv.Load(envPath); err != nil {
		panic("error while loading .env file: " + err.Error())
	}
}
