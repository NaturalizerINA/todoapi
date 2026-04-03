package repository

import (
	"errors"
	"todoapi/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubtaskRepository interface {
	Create(subtask *models.Subtask, userID uuid.UUID) error
	GetByNoteID(noteID int, userID uuid.UUID) ([]models.Subtask, error)
	Update(subtask *models.Subtask, userID uuid.UUID) error
	Delete(id int, userID uuid.UUID) error
	Toggle(id int, userID uuid.UUID) error
}

type subtaskRepositoryImpl struct {
	db *gorm.DB
}

func NewSubtaskRepository(db *gorm.DB) SubtaskRepository {
	return &subtaskRepositoryImpl{db}
}

func (r *subtaskRepositoryImpl) Create(subtask *models.Subtask, userID uuid.UUID) error {
	// First verify note ownership
	var count int64
	r.db.Model(&models.MasterNote{}).Where("id = ? AND user_id = ?", subtask.NoteID, userID).Count(&count)
	if count == 0 {
		return errors.New("unauthorized or note not found")
	}
	return r.db.Create(subtask).Error
}

func (r *subtaskRepositoryImpl) GetByNoteID(noteID int, userID uuid.UUID) ([]models.Subtask, error) {
	var subtasks []models.Subtask
	err := r.db.Joins("JOIN master_notes ON master_notes.id = subtasks.note_id").
		Where("subtasks.note_id = ? AND master_notes.user_id = ?", noteID, userID).
		Find(&subtasks).Error
	return subtasks, err
}

func (r *subtaskRepositoryImpl) Update(subtask *models.Subtask, userID uuid.UUID) error {
	// Verify ownership join
	var existing models.Subtask
	err := r.db.Joins("JOIN master_notes ON master_notes.id = subtasks.note_id").
		Where("subtasks.id = ? AND master_notes.user_id = ?", subtask.ID, userID).
		First(&existing).Error
	if err != nil {
		return errors.New("unauthorized or subtask not found")
	}
	return r.db.Save(subtask).Error
}

func (r *subtaskRepositoryImpl) Delete(id int, userID uuid.UUID) error {
	var existing models.Subtask
	err := r.db.Joins("JOIN master_notes ON master_notes.id = subtasks.note_id").
		Where("subtasks.id = ? AND master_notes.user_id = ?", id, userID).
		First(&existing).Error
	if err != nil {
		return errors.New("unauthorized or subtask not found")
	}
	return r.db.Delete(&models.Subtask{}, id).Error
}

func (r *subtaskRepositoryImpl) Toggle(id int, userID uuid.UUID) error {
	var subtask models.Subtask
	err := r.db.Joins("JOIN master_notes ON master_notes.id = subtasks.note_id").
		Where("subtasks.id = ? AND master_notes.user_id = ?", id, userID).
		First(&subtask).Error
	if err != nil {
		return errors.New("unauthorized or subtask not found")
	}
	
	subtask.IsCompleted = !subtask.IsCompleted
	return r.db.Save(&subtask).Error
}
