package event

import (
	"context"

	"github.com/thienhaole92/vnd/logger"
)

func Consume[D any](ctx context.Context, message EventString, name string, delegate Delegate[D]) error {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("completed")
		log.Sync()
	}()

	schema := EventSchema[D]{}
	if err := message.UnpackEvent(&schema); err != nil {
		return err
	}

	return delegate(log, ctx, &schema)
}
