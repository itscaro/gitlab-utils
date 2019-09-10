package utils

import "regexp"

func SanitizeString(str string) string {
	re := regexp.MustCompile("(?i)[^a-z0-9]")
	return re.ReplaceAllString(str, "-")
}
