package main

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_export/controller"
)

func main() {
	ctx := context.Background()
	exporterRepo := repository.NewMongoExporterRepository(config.GetDatabase(config.ConnectDB(ctx)))
	exporterService := service.NewExporterService(exporterRepo)
	// TODO fix restart bug
	exporterService.InitInCodeExports(ctx)

	exporterController := controller.NewExportController(exporterService)
	app := util.DefaultFiberApp()
	exporterController.RegisterRoutes(app).Listen(":8081")
}
