package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	url2 "net/url"
	"tracker-manager-go/model"
)

type Syncer interface {
	Login(client model.Client, tracker string) error
	Tracker() error
}

func Sync(syncer Syncer, client model.Client, tracker string) error {
	if err := syncer.Login(client, tracker); err != nil {
		return err
	}
	if err := syncer.Tracker(); err != nil {
		return err
	}
	return nil
}

func convertUrl(url string) string {
	if url[:len(url)-1] == "/" {
		url = url[:len(url)-1]
	}
	return url
}

func post(url string, params map[string]any, headers map[string]string, isFormData bool, header *http.Header) (string, error) {
	var bs []byte
	if isFormData {
		form := url2.Values{}
		for k, v := range params {
			form.Set(k, v.(string))
		}
		bs = []byte(form.Encode())
	} else {
		var err error
		bs, err = json.Marshal(params)
		if err != nil {
			return "", err
		}
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(bs))
	if err != nil {
		return "", err
	}
	if isFormData {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	}
	if len(headers) != 0 {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
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
	if response.StatusCode == http.StatusOK {
		// 判断是否需要header
		if header != nil {
			*header = response.Header
		}
		bs1, err := io.ReadAll(body)
		if err != nil {
			return "", err
		}
		return string(bs1), nil
	}
	return "", errors.New("请求失败")
}
