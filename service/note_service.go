package service

import (
	"todoapi/models"
	"todoapi/repository"

	"github.com/google/uuid"
)

type NoteService interface {
	GetAllNotes(userID uuid.UUID) ([]models.MasterNote, error)
	GetNoteByID(id int, userID uuid.UUID) (models.MasterNote, error)
	CreateNote(note *models.MasterNote) error // Already contains UserID in note
	UpdateNote(id int, userID uuid.UUID, updateData models.MasterNote) (models.MasterNote, error)
	DeleteNote(id int, userID uuid.UUID) error
}

type noteServiceImpl struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteServiceImpl{repo}
}

func (s *noteServiceImpl) GetAllNotes(userID uuid.UUID) ([]models.MasterNote, error) {
	return s.repo.FindAll(userID)
}

func (s *noteServiceImpl) GetNoteByID(id int, userID uuid.UUID) (models.MasterNote, error) {
	return s.repo.FindByID(id, userID)
}

func (s *noteServiceImpl) CreateNote(note *models.MasterNote) error {
	return s.repo.Create(note)
}

func (s *noteServiceImpl) UpdateNote(id int, userID uuid.UUID, updateData models.MasterNote) (models.MasterNote, error) {
	note, err := s.repo.FindByID(id, userID)
	if err != nil {
		return note, err
	}

	note.Name = updateData.Name
	note.Status = updateData.Status
	note.DateUpdated = updateData.DateUpdated

	err = s.repo.Update(&note)
	return note, err
}

func (s *noteServiceImpl) DeleteNote(id int, userID uuid.UUID) error {
	note, err := s.repo.FindByID(id, userID)
	if err != nil {
		return err
	}

	return s.repo.Delete(&note)
}
