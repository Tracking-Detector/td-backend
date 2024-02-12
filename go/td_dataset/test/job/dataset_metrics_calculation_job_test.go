package job_test

import (
	"testing"

	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/mocks"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/testsupport"
	"github.com/Tracking-Detector/td-backend/go/td_dataset/job"
	"github.com/stretchr/testify/suite"
)

func TestDatasetMetricsCalculationJob(t *testing.T) {
	suite.Run(t, &DatasetMetricsCalculationJob{})
}

type DatasetMetricsCalculationJob struct {
	suite.Suite
	calculationJob *job.DatasetMetricJob
	datasetService *mocks.IDatasetService
	requestService *mocks.IRequestService
}

func (suite *DatasetMetricsCalculationJob) SetupTest() {
	suite.datasetService = new(mocks.IDatasetService)
	suite.requestService = new(mocks.IRequestService)
	suite.calculationJob = job.NewDatasetMetricJob(suite.datasetService, suite.requestService)
}

func (suite *DatasetMetricsCalculationJob) TestExecute_Success() {
	// given
	datasets := []*model.Dataset{
		{
			ID:          "someId1",
			Name:        "All",
			Description: "someDescription",
			Label:       "",
		},
		{
			ID:          "someId2",
			Name:        "Test",
			Description: "someDescription",
			Label:       "test",
		},
	}
	requests := testsupport.LoadRequestJson()
	for i, request := range requests {
		if i%2 == 0 {
			request.Dataset = "test"
		}
	}
	// when
	// suite.calculationJob.Execute()
	// then
	suite.Len(datasets, 2)
}
