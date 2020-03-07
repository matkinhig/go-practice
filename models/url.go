package models

import "time"

type Url struct {
	ID        uint64
	Url       string
	Status    bool
	CreatedAt time.Time
	DeletedAt time.Time
	UpdatedAt time.Time
}
