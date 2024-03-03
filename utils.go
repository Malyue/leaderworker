package leaderworker

import "strings"

// TrimPrefixes("/tmp/file","/tmp") => "/file"
func TrimPrefixes(s string, prefixes ...string) string {
	originLen := len(s)
	for i := range prefixes {
		trimmed := strings.TrimPrefix(s, prefixes[i])
		if len(trimmed) != originLen {
			return trimmed
		}
	}
	return s
}
