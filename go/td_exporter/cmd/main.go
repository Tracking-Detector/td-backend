package main

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/queue"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/storage"
	"github.com/Tracking-Detector/td-backend/go/td_exporter/consumer"
	"github.com/Tracking-Detector/td-backend/go/td_exporter/job"
)

func main() {
	ctx := context.Background()
	db := config.GetDatabase(config.ConnectDB(ctx))
	minioAdpater := storage.NewMinIOStorageAdapter(config.ConnectMinio())
	rabbitMQAdapter := queue.NewRabbitMQChannelAdapter(config.ConnectRabbitMQ())

	requestRepo := repository.NewMongoRequestRepository(db)
	exporterRepo := repository.NewMongoExporterRepository(db)
	datasetRepo := repository.NewMongoDatasetRepository(db)
	exportRunRepo := repository.NewMongoExportRunRunRepository(db)

	storageService := service.NewMinIOStorageService(minioAdpater)

	storageService.VerifyBucketExists(ctx, config.EnvExtractorBucketName())
	storageService.VerifyBucketExists(ctx, config.EnvModelBucketName())
	storageService.VerifyBucketExists(ctx, config.EnvExportBucketName())

	exporterService := service.NewExporterService(exporterRepo)
	datasetService := service.NewDatasetService(datasetRepo, requestRepo)
	exportRunService := service.NewExportRunService(exportRunRepo)

	internalExportJob := job.NewInternalExportJob(requestRepo, storageService)
	externalExportJob := job.NewExternalExportJob(requestRepo, storageService)

	consumer.NewExportMessageConsumer(internalExportJob, externalExportJob, exportRunService, rabbitMQAdapter, exporterService, datasetService).Consume()
	select {}
}
