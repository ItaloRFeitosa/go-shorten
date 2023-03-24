package linkmanager

import (
	"context"

	"github.com/italorfeitosa/go-shorten/pkg/event"
)

type LinkProducer interface {
	Send(ctx context.Context, ev event.IntegrationEvent[any]) error
}
