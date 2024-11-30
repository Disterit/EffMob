package service

import (
	"EffMob/models"
	"EffMob/pkg/repositroy"
)

type GroupService struct {
	repositroy.Group
}

func NewGroupService(group repositroy.Group) *GroupService {
	return &GroupService{Group: group}
}

func (s *GroupService) CreateGroup(groupName string) (int, error) {
	return s.Group.CreateGroup(groupName)
}

func (s *GroupService) GetAllLibrary() (map[string][]models.Song, error) {
	return s.Group.GetAllLibrary()
}

func (s *GroupService) GetAllSongGroupById(id int) (map[string][]models.Song, error) {
	return s.Group.GetAllSongGroupById(id)
}

func (s *GroupService) UpdateGroup(id int, input models.Group) error {
	return s.Group.UpdateGroup(id, input)
}

func (s *GroupService) DeleteGroup(id int) error {
	return s.Group.DeleteGroup(id)
}
