package util

import (
	"io"
	"net/http"
)

func Get(url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	body := response.Body
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(body)
	bytes, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
