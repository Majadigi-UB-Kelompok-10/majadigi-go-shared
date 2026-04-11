package util

import "strings"

func MaskDBSensitiveData(msg string) string {
	if strings.Contains(msg, "postgres://") {
		msg = "database connection error (credentials masked)"
	}
	if strings.Contains(msg, "cloudinary://") {
		msg = "cloudinary error (credentials masked)"
	}
	return msg
}
