package handler

import (
	"EffMob/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrNoResponse = errors.New("no response")
)

func GetSongInfo(externalApiUrl, group, song string) (*models.SongInfo, error) {
	encodedGroup := url.QueryEscape(group)
	encodedSong := url.QueryEscape(song)

	resp, err := http.Get(fmt.Sprintf("%s/info?group=%s&song=%s", externalApiUrl, encodedGroup, encodedSong))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			return nil, ErrBadRequest
		} else {
			return nil, ErrNoResponse
		}
	}

	var songInfo models.SongInfo
	if err = json.NewDecoder(resp.Body).Decode(&songInfo); err != nil {
		return nil, err
	}

	return &songInfo, nil
}
