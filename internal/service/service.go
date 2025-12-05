package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseCode(str string) bool {
	if str == "" {
		return false
	}

	for _, v := range str {
		if v != '.' && v != '-' && v != '/' && v != ' ' && v != '\n' && v != '\r' {
			return false
		}
	}

	if !strings.ContainsAny(str, ".-") {
		return false
	}

	return true
}

func DetectAndConvert(str string) (string, error) {
	if str == "" {
		return "", errors.New("empty line")
	}

	fmt.Printf("[DEBUG] Input: %q\n", str)
	fmt.Printf("[DEBUG] MorseCode(): %v\n", MorseCode(str))

	prefix := "Конвертированный текст: "
	if strings.HasPrefix(str, prefix) {
		str = strings.TrimPrefix(str, prefix)
		str = strings.TrimSpace(str)
		fmt.Printf("[DEBUG] After removing prefix: %q\n", str)
	}

	if MorseCode(str) {
		result := morse.ToText(str)
		fmt.Printf("[DEBUG] morse.ToText(): %q\n", result)
		if result == "" {
			return "", errors.New("failed to convert input")
		}

		return result, nil
	} else {
		result := morse.ToMorse(str)
		fmt.Printf("[DEBUG] morse.ToMorse(): %q\n", result)

		if result == "" {
			return "", errors.New("failed to encode text")
		}

		return result, nil
	}
}
