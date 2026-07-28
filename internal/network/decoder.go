package network

import (
	"encoding/hex"
	"strings"
)

func DecodeESSID(s string) string {

	s = strings.TrimSpace(s)

	if s == "" || s == "<MISSING>" {
		return "<hidden>"
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return s
	}

	essid := strings.TrimSpace(string(b))
	if essid == "" {
		return "<hidden>"
	}

	return essid
}
