package utils

import "os"

func GetEnvOrDefault(key string, otherwise string) string {
	if val, exists := os.LookupEnv(key); exists != false {
		return val
	}
	return otherwise
}
