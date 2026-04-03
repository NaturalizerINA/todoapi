package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// CustomTime helps us format JSON with exactly 3 digits
type CustomTime struct {
	time.Time
}

// MarshalJSON returns JSON string with 3-digit precision
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	// Use RFC3339 pattern but force 3 decimals (milliseconds)
	formatted := fmt.Sprintf("\"%s\"", ct.Time.Format("2006-01-02T15:04:05.000Z07:00"))
	return []byte(formatted), nil
}

// UnmarshalJSON handles incoming JSON strings
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		ct.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// Value implements the driver.Valuer interface for GORM
func (ct CustomTime) Value() (driver.Value, error) {
	if ct.Time.IsZero() {
		return nil, nil
	}
	return ct.Time, nil
}

// Scan implements the sql.Scanner interface for GORM
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	if t, ok := value.(time.Time); ok {
		ct.Time = t
		return nil
	}
	return fmt.Errorf("can't scan %T into CustomTime", value)
}

// MasterNote represents the existing master_notes table
type MasterNote struct {
	ID          int        `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserID      uuid.UUID  `gorm:"column:user_id;type:uuid" json:"user_id"`
	Name        string     `gorm:"column:name" json:"name"`
	Status      string     `gorm:"column:status" json:"status"`
	DateCreated CustomTime `gorm:"column:date_created" json:"date_created"`
	DateUpdated CustomTime `gorm:"column:date_updated" json:"date_updated"`
	Subtasks    []Subtask  `gorm:"foreignKey:NoteID;constraint:OnDelete:CASCADE" json:"subtasks"`
}

// TableName tells GORM to use the exactly named table `master_notes`
func (MasterNote) TableName() string {
	return "master_notes"
}
