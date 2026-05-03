package httpapi

import (
	"regexp"
	"strconv"
	"strings"
)

var positiveIDPattern = regexp.MustCompile(`^[1-9][0-9]*$`)
var emailPattern = regexp.MustCompile(`^[A-Za-z0-9._%+\-]+@[A-Za-z0-9.\-]+\.[A-Za-z]{2,}$`)

func parsePositiveID(raw string) (int64, bool) {
	value := strings.TrimSpace(raw)
	if !positiveIDPattern.MatchString(value) {
		return 0, false
	}

	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, false
	}

	return id, true
}

func normalizeEmail(raw string) (string, bool) {
	value := strings.ToLower(strings.TrimSpace(raw))
	if len(value) < 5 || len(value) > 254 {
		return "", false
	}

	if !emailPattern.MatchString(value) {
		return "", false
	}

	return value, true
}