package service

import (
	"todoapi/models"
	"todoapi/repository"
)

type NoteService interface {
	GetAllNotes() ([]models.MasterNote, error)
	GetNoteByID(id int) (models.MasterNote, error)
	CreateNote(note *models.MasterNote) error
	UpdateNote(id int, updateData models.MasterNote) (models.MasterNote, error)
	DeleteNote(id int) error
}

type noteServiceImpl struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteServiceImpl{repo}
}

func (s *noteServiceImpl) GetAllNotes() ([]models.MasterNote, error) {
	return s.repo.FindAll()
}

func (s *noteServiceImpl) GetNoteByID(id int) (models.MasterNote, error) {
	return s.repo.FindByID(id)
}

func (s *noteServiceImpl) CreateNote(note *models.MasterNote) error {
	return s.repo.Create(note)
}

func (s *noteServiceImpl) UpdateNote(id int, updateData models.MasterNote) (models.MasterNote, error) {
	note, err := s.repo.FindByID(id)
	if err != nil {
		return note, err
	}

	note.Name = updateData.Name
	note.Status = updateData.Status
	note.DateUpdated = updateData.DateUpdated

	err = s.repo.Update(&note)
	return note, err
}

func (s *noteServiceImpl) DeleteNote(id int) error {
	note, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(&note)
}
