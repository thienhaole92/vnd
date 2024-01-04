package rest

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RestLogFieldExtractor(c echo.Context) []zapcore.Field {
	if req := c.Get(RequestObjectContextKey); req != nil {
		reqObjectString := ""
		if b, err := json.Marshal(req); err != nil {
			reqObjectString = fmt.Sprintf("fail to parse reqObject as string: %+v", err)
		} else {
			reqObjectString = string(b)
		}

		return []zapcore.Field{zap.String("requestObject", reqObjectString)}
	}

	return nil
}
