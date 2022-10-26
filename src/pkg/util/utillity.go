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

func EscapeChar(res []byte) []byte {
	response := string(res)
	r := strings.NewReplacer("&lt;", "<", "&gt;", ">")
	result := r.Replace(response)
	return []byte(result)
}

func ResponseCode(code int) string {
	var label string
	switch code {
	case 0:
		label = "Successful"
	case 52:
		label = "Low Balance"
	case 54:
		label = "Billing Error"
	case 16:
		label = "Invalid Account Name or ESME ID"
	case 53:
		label = "Unknown Subscriber"
	case 19:
		label = "Monthly Quota Limit Reached"
	case 8:
		label = "Account throttled"
	}
	return label
}
