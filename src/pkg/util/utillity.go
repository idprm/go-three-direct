package util

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateTransactionId() string {
	id := uuid.New()
	return id.String()
}

func Day(date time.Time) int {
	diff := date.Sub(time.Now())
	return int(diff.Hours() / 24)
}

func TrimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func TrimByteToString(b []byte) string {
	str := string(b)
	return strings.Join(strings.Fields(str), " ")
}
