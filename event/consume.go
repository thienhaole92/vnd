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

	err := delegate(log, ctx, &schema)
	if err != nil {
		log.Errorw("fail to consume message", "error", err)
	}

	return nil
}
