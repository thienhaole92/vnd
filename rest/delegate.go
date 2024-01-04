package rest

import (
	"github.com/thienhaole92/vnd/context"
	"github.com/thienhaole92/vnd/logger"
)

type Delegate[REQ any] func(*logger.Logger, context.Context, *REQ) (*Result, error)
