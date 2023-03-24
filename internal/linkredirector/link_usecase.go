package linkredirector

import (
	"context"
)

type LinkUseCase interface {
	Create(ctx context.Context, input LinkDTO) error
	ChangeOriginalLink(ctx context.Context, input LinkDTO) error
	Delete(context.Context, uint) error
	GetOriginalLink(context.Context, uint) (string, error)
}

type linkUseCase struct {
	repo LinkRepository
}

func NewLinkUseCase(repo LinkRepository) *linkUseCase {
	return &linkUseCase{repo}
}

func (u *linkUseCase) Create(ctx context.Context, input LinkDTO) error {
	link := new(Link)

	link.OriginalLink = input.OriginalLink

	return u.repo.Save(ctx, link)
}

func (u *linkUseCase) ChangeOriginalLink(ctx context.Context, input LinkDTO) error {
	link, err := u.repo.Get(ctx, input.ID)
	if err != nil {
		return err
	}

	link.OriginalLink = input.OriginalLink

	return u.repo.Save(ctx, link)
}

func (u *linkUseCase) Delete(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}

func (u *linkUseCase) GetOriginalLink(ctx context.Context, id uint) (string, error) {
	link, err := u.repo.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return link.OriginalLink, nil
}
