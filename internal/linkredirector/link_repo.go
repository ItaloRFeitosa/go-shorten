package linkredirector

import (
	"context"

	"gorm.io/gorm"
)

type LinkRepository interface {
	Save(ctx context.Context, link *Link) error
	Get(ctx context.Context, id uint) (*Link, error)
	Delete(ctx context.Context, id uint) error
}

type linkRepo struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) *linkRepo {
	return &linkRepo{db}
}

func (r *linkRepo) Save(ctx context.Context, link *Link) error {
	return r.db.WithContext(ctx).Save(link).Error
}

func (r *linkRepo) Get(ctx context.Context, id uint) (*Link, error) {
	link := new(Link)

	err := r.db.WithContext(ctx).First(link, "id = ?", id).Error

	return link, err
}

func (r *linkRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(Link{}).Delete("id = ?", id).Error
}
