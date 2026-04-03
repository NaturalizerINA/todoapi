package models

import "time"

type Subtask struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	NoteID      int       `gorm:"column:note_id;not null" json:"note_id"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	IsCompleted bool      `gorm:"column:is_completed;default:false" json:"is_completed"`
	DateCreated time.Time `gorm:"column:date_created;default:CURRENT_TIMESTAMP" json:"date_created"`
	DateUpdated time.Time `gorm:"column:date_updated;default:CURRENT_TIMESTAMP" json:"date_updated"`
}

func (Subtask) TableName() string {
	return "subtasks"
}
