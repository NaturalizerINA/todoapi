package models

import "time"

// MasterNote represents the existing master_notes table
type MasterNote struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Status      string    `gorm:"column:status" json:"status"`
	DateCreated time.Time `gorm:"column:date_created;autoCreateTime" json:"date_created"`
	DateUpdated time.Time `gorm:"column:date_updated;autoUpdateTime" json:"date_updated"`
}

// TableName tells GORM to use the exactly named table `master_notes`
func (MasterNote) TableName() string {
	return "master_notes"
}
