package linkmanager

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type LinkRepository interface {
	Save(ctx context.Context, link *Link) error
	Get(ctx context.Context, id uint) (*Link, error)
	GetAll(ctx context.Context, query GetAllLinksQuery) ([]LinkDTO, error)
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

	err := r.db.WithContext(ctx).First(link, "id = ? and deleted_at is null", id).Error

	return link, err
}

func (r *linkRepo) GetAll(ctx context.Context, q GetAllLinksQuery) ([]LinkDTO, error) {
	var links []LinkDTO

	query := r.db.WithContext(ctx).Model(Link{}).Where("owner_id = ? and deleted_at is null", q.OwnerID)

	if q.ID != 0 {
		query.Where("id = ?", q.ID)
	}

	err := query.Offset(q.Offset()).Limit(q.Limit).Find(&links).Error

	return links, err
}

func (r *linkRepo) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Model(Link{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}
