package entity

import "time"

type Ingredient struct {
	ID        string     `gorm:"primaryKey;size:50"`
	Name      string     `gorm:"not null;unique;size:50" sql:"index"`
	Desc      string     `gorm:"not null;size:250"`
	CreatedAt *time.Time `gorm:"default:current_timestamp"`
	UpdatedAt *time.Time
}

// Custom table name
func (Ingredient) TableName() string {
	return "ingredients"
}
