package settings

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type DBSetting struct {
	DBPort     string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	SSLMode    string
}

func NewDBSetting() (*DBSetting, error) {
	s := &DBSetting{}
	var err error
	err = godotenv.Load("configs/conf.env")
	if err != nil {
		return nil, err
	}
	s.DBUser, err = GetEnvDefault("ID", "user")
	if err != nil {
		return s, err
	}
	s.DBPassword, err = GetEnvDefault("PASS", "secret")
	if err != nil {
		return s, err
	}
	s.DBPort, err = GetEnvDefault("PORT", "5432")
	if err != nil {
		return s, err
	}
	s.DBName, err = GetEnvDefault("DB", "postgres")
	if err != nil {
		return s, err
	}
	s.DBHost, err = GetEnvDefault("SERVICE_HOSTNAME", "localhost")
	if err != nil {
		return s, err
	}
	s.SSLMode, err = GetEnvDefault("SSL", "false")
	if err != nil {
		return s, err
	}
	return s, nil
}

func GetEnvDefault(key, defaultValue string) (string, error) {
	value := os.Getenv(key)
	if key == "" {
		if defaultValue == "" {
			return "", errors.New("environment variable isn't set")
		}
		return defaultValue, nil
	}

	return value, nil

}
