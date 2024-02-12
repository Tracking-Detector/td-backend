package acceptance

import (
	"os"

	"testing"
	"time"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/queue"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/storage"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/testsupport"
	"github.com/Tracking-Detector/td-backend/go/td_exporter/consumer"
	"github.com/Tracking-Detector/td-backend/go/td_exporter/job"
	"github.com/stretchr/testify/suite"
)

func TestConsumerAcceptance(t *testing.T) {
	suite.Run(t, &ExportConsumerAcceptanceTest{})
}

type ExportConsumerAcceptanceTest struct {
	testsupport.AcceptanceTest
	suite.Suite
	storageService   *service.MinIOStorageService
	publisherService *service.PublishService
	datasetService   *service.DatasetService
	requestRepo      *repository.MongoRequestRepository
	datasetRepo      *repository.MongoDatasetRepository
	exporterRepo     *repository.MongoExporterRepository
	exportRunRepo    *repository.MongoExportRunRunRepository
	exportRunService *service.ExportRunService
	exportConsumer   *consumer.ExportMessageConsumer
	queueAdapter     *queue.RabbitMQChannelAdapter
}

func (suite *ExportConsumerAcceptanceTest) SetupSuite() {
	suite.SetupIntegration()
	suite.exporterRepo = repository.NewMongoExporterRepository(config.GetDatabase(config.ConnectDB(suite.Ctx)))
	suite.requestRepo = repository.NewMongoRequestRepository(config.GetDatabase(config.ConnectDB(suite.Ctx)))
	suite.exportRunRepo = repository.NewMongoExportRunRunRepository(config.GetDatabase(config.ConnectDB(suite.Ctx)))
	suite.datasetRepo = repository.NewMongoDatasetRepository(config.GetDatabase(config.ConnectDB(suite.Ctx)))
	suite.datasetService = service.NewDatasetService(suite.datasetRepo, suite.requestRepo)
	suite.exportRunService = service.NewExportRunService(suite.exportRunRepo)
	minioClient := config.ConnectMinio()
	rabbitMqChannel := config.ConnectRabbitMQ()
	suite.queueAdapter = queue.NewRabbitMQChannelAdapter(rabbitMqChannel)
	minioStorageAdapter := storage.NewMinIOStorageAdapter(minioClient)
	suite.storageService = service.NewMinIOStorageService(minioStorageAdapter)
	suite.publisherService = service.NewPublishService(suite.queueAdapter)
	suite.datasetRepo.DeleteAll(suite.Ctx)
	internalJob := job.NewInternalExportJob(suite.requestRepo, suite.storageService)
	externJob := job.NewExternalExportJob(suite.requestRepo, suite.storageService)
	suite.exportConsumer = consumer.NewExportMessageConsumer(internalJob, externJob,
		suite.exportRunService, suite.queueAdapter, service.NewExporterService(suite.exporterRepo), suite.datasetService)
	go func() {
		suite.exportConsumer.Consume()
	}()
}

func (suite *ExportConsumerAcceptanceTest) SetupTest() {
	suite.requestRepo.DeleteAll(suite.Ctx)
	suite.exporterRepo.DeleteAll(suite.Ctx)
	suite.exportRunRepo.DeleteAll(suite.Ctx)
	suite.queueAdapter.PurgeQueue(config.EnvExportBucketName(), false)
}

func (suite *ExportConsumerAcceptanceTest) TearDownSuite() {
	suite.exportConsumer.Stop()
	suite.TeardownIntegration()
}

func (suite *ExportConsumerAcceptanceTest) TestExportConsumer_ForExternalExporterSuccess() {
	// given
	suite.storageService.VerifyBucketExists(suite.Ctx, config.EnvExportBucketName())
	suite.storageService.VerifyBucketExists(suite.Ctx, config.EnvExtractorBucketName())
	extractorFilePath := "./testdata/exporter/exporter204.js"
	fileLoc := "exporter204.js"
	file, _ := os.Open(extractorFilePath)
	suite.storageService.PutObject(suite.Ctx, config.EnvExtractorBucketName(), "exporter204.js", file, -1, "application/javascript")
	exporter := &model.Exporter{
		ID:                   "someId",
		Name:                 "someName",
		Description:          "someDescription",
		Dimensions:           []int{204, 1},
		Type:                 model.JS,
		ExportScriptLocation: &fileLoc,
	}
	dataset, _ := suite.datasetRepo.Save(suite.Ctx, &model.Dataset{
		Name:        "someName",
		Description: "someDescription",
		Label:       "",
	})
	suite.exporterRepo.Save(suite.Ctx, exporter)
	requests := testsupport.LoadRequestJson()
	suite.requestRepo.SaveAll(suite.Ctx, requests)

	// when
	suite.publisherService.EnqueueExportJob("someId", "EasyPrivacy", dataset.ID)
	time.Sleep(1 * time.Second)
	// then
	suite.exportConsumer.Wg.Wait()
	export, err := suite.storageService.GetObject(suite.Ctx, config.EnvExportBucketName(), "someName_EasyPrivacy_.csv.gz")
	suite.NoError(err)
	expectedCsv := testsupport.LoadFile("./testdata/requests/expected_encoding.csv")
	actualCsv := testsupport.Unzip(export)
	suite.Equal(expectedCsv, actualCsv)
	count, _ := suite.exportRunRepo.Count(suite.Ctx)
	suite.Equal(int64(1), count)
	exportRuns, _ := suite.exportRunRepo.FindAll(suite.Ctx)
	suite.Equal("someId", exportRuns[0].ExporterId)
	suite.Equal("someName", exportRuns[0].Name)
	suite.Equal("EasyPrivacy", exportRuns[0].Reducer)
	suite.Equal(dataset.ID, exportRuns[0].Dataset)
	suite.Equal(9, exportRuns[0].Metrics.NonTracker)
	suite.Equal(1, exportRuns[0].Metrics.Tracker)
	suite.Equal(10, exportRuns[0].Metrics.Total)

}
