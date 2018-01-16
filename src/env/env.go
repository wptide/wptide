package env

import "os"

// GetEnv uses "os.LookupEnv", but provides a fallback.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}