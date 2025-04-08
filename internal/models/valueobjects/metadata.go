package valueobjects

import "time"

type Metadata struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMetadata(createdAt, updatedAt time.Time) *Metadata {
	return &Metadata{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
