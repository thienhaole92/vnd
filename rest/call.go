package rest

import (
	"github.com/thienhaole92/vnd/context"
	"github.com/thienhaole92/vnd/logger"
)

func Call[REQ any](e context.Context, req *REQ, name string, delegate Delegate[REQ]) (*Result, error) {
	log := logger.GetLogger(name)
	defer func() {
		log.Infow("completed")
		log.Sync()
	}()

	requestId := e.RequestId()
	log.With([]interface{}{
		"request_id", requestId,
	}...)

	return delegate(log, e, req)
}
