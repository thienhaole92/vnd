package event

import (
	"context"

	"github.com/thienhaole92/vnd/logger"
)

type Delegate[D any] func(*logger.Logger, context.Context, *EventSchema[D]) error
