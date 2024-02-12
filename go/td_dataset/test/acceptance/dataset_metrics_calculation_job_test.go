package acceptance

import (
	"testing"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/testsupport"
	"github.com/Tracking-Detector/td-backend/go/td_dataset/job"
	"github.com/stretchr/testify/suite"
)

func TestDatasetMetricsCalculationJobAcceptance(t *testing.T) {
	suite.Run(t, &DatasetMetricsCalculationJobAcceptanceTest{})
}

type DatasetMetricsCalculationJobAcceptanceTest struct {
	testsupport.AcceptanceTest
	suite.Suite
	requestRepo      model.RequestRepository
	requestService   service.IRequestService
	datasetRepo      model.DatasetRepository
	datasetService   service.IDatasetService
	datasetMetricJob job.DatasetMetricJob
}

func (suite *DatasetMetricsCalculationJobAcceptanceTest) SetupTest() {
	suite.SetupIntegration()
	suite.requestRepo = repository.NewMongoRequestRepository(config.GetDatabase(config.ConnectDB(suite.Ctx)))
	suite.datasetRepo = repository.NewMongoDatasetRepository(config.GetDatabase(config.ConnectDB(suite.Ctx)))
	suite.requestService = service.NewRequestService(suite.requestRepo)
	suite.datasetService = service.NewDatasetService(suite.datasetRepo, suite.requestRepo)
	suite.datasetMetricJob = *job.NewDatasetMetricJob(suite.datasetService, suite.requestService)
	suite.requestRepo.DeleteAll(suite.Ctx)
	suite.datasetRepo.DeleteAll(suite.Ctx)
}

func (suite *DatasetMetricsCalculationJobAcceptanceTest) TearDownSuite() {
	suite.TeardownIntegration()
}

func (suite *DatasetMetricsCalculationJobAcceptanceTest) TestExecute_Success() {
	// given
	suite.datasetRepo.Save(suite.Ctx, &model.Dataset{
		Name:        "Training dataset",
		Description: "This is a training dataset.",
		Label:       "train",
	})
	requests := testsupport.LoadRequestJson()
	for _, request := range requests {
		request.Dataset = "train"
	}
	suite.requestRepo.SaveAll(suite.Ctx, requests)
	suite.datasetService.ReloadCache(suite.Ctx)
	// when
	suite.datasetMetricJob.Execute()
	// then
	datasets := suite.datasetService.GetAllDatasets()
	suite.Equal(1, len(datasets))
	suite.Equal(10, datasets[0].Metrics.Total)
	suite.Equal(10, datasets[0].Metrics.ReducerMetric[0].Total)
	suite.Equal(10, datasets[0].Metrics.ReducerMetric[0].NonTracker)
	suite.Equal(0, datasets[0].Metrics.ReducerMetric[0].Tracker)
	suite.Equal(10, datasets[0].Metrics.ReducerMetric[1].Total)
	suite.Equal(9, datasets[0].Metrics.ReducerMetric[1].NonTracker)
	suite.Equal(1, datasets[0].Metrics.ReducerMetric[1].Tracker)
}

func (suite *DatasetMetricsCalculationJobAcceptanceTest) TestExecuteMultiple_Success() {
	// given
	suite.datasetRepo.Save(suite.Ctx, &model.Dataset{
		Name:        "Training dataset",
		Description: "This is a training dataset.",
		Label:       "train",
	})
	suite.datasetRepo.Save(suite.Ctx, &model.Dataset{
		Name:        "Test dataset",
		Description: "This is a test dataset.",
		Label:       "test",
	})
	requests := testsupport.LoadRequestJson()
	for i, request := range requests {
		if i%3 == 0 {
			request.Dataset = "train"
		} else {
			request.Dataset = "test"
		}
	}
	suite.requestRepo.SaveAll(suite.Ctx, requests)
	suite.datasetService.ReloadCache(suite.Ctx)
	// when
	suite.datasetMetricJob.Execute()
	// then
	datasets := suite.datasetService.GetAllDatasets()
	suite.Equal(2, len(datasets))
	suite.Equal(4, datasets[0].Metrics.Total)
	suite.Equal(4, datasets[0].Metrics.ReducerMetric[0].Total)
	suite.Equal(4, datasets[0].Metrics.ReducerMetric[0].NonTracker)
	suite.Equal(0, datasets[0].Metrics.ReducerMetric[0].Tracker)
	suite.Equal(4, datasets[0].Metrics.ReducerMetric[1].Total)
	suite.Equal(3, datasets[0].Metrics.ReducerMetric[1].NonTracker)
	suite.Equal(1, datasets[0].Metrics.ReducerMetric[1].Tracker)
	suite.Equal(6, datasets[1].Metrics.Total)
	suite.Equal(6, datasets[1].Metrics.ReducerMetric[0].Total)
	suite.Equal(6, datasets[1].Metrics.ReducerMetric[0].NonTracker)
	suite.Equal(0, datasets[1].Metrics.ReducerMetric[0].Tracker)
	suite.Equal(6, datasets[1].Metrics.ReducerMetric[1].Total)
	suite.Equal(6, datasets[1].Metrics.ReducerMetric[1].NonTracker)
	suite.Equal(0, datasets[1].Metrics.ReducerMetric[1].Tracker)
}
