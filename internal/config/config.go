package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT   string
	Env    string
	DB_URL string
}

/*
 @Must prefix: it used to indicate that the variable must be set, otherwise it
 will throw an fatal error and the application will not start
*/

func MustLoadConfig() Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	// if port != "" {
	// 	panic("PORT is required")
	// }
	env := os.Getenv("ENV")
	log.Printf("Server Running in %s mode", env)
	if env == "" {
		panic("ENV is required")
	}
	return Config{
		PORT:   port,
		Env:    env,
		DB_URL: os.Getenv("DATABASE_URL"),
	}
}
