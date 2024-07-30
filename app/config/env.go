package config

import (
	"os"
)

var envVars *EnvVars

func init() {
	envVars = &EnvVars{
		dbHost:     os.Getenv("DB_HOST"),
		dbPort:     os.Getenv("DB_PORT"),
		dbDatabase: os.Getenv("DB_DATABASE"),
		dbUser:     os.Getenv("DB_USER"),
		dbPassword: os.Getenv("DB_PASSWORD"),
	}
}

type EnvVars struct {
	/* About Database */
	dbHost     string
	dbPort     string
	dbDatabase string
	dbUser     string
	dbPassword string
}
