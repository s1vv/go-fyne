package handlers

import (
	"errors"
)

func FetchData(mode string, args ...func() string) (string, error) {
	switch mode {
	case "network":
		return jsonRequest()
	case "python":
		return mergeTxtFiles(args...)
	default:
		return "", errors.New("неизвестный режим")
	}
}
