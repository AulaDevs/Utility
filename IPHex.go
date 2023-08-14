package Utility

import (
	"fmt"
	"strings"
	"strconv"
	"crypto/sha1"
	"encoding/hex"
)

func EncodeIP(ip string) string {
	parts := strings.Split(ip, ".")
	encodedParts := make([]string, len(parts))

	for i, part := range parts {
		val, _ := strconv.Atoi(part)
		encodedParts[i] = fmt.Sprintf("%02X", val)
	}

	return "#" + strings.Join(encodedParts, ".")
}

func DecodeIP(ip string) string {
	encodedParts := strings.Split(ip[1:], ".")
	decodedParts := make([]string, len(encodedParts))

	for i, part := range encodedParts {
		val, _ := strconv.ParseInt(part, 16, 0)
		decodedParts[i] = fmt.Sprintf("%d", val)
	}

	return strings.Join(decodedParts, ".")
}

func GetIPColor(ip string) string {
	hash := sha1.Sum([]byte(ip))
	hexDigest := hex.EncodeToString(hash[:])
	return hexDigest[:6]
}