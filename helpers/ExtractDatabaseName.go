package helpers

import "strings"

func ExtractDatabaseName(dsn string) string {
	parts := strings.Split(dsn, "/")
	if len(parts) == 0 {
		return ""
	}

	last := parts[len(parts)-1]
	if idx := strings.Index(last, "?"); idx != -1 {
		return last[:idx]
	}

	return last
}
