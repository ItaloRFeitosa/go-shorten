package linkmanager

import (
	"context"

	"github.com/italorfeitosa/go-shorten/pkg/hashid"
)

type LinkUseCaseHashIDDecorator struct {
	LinkUseCase
}

func NewLinkUseCaseHashIDDecorator(u LinkUseCase) *LinkUseCaseHashIDDecorator {
	return &LinkUseCaseHashIDDecorator{u}
}

func (u *LinkUseCaseHashIDDecorator) UpdateInfo(ctx context.Context, input UpdateLinkInfoInput) error {
	input.ID = input.Hash.ID()

	return u.LinkUseCase.UpdateInfo(ctx, input)
}

func (u *LinkUseCaseHashIDDecorator) ChangeOriginalLink(ctx context.Context, input ChangeOriginalLinkInput) error {
	input.ID = input.Hash.ID()

	return u.LinkUseCase.ChangeOriginalLink(ctx, input)
}

func (u *LinkUseCaseHashIDDecorator) Delete(ctx context.Context, input DeleteLinkInput) error {
	input.ID = input.Hash.ID()

	return u.LinkUseCase.Delete(ctx, input)
}

func (d *LinkUseCaseHashIDDecorator) Find(ctx context.Context, query GetAllLinksQuery) ([]LinkDTO, error) {
	links, err := d.LinkUseCase.Find(ctx, query)
	if err != nil {
		return links, err
	}

	if err := composeHashes(links); err != nil {
		return links, err
	}

	return links, nil
}

func composeHashes(links []LinkDTO) error {
	var err error
	for i := range links {
		links[i].Hash, err = hashid.New(links[i].ID)
		if err != nil {
			return err
		}
	}

	return nil
}
