package main

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/storage"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_download/controller"
)

func main() {
	ctx := context.Background()
	minioClient := config.ConnectMinio()
	minioStorageAdapter := storage.NewMinIOStorageAdapter(minioClient)
	storageService := service.NewMinIOStorageService(minioStorageAdapter)
	storageService.VerifyBucketExists(ctx, config.EnvExportBucketName())
	storageService.VerifyBucketExists(ctx, config.EnvModelBucketName())
	downloadController := controller.NewDownloadController(storageService)

	app := util.DefaultFiberApp()
	downloadController.RegisterRoutes(app).Listen(":8081")
}
