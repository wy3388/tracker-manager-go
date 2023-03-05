package client

import (
	"errors"
	"net/http"
	"strings"
	"tracker-manager-go/model"
)

type QBittorrent struct {
	url     string
	token   string
	tracker string
}

func (q *QBittorrent) Login(client model.Client, tracker string) error {
	url := convertUrl(client.Url)
	q.url = url
	q.tracker = tracker
	url = url + QbittorrentLoginUrl
	params := map[string]any{
		"username": client.Username,
		"password": client.Password,
	}
	var header http.Header
	_, err := post(url, params, map[string]string{}, true, &header)
	if err != nil {
		return err
	}
	cookie := header.Get("set-cookie")
	if cookie != "" {
		sArr := strings.Split(cookie, ";")
		if len(sArr) > 0 {
			q.token = sArr[0]
			return nil
		}
	}
	return errors.New("登录客户端失败")
}

func (q *QBittorrent) Tracker() error {
	url := q.url + QbittorrentSetPreferences
	headers := map[string]string{
		"Cookie": q.token,
	}
	params := map[string]any{
		"json": "{\"add_trackers_enabled\": true, \"add_trackers\": \"" + q.tracker + "\"}",
	}
	_, err := post(url, params, headers, true, nil)
	if err != nil {
		return err
	}
	return nil
}
