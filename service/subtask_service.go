package service

import (
	"todoapi/models"
	"todoapi/repository"

	"github.com/google/uuid"
)

type SubtaskService interface {
	Create(subtask *models.Subtask, userID uuid.UUID) error
	Update(subtask *models.Subtask, userID uuid.UUID) error
	Delete(id int, userID uuid.UUID) error
	Toggle(id int, userID uuid.UUID) error
}

type subtaskServiceImpl struct {
	repo repository.SubtaskRepository
}

func NewSubtaskService(repo repository.SubtaskRepository) SubtaskService {
	return &subtaskServiceImpl{repo}
}

func (s *subtaskServiceImpl) Create(subtask *models.Subtask, userID uuid.UUID) error {
	return s.repo.Create(subtask, userID)
}

func (s *subtaskServiceImpl) Update(subtask *models.Subtask, userID uuid.UUID) error {
	return s.repo.Update(subtask, userID)
}

func (s *subtaskServiceImpl) Delete(id int, userID uuid.UUID) error {
	return s.repo.Delete(id, userID)
}

func (s *subtaskServiceImpl) Toggle(id int, userID uuid.UUID) error {
	return s.repo.Toggle(id, userID)
}
