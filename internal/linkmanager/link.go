package linkmanager

import (
	"time"
)

type Link struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Description  string
	OriginalLink string     `gorm:"not null"`
	OwnerID      string     `gorm:"not null"`
	CreatedAt    time.Time  `gorm:"not null"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime"`
	DeletedAt    *time.Time
}

func (l *Link) BelongsTo(ownerID string) bool {
	return l.OwnerID == ownerID
}
