package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseCode(str string) bool {
	if str == "" {
		return false
	}

	for _, v := range str {
		if v != '.' && v != '-' && v != '/' {
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

	if MorseCode(str) {
		result := morse.ToText(str)
		if result == " " {
			return "", errors.New("failed to convert input")
		}
	}

	result := morse.ToMorse(str)
	if result == " " {
		return "", errors.New("failed to convert input")
	}

	return result, nil
}
