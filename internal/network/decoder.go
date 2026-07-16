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

	essid := string(b)

	if strings.TrimSpace(essid) == "" {
		return "<hidden>"
	}

	return essid
}
