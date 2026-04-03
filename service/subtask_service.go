package service

import (
	"todoapi/models"
	"todoapi/repository"
)

type SubtaskService interface {
	Create(subtask *models.Subtask) error
	Update(subtask *models.Subtask) error
	Delete(id int) error
	Toggle(id int) error
}

type subtaskServiceImpl struct {
	repo repository.SubtaskRepository
}

func NewSubtaskService(repo repository.SubtaskRepository) SubtaskService {
	return &subtaskServiceImpl{repo}
}

func (s *subtaskServiceImpl) Create(subtask *models.Subtask) error {
	return s.repo.Create(subtask)
}

func (s *subtaskServiceImpl) Update(subtask *models.Subtask) error {
	return s.repo.Update(subtask)
}

func (s *subtaskServiceImpl) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *subtaskServiceImpl) Toggle(id int) error {
	return s.repo.Toggle(id)
}
