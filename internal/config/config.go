package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	Env  string
}

/*
 @Must prefix: it used to indicate that the variable must be set, otherwise it
 will throw an fatal error and the application will not start
*/

func MustLoadConfig() Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port != "" {
		panic("PORT is required")
	}
	env := os.Getenv("ENV")
	if env != "" {
		panic("ENV is required")
	}
	return Config{
		PORT: port,
		Env:  env,
	}
}
