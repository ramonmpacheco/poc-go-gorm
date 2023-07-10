package model

import "time"

type Ingredient struct {
	ID        string
	Name      string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
