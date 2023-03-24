package linkmanager

import (
	"time"

	"github.com/google/uuid"
	"github.com/italorfeitosa/go-shorten/pkg/event"
)

var LinkEntity = "Link"

type linkCreated struct {
	EventName    string    `json:"-"`
	LinkID       uint      `json:"linkId"`
	OriginalLink string    `json:"originalLink"`
	CreatedAt    time.Time `json:"createdAt"`
}

var (
	LinkCreated         linkCreated
	LinkDeleted         linkDeleted
	OriginalLinkChanged originalLinkChanged
)

func (l linkCreated) WithLinkID(id uint) linkCreated {
	l.LinkID = id

	return l
}

func (l linkCreated) WithOriginalLink(original string) linkCreated {
	l.OriginalLink = original

	return l
}

func (l linkCreated) WithCreatedAt(at time.Time) linkCreated {
	l.CreatedAt = at

	return l
}

func (l linkCreated) RaiseEvent() event.IntegrationEvent[any] {
	var e event.IntegrationEvent[any]

	e.ID = uuid.NewString()
	e.Entity = LinkEntity
	e.EntityID = l.LinkID
	e.Name = "LinkCreated"
	e.RaisedAt = time.Now()
	e.Data = l

	return e
}

type linkDeleted struct {
	LinkID    uint      `json:"linkId"`
	DeletedAt time.Time `json:"deletedAt"`
}

func (l linkDeleted) WithLinkID(id uint) linkDeleted {
	l.LinkID = id

	return l
}

func (l linkDeleted) WithDeletedAt(at time.Time) linkDeleted {
	l.DeletedAt = at

	return l
}

func (l linkDeleted) RaiseEvent() event.IntegrationEvent[any] {
	var e event.IntegrationEvent[any]

	e.ID = uuid.NewString()
	e.Entity = LinkEntity
	e.EntityID = l.LinkID
	e.Name = "LinkDeleted"
	e.RaisedAt = time.Now()
	e.Data = l

	return e
}

type originalLinkChanged struct {
	LinkID       uint   `json:"linkId"`
	OriginalLink string `json:"originalLink"`
}

func (l originalLinkChanged) WithLinkID(id uint) originalLinkChanged {
	l.LinkID = id

	return l
}

func (l originalLinkChanged) WithOriginalLink(original string) originalLinkChanged {
	l.OriginalLink = original

	return l
}

func (l originalLinkChanged) RaiseEvent() event.IntegrationEvent[any] {
	var e event.IntegrationEvent[any]

	e.ID = uuid.NewString()
	e.Entity = LinkEntity
	e.EntityID = l.LinkID
	e.Name = "OriginalLinkChanged"
	e.RaisedAt = time.Now()
	e.Data = l

	return e
}
