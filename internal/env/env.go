package env

import "os"

type EnvError struct {
	error
	env string
}

func (e *EnvError) Error() string { return "Env not found: " + e.env }

func GetOrDefault(key string, def string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return def
}

func Get(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if exists {
		return value, nil
	}
	return "", &EnvError{env: key}
}
