package main

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_request/controller"
)

func main() {
	ctx := context.Background()
	requestRepo := repository.NewMongoRequestRepository(config.GetDatabase(config.ConnectDB(ctx)))
	requestService := service.NewRequestService(requestRepo)
	requestController := controller.NewRequestController(requestService)

	app := util.DefaultFiberApp()
	requestController.RegisterRoutes(app).Listen(":8081")
}
