package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	vnderror "github.com/thienhaole92/vnd/error"
)

func Put[RES any](ctx context.Context, requestId string, url string, heades map[string]string, body any) (*RES, error) {
	var req *http.Request

	if body != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}

		r, err := http.NewRequestWithContext(ctx, http.MethodPut, url, buf)
		if err != nil {
			return nil, err
		}
		req = r
	} else {
		r, err := http.NewRequestWithContext(ctx, http.MethodPut, url, nil)
		if err != nil {
			return nil, err
		}
		req = r
	}

	for k, v := range heades {
		req.Header.Set(k, v)
	}
	req.Header.Set(echo.HeaderContentType, "application/json")
	req.Header.Set(echo.HeaderXRequestID, requestId)

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var httpErr vnderror.Error
		if err := json.Unmarshal(body, &httpErr); err != nil {
			return nil, err
		}

		return nil, &httpErr
	} else {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var result RES
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		return &result, nil
	}
}
