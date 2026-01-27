package utils

import "os"

func GetEnvOrDefault(key string, otherwise string) string {
	if val, exists := os.LookupEnv(key); exists != false {
		return val
	}
	return otherwise
}

func SetEnvIfNotExists(key, value string) error {
	if _, exists := os.LookupEnv(key); exists != false {
		return nil
	}
	return os.Setenv(key, value)
}
