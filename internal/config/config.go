package config

import "os"

type Config struct {
	Addr      string
	HTTPAddr  string
	HTTPSHost string
	CertFile  string
	KeyFile   string
	DSN       string
}

func New() Config {
	return Config{
		Addr:      env("HTTPS_ADDR", ":8443"),
		HTTPAddr:  env("HTTP_ADDR", ":8080"),
		HTTPSHost: env("HTTPS_HOST", "localhost:8443"),
		CertFile:  env("TLS_CERT_FILE", "certs/server.crt"),
		KeyFile:   env("TLS_KEY_FILE", "certs/server.key"),
		DSN:       env("DATABASE_DSN", "postgres://postgres:postgres@localhost:5432/study_security?sslmode=disable"),
	}
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}