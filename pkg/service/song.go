package service

import (
	"EffMob/models"
	"EffMob/pkg/repositroy"
)

type SongService struct {
	repo repositroy.Song
}

func NewSongService(repo repositroy.Song) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(groupName, songName string) (int, error) {
	return s.repo.CreateSong(groupName, songName)
}

func (s *SongService) GetAllSongs() ([]models.Song, error) {
	return s.repo.GetAllSongs()
}

func (s *SongService) GetSongById(id int) (models.Song, error) {
	return s.repo.GetSongById(id)
}

func (s *SongService) UpdateSong(id int, input models.UpdateSong) error {
	return s.repo.UpdateSong(id, input)
}

func (s *SongService) DeleteSong(id int) error {
	return s.repo.DeleteSong(id)
}
