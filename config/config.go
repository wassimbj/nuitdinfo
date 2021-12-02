package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b) // this is the path where this func is ran which is ./config

	//? envPath = ./config/../.env = server/.env
	envPath := filepath.Join(basepath, "..", ".env")
	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
