package service

import (
	"EffMob/logger"
	"EffMob/pkg/repositroy"
	"errors"
	"strconv"
	"strings"
)

type VerseService struct {
	repo repositroy.Verse
}

func NewVerseService(repo repositroy.Verse) *VerseService {
	return &VerseService{repo: repo}
}

func (s *VerseService) GetVerses(songId, verseId, limit int) (map[string]string, error) {
	output := make(map[string]string)

	text, err := s.repo.GetVerses(songId)
	if err != nil {
		logger.Log.Error("error to get song text from repository", err.Error())
		return nil, err
	}

	if text == "" {
		logger.Log.Error("text song is empty")
		return nil, errors.New("text song is empty")
	}

	verses := strings.Split(text, "\n\n")
	if verseId < 1 || verseId > len(verses) {
		logger.Log.Error("not enough id verses in song text")
		return nil, errors.New("not enough id verses in song text")
	}

	startIndex := verseId - 1
	endIndex := startIndex + limit + 1
	if endIndex > len(verses) {
		endIndex = len(verses)
	}

	for i := startIndex; i < endIndex; i++ {
		output["verse"+strconv.Itoa(i+1)] = verses[i]
	}

	return output, nil
}
