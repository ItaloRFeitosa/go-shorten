package linkmanager

import (
	"context"
)

type LinkUseCase interface {
	Create(context.Context, CreateLinkInput) error
	UpdateInfo(context.Context, UpdateLinkInfoInput) error
	ChangeOriginalLink(ctx context.Context, input ChangeOriginalLinkInput) error
	Delete(context.Context, DeleteLinkInput) error
	Find(context.Context, GetAllLinksQuery) ([]LinkDTO, error)
}

type linkUseCase struct {
	repo     LinkRepository
	producer LinkProducer
}

func NewLinkUseCase(repo LinkRepository) *linkUseCase {
	return &linkUseCase{repo, nil}
}

func (u *linkUseCase) Create(ctx context.Context, input CreateLinkInput) error {
	link := new(Link)

	link.Name = input.Name
	link.Description = input.Description
	link.OriginalLink = input.OriginalLink
	link.OwnerID = input.OwnerID

	if err := u.repo.Save(ctx, link); err != nil {
		return err
	}

	event := LinkCreated.WithLinkID(link.ID).
		WithOriginalLink(link.OriginalLink).
		WithCreatedAt(link.CreatedAt).
		RaiseEvent()

	return u.producer.Send(ctx, event)
}

func (u *linkUseCase) UpdateInfo(ctx context.Context, input UpdateLinkInfoInput) error {
	link, err := u.getLink(ctx, input.ID, input.OwnerID)
	if err != nil {
		return err
	}

	link.Name = input.Name

	if input.Description != "" {
		link.Description = input.Description
	}

	return u.repo.Save(ctx, link)
}

func (u *linkUseCase) ChangeOriginalLink(ctx context.Context, input ChangeOriginalLinkInput) error {
	link, err := u.getLink(ctx, input.ID, input.OwnerID)
	if err != nil {
		return err
	}

	link.OriginalLink = input.OriginalLink

	return u.repo.Save(ctx, link)
}

func (u *linkUseCase) Delete(ctx context.Context, input DeleteLinkInput) error {
	link, err := u.getLink(ctx, input.ID, input.OwnerID)
	if err != nil {
		return err
	}

	return u.repo.Delete(ctx, link.ID)
}

func (u *linkUseCase) getLink(ctx context.Context, id uint, ownerID string) (*Link, error) {
	link, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if !link.BelongsTo(ownerID) {
		return nil, ErrWrongOwnerID
	}

	return link, nil
}

func (u *linkUseCase) Find(ctx context.Context, query GetAllLinksQuery) ([]LinkDTO, error) {
	return u.repo.GetAll(ctx, query)
}
