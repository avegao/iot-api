package util

import "strconv"

func ParseUint64(value string) uint64 {
	parsed, _ := strconv.ParseUint(value, 0, 64)

	return parsed
}
