package validation

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func ValidateEmail(email string) bool {
	return emailRegex.MatchString(strings.ToLower(email))
}

func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}
	// Add more password requirements as needed
	return true, ""
}

func ValidateUsername(username string) (bool, string) {
	if len(username) < 3 {
		return false, "Username must be at least 3 characters long"
	}
	if len(username) > 30 {
		return false, "Username must be less than 30 characters"
	}
	return true, ""
}
