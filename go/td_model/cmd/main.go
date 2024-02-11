package main

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_model/controller"
)

func main() {
	ctx := context.Background()
	trainingrunRepo := repository.NewMongoTrainingRunRepository(config.GetDatabase(config.ConnectDB(ctx)))
	modelRepo := repository.NewMongoModelRepository(config.GetDatabase(config.ConnectDB(ctx)))
	trainingrunService := service.NewTraingingrunService(trainingrunRepo)
	modelService := service.NewModelService(modelRepo, trainingrunService)
	modelController := controller.NewModelController(trainingrunService, modelService)

	app := util.DefaultFiberApp()
	modelController.RegisterRoutes(app).Listen(":8081")
}
