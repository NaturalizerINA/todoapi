package repository

import (
	"todoapi/models"

	"gorm.io/gorm"
)

type SubtaskRepository interface {
	Create(subtask *models.Subtask) error
	GetByNoteID(noteID int) ([]models.Subtask, error)
	Update(subtask *models.Subtask) error
	Delete(id int) error
	Toggle(id int) error
}

type subtaskRepositoryImpl struct {
	db *gorm.DB
}

func NewSubtaskRepository(db *gorm.DB) SubtaskRepository {
	return &subtaskRepositoryImpl{db}
}

func (r *subtaskRepositoryImpl) Create(subtask *models.Subtask) error {
	return r.db.Create(subtask).Error
}

func (r *subtaskRepositoryImpl) GetByNoteID(noteID int) ([]models.Subtask, error) {
	var subtasks []models.Subtask
	err := r.db.Where("note_id = ?", noteID).Find(&subtasks).Error
	return subtasks, err
}

func (r *subtaskRepositoryImpl) Update(subtask *models.Subtask) error {
	return r.db.Save(subtask).Error
}

func (r *subtaskRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&models.Subtask{}, id).Error
}

func (r *subtaskRepositoryImpl) Toggle(id int) error {
	var subtask models.Subtask
	if err := r.db.First(&subtask, id).Error; err != nil {
		return err
	}
	subtask.IsCompleted = !subtask.IsCompleted
	return r.db.Save(&subtask).Error
}
