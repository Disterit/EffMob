package handler

import (
	"EffMob/models"
	"errors"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrNoResponse = errors.New("no response")
)

//func GetSongInfo(externalApiUrl, group, song string) (*models.SongInfo, error) {
//	encodedGroup := url.QueryEscape(group)
//	encodedSong := url.QueryEscape(song)
//
//	resp, err := http.Get(fmt.Sprintf("%s/info?group=%s&song=%s", externalApiUrl, encodedGroup, encodedSong))
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		if resp.StatusCode == http.StatusBadRequest {
//			return nil, ErrBadRequest
//		} else {
//			return nil, ErrNoResponse
//		}
//	}
//
//	var songInfo models.SongInfo
//	if err = json.NewDecoder(resp.Body).Decode(&songInfo); err != nil {
//		return nil, err
//	}
//
//	return &songInfo, nil
//}

func GetSongInfo(externalApiUrl, group, song string) (*models.SongInfo, error) {

	var songInfo models.SongInfo

	songInfo.Text = "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"
	songInfo.Link = "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
	songInfo.ReleaseDate = "16.07.2006"

	return &songInfo, nil
}
