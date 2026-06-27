package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	Dsn                string
	JwtSecret          string
	CorsAllowedOrigins []string
}

func LoadEnv() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:               getEnvOrDefault("PORT", "8080"),
		Dsn:                getEnvOrDefault("DSN", ""),
		JwtSecret:          getEnvOrDefault("JWT_SECRET", "change-me"),
		CorsAllowedOrigins: parseCorsOrigins(getEnvOrDefault("CORS_ALLOWED_ORIGINS", "*")),
	}
}

func getEnvOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func parseCorsOrigins(raw string) []string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return []string{"*"}
	}

	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			origins = append(origins, part)
		}
	}

	if len(origins) == 0 {
		return []string{"*"}
	}

	return origins
}
