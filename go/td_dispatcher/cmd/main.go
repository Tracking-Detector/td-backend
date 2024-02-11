package main

import (
	"context"
	"time"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/queue"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_dispatcher/controller"
)

func main() {
	time.Sleep(30 * time.Second)
	ctx := context.Background()
	db := config.GetDatabase(config.ConnectDB(ctx))
	rabbitCh := config.ConnectRabbitMQ()

	channelAdapter := queue.NewRabbitMQChannelAdapter(rabbitCh)
	exporterRepo := repository.NewMongoExporterRepository(db)
	modelRepo := repository.NewMongoModelRepository(db)
	requestRepo := repository.NewMongoRequestRepository(db)
	trainingRunRepo := repository.NewMongoTrainingRunRepository(db)
	exportRunRepo := repository.NewMongoExportRunRunRepository(db)
	datasetRepo := repository.NewMongoDatasetRepository(db)

	trainingRunService := service.NewTraingingrunService(trainingRunRepo)
	exporterService := service.NewExporterService(exporterRepo)
	exportRunService := service.NewExportRunService(exportRunRepo)
	datasetService := service.NewDatasetService(datasetRepo, requestRepo)
	modelService := service.NewModelService(modelRepo, trainingRunService)
	publisherService := service.NewPublishService(channelAdapter)

	dispatchController := controller.NewDispatchController(exporterService, publisherService, modelService, datasetService, exportRunService)
	app := util.DefaultFiberApp()
	dispatchController.RegisterRoutes(app).Listen(":8081")
}
