package main

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_dataset/controller"
	"github.com/Tracking-Detector/td-backend/go/td_dataset/job"
	"github.com/robfig/cron"
)

func StartCron(datasetCalculatorJob *job.DatasetMetricJob) {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		datasetCalculatorJob.Execute()
	})
	c.Start()
}

func main() {
	ctx := context.Background()
	db := config.GetDatabase(config.ConnectDB(ctx))
	requestRepo := repository.NewMongoRequestRepository(db)
	datasetRepo := repository.NewMongoDatasetRepository(db)

	datasetService := service.NewDatasetService(datasetRepo, requestRepo)
	requestService := service.NewRequestService(requestRepo)
	datasetCalculationJob := job.NewDatasetMetricJob(datasetService, requestService)
	datasetController := controller.NewDatasetController(datasetService)
	go StartCron(datasetCalculationJob)
	app := util.DefaultFiberApp()
	go datasetController.RegisterRoutes(app).Listen(":8081")
}
