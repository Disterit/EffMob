package service

import "EffMob/pkg/repositroy"

type SongService struct {
	repo repositroy.Song
}

func NewSongService(repo repositroy.Song) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(groupName, songName string) (int, error) {
	return s.repo.CreateSong(groupName, songName)
}
