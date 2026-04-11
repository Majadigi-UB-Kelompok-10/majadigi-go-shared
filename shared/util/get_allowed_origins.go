package util

import (
	"strings"
)

func GetAllowedOrigins(origins string) []string {
	if origins == "" {
		origins = "http://localhost:3000,http://localhost:3001,http://localhost:5173"
	}

	var result []string

	for _, origin := range strings.Split(origins, ",") {
		result = append(result, strings.TrimSpace(origin))
	}

	return result
}
