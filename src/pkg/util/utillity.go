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

func ResponseStatusCode(code int) string {
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

func DRStatus(name string) string {
	var label string
	switch name {
	case "ENROUTE":
		label = "The message is enroute."
	case "DELIVRD":
		label = "The message was successfully delivered."
	case "EXPIRED":
		label = "The SMSC was unable to deliver the message in a specified amount of time. For instance when the phone was turned off."
	case "DELETED":
		label = "The message was deleted."
	case "UNDELIV":
		label = "The SMS was unable to deliver the message. For instance, when the number does not exist."
	case "ACCEPTD":
		label = "The SMS was accepted and will be send."
	case "UNKNOWN":
		label = "Unknown error occurred."
	case "REJECTD":
		label = "The message was rejected. The provider could have blocked phone numbers in this range."
	case "SKIPPED":
		label = "The message was skipped."
	}
	return label
}

func FilterMessage(message string) string {
	i := strings.Index(message, " ")
	if i > -1 {
		keyword := message[i+1:]
		return keyword
	} else {
		return message
	}
}

func KeywordDefine(message string) string {
	var subkey string

	if strings.Contains(strings.ToUpper(message), "REG KEREN") {
		msg := strings.Split(message, " ")
		index := msg[1]
		subkey = index[5:]
	}

	if strings.Contains(strings.ToUpper(message), "REG GM") {
		msg := strings.Split(message, " ")
		index := msg[1]
		subkey = index[2:]
	}

	return subkey
}

func FilterReg(message string) bool {
	index := strings.Split(strings.ToUpper(message), " ")
	if index[0] == "REG" && (strings.Contains(strings.ToUpper(message), "REG KEREN") || strings.Contains(strings.ToUpper(message), "REG GM")) {
		return true
	}
	return false
}

func FilterUnreg(message string) bool {
	index := strings.Split(strings.ToUpper(message), " ")
	if index[0] == "UNREG" && (strings.Contains(strings.ToUpper(message), "UNREG KEREN") || strings.Contains(strings.ToUpper(message), "UNREG GM")) {
		return true
	}
	return false
}
