package utils

import "strings"

func NormalizeString(input string) string {
    return strings.ReplaceAll(strings.ToLower(input), " ", "_")
}
