package repository

import (
	"todoapi/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteRepository interface {
	FindAll(userID uuid.UUID) ([]models.MasterNote, error)
	FindByID(id int, userID uuid.UUID) (models.MasterNote, error)
	Create(note *models.MasterNote) error
	Update(note *models.MasterNote) error
	Delete(note *models.MasterNote) error
}

type noteRepositoryImpl struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepositoryImpl{db}
}

func (r *noteRepositoryImpl) FindAll(userID uuid.UUID) ([]models.MasterNote, error) {
	var notes []models.MasterNote
	result := r.db.Preload("Subtasks").Where("user_id = ?", userID).Find(&notes)
	return notes, result.Error
}

func (r *noteRepositoryImpl) FindByID(id int, userID uuid.UUID) (models.MasterNote, error) {
	var note models.MasterNote
	result := r.db.Preload("Subtasks").Where("id = ? AND user_id = ?", id, userID).First(&note)
	return note, result.Error
}

func (r *noteRepositoryImpl) Create(note *models.MasterNote) error {
	return r.db.Create(note).Error
}

func (r *noteRepositoryImpl) Update(note *models.MasterNote) error {
	return r.db.Save(note).Error
}

func (r *noteRepositoryImpl) Delete(note *models.MasterNote) error {
	return r.db.Delete(note).Error
}
