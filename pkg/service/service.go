package service

import "EffMob/pkg/repositroy"

type Song interface {
	CreateSong(groupName, songName string) (int, error)
}

type Verse interface {
}

type Service struct {
	Song
	Verse
}

func NewService(repo *repositroy.Repository) *Service {
	return &Service{
		Song: NewSongService(repo),
	}
}
