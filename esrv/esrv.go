package esrv

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/thienhaole92/vnd/logger"
)

type Esrv struct {
	name        string
	gracePeriod time.Duration
	e           *echo.Echo
	h           *http.Server
}

func NewEsrv(e *echo.Echo, sc *serverConfig) *Esrv {
	log := logger.GetLogger(sc.Name)
	defer log.Sync()

	log.Infow("init server", "name", sc.Name, "host", sc.Host, "port", sc.Port)
	h := &http.Server{
		Addr:         sc.Addr(),
		Handler:      e,
		ReadTimeout:  sc.ReadTimeout,
		WriteTimeout: sc.WriteTimeout,
	}

	return &Esrv{
		name:        sc.Name,
		gracePeriod: sc.GracePeriod,
		e:           e,
		h:           h,
	}
}

func (es *Esrv) ListenAndServe() {
	go func(s *Esrv) {
		log := logger.GetLogger(s.name)
		defer log.Sync()

		log.Infow("start server", "name", s.name)
		if err := s.h.ListenAndServe(); err != http.ErrServerClosed {
			log.Panic(err)
		}
	}(es)
}

func (es *Esrv) Shutdown() {
	log := logger.GetLogger(es.name)
	defer log.Sync()

	log.Infow("shutdown server", "name", es.name)

	ctx, cancel := context.WithTimeout(context.TODO(), es.gracePeriod)
	defer cancel()

	if err := es.h.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("could not gracefully shut down web server. error: %s" + err.Error()))
	}
}
