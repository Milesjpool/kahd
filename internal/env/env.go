package env

import "os"

func GetOrDefault(key string, def string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return def
}
