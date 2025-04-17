package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchData() (string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var pretty map[string]interface{}
	if err := json.Unmarshal(body, &pretty); err != nil {
		return "", err
	}

	return fmt.Sprintf("ID: %v, Title: %v", pretty["id"], pretty["title"]), nil
}
