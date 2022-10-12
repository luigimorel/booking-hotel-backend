package middleware

import "strings"

func ValidateEmail(email string) bool {
	return !strings.Contains(email, "@")
}
