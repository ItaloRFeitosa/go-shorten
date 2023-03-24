package linkmanager

import (
	"time"

	"github.com/italorfeitosa/go-shorten/pkg/hashid"
)

type CreateLinkInput struct {
	Name         string `json:"name" validate:"required,min=2"`
	Description  string `json:"description" validate:"omitempty,min=2,max=512"`
	OriginalLink string `json:"originalLink" validate:"required,url"`
	OwnerID      string `json:"-" validate:"required"`
}

type UpdateLinkInfoInput struct {
	ID          uint           `json:"-"`
	Hash        *hashid.HashID `json:"-" validate:"required"`
	Name        string         `json:"name" validate:"required,min=2"`
	Description string         `json:"description" validate:"omitempty,min=2,max=512"`
	OwnerID     string         `json:"-" validate:"required"`
}

type ChangeOriginalLinkInput struct {
	ID           uint           `json:"-"`
	Hash         *hashid.HashID `json:"-" validate:"required"`
	OriginalLink string         `json:"originalLink" validate:"required,url"`
	OwnerID      string         `json:"-" validate:"required"`
}

type DeleteLinkInput struct {
	ID      uint
	Hash    *hashid.HashID `validate:"required"`
	OwnerID string         `validate:"required"`
}

type LinkDTO struct {
	ID           uint           `json:"-"`
	Hash         *hashid.HashID `json:"hash" gorm:"-"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	OriginalLink string         `json:"originalLink"`
	OwnerID      string         `json:"ownerId"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
}

type GetAllLinksQuery struct {
	ID      uint   `query:"-"`
	OwnerID string `query:"-"`
	Hash    string `query:"hash"`
	Limit   int    `query:"limit"`
	Page    int    `query:"page"`
}

func (q GetAllLinksQuery) Offset() int {
	return (q.Page - 1) * q.Limit
}
