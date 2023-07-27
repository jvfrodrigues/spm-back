package config

import "os"

var config = map[string]string{
	"HTTP_PORT": "8080",
	"API_MODE":  "debug",
}

func GetString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		return config[k]
	}
	return v
}
