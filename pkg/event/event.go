package event

import "time"

type IntegrationEvent[T any] struct {
	ID       string    `json:"id"`
	Entity   string    `json:"entity"`
	EntityID uint      `json:"entityId"`
	Name     string    `json:"name"`
	RaisedAt time.Time `json:"raisedAt"`
	Data     T         `json:"data"`
}
