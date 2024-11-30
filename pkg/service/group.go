package service

import (
	"EffMob/pkg/handler"
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

func (s *GroupService) GetAllLibrary() (handler.OutputLibrary, error) {
	return s.Group.GetAllLibrary()
}
