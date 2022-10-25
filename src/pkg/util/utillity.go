package util

import (
	"crypto/md5"
	"crypto/rc4"
	"encoding/base64"
	"encoding/hex"
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

func Key(s string) []byte {
	data := md5.Sum([]byte(s))
	secretKey := hex.EncodeToString(data[:])
	return []byte(secretKey)
}

func Cipher(secretKey string) *rc4.Cipher {
	key := Key(secretKey)
	hex, _ := hex.DecodeString(string(key))
	cipher, _ := rc4.NewCipher([]byte(hex))
	return cipher
}

func Encrypt(secretKey string, v string) []byte {
	cipher := Cipher(secretKey)
	text := []byte(v)
	cipher.XORKeyStream(text, text)
	encode := base64.StdEncoding.EncodeToString(text)
	return []byte(encode)
}

func Decrypt(secretKey string, v string) []byte {
	decode, _ := base64.StdEncoding.DecodeString(v)
	cipher := Cipher(secretKey)
	bytes := []byte(decode)
	cipher.XORKeyStream(bytes, bytes)
	return bytes
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
