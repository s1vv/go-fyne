package handlers

import (
	"io"
	"net/http"
)

func jsonRequest() (string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil

}
