package route

import (
	topic "user-service-v5/internal/event"
	"user-service-v5/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thienhaole92/vnd/logger"
	vndmiddleware "github.com/thienhaole92/vnd/middleware"
	"github.com/thienhaole92/vnd/publisher"
	"github.com/thienhaole92/vnd/rest"
	"github.com/thienhaole92/vnd/runner"
	"github.com/thienhaole92/vnd/wpub"
)

type V1 struct {
	*echo.Group
}

func (v1 *V1) Configure(rn *runner.Runner) error {
	pub, err := createPublisher(rn)
	if err != nil {
		return err
	}

	s := service.NewService(pub)

	return v1.registerRoutes(s)
}

func (v1 *V1) registerRoutes(s *service.Service) error {
	v1.Use(vndmiddleware.RequestID(middleware.DefaultSkipper))

	v1.GET("/sync-all-user", rest.Wrapper(s.SyncAllUser))

	return nil
}

func createPublisher(rn *runner.Runner) (publisher.Publisher, error) {
	log := logger.GetLogger("createPublisher")
	defer log.Sync()

	pc, err := wpub.NewConfig()
	if err != nil {
		log.Errorw("failed to get publisher config", "error", err)
		return nil, err
	}

	pub, err := wpub.NewPublisher(pc, topic.TOPIC_PRIVATE_USER_INFO)
	if err != nil {
		log.Errorw("failed to create publisher", "error", err)
		return nil, err
	}

	rn.AddShutdownHook("close_publisher", func(*runner.Runner) error {
		pub.Close()
		return nil
	})

	return pub, nil
}
