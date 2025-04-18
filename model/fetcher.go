package model

import (
	"errors"
	"io"
	"net/http"
	"os/exec"
)

func FetchData(mode string, args ...func() string) (string, error) {
	switch mode {
	case "network":
		resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return string(body), nil

	case "python":
		if len(args[0]())+len(args[1]()) < 6 {
			return "", errors.New("требуется два пути к файлам")
		}
		cmd := exec.Command("python3", "scripts/merge.py", args[0](), args[1]())
		out, err := cmd.CombinedOutput()
		return string(out), err

	default:
		return "", errors.New("неизвестный режим")
	}
}
