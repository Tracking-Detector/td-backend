package consumer_test

import (
	"os"

	"testing"
	"time"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/message"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/mocks"
	"github.com/Tracking-Detector/td-backend/go/td_exporter/consumer"
	mocks_exporter "github.com/Tracking-Detector/td-backend/go/td_exporter/test/mocks"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestExportConsumer(t *testing.T) {
	suite.Run(t, &ExportConsumerTest{})
}

type ExportConsumerTest struct {
	suite.Suite
	exportConsumer   *consumer.ExportMessageConsumer
	internalJob      *mocks_exporter.IExportJob
	externalJob      *mocks_exporter.IExportJob
	queueAdapter     *mocks.IQueueChannelAdapter
	exportRunService *mocks.IExportRunService
	exporterService  *mocks.IExporterService
	datasetService   *mocks.IDatasetService
}

func (suite *ExportConsumerTest) SetupTest() {
	suite.internalJob = new(mocks_exporter.IExportJob)
	suite.externalJob = new(mocks_exporter.IExportJob)
	suite.queueAdapter = new(mocks.IQueueChannelAdapter)
	suite.exporterService = new(mocks.IExporterService)
	suite.exportRunService = new(mocks.IExportRunService)
	suite.datasetService = new(mocks.IDatasetService)
	suite.exportConsumer = consumer.NewExportMessageConsumer(suite.internalJob, suite.externalJob, suite.exportRunService, suite.queueAdapter, suite.exporterService, suite.datasetService)
}

func (suite *ExportConsumerTest) TestConsume_SuccessInternal() {
	// given
	os.Setenv("EXPORT_QUEUE", "export")
	exporter := &model.Exporter{
		ID:          "someId",
		Name:        "someName",
		Description: "someDescription",
		Dimensions:  []int{204, 1},
		Type:        model.IN_SERVICE,
	}
	dataset := &model.Dataset{
		ID:    "someId",
		Name:  "someName",
		Label: "",
	}
	jobs := []*message.JobPayload{message.NewJob("export", []string{"someId", "or", "someId"})}
	suite.datasetService.On("GetDatasetByID", mock.Anything, "someId").Return(dataset, nil)
	suite.exportRunService.On("Save", mock.Anything, mock.Anything).Return(&model.ExportRun{
		ID:         "someRunId",
		ExporterId: exporter.ID,
		Name:       exporter.Name,
		Dataset:    "someId",
		Start:      time.Now(),
		End:        time.Now(),
	}, nil)
	suite.queueAdapter.On("Consume", config.EnvExportQueueName(), "ExportConsumer", true, false, false, false, mock.Anything).Return(suite.createChan(jobs), nil)
	suite.exporterService.On("FindByID", mock.Anything, "someId").Return(exporter, nil)
	suite.internalJob.On("Execute", exporter, "or", "").Return(nil)
	// when
	suite.exportConsumer.Consume()
	suite.exportConsumer.Wg.Wait()
	// then
	suite.queueAdapter.AssertCalled(suite.T(), "Consume", config.EnvExportQueueName(), "ExportConsumer", true, false, false, false, mock.Anything)
	suite.exporterService.AssertCalled(suite.T(), "FindByID", mock.Anything, "someId")
	suite.exportRunService.AssertNumberOfCalls(suite.T(), "Save", 2)
	suite.internalJob.AssertCalled(suite.T(), "Execute", exporter, "or", "")
}

func (suite *ExportConsumerTest) TestConsume_SuccessExternal() {
	// given
	os.Setenv("EXPORT_QUEUE", "export")
	exporter := &model.Exporter{
		ID:          "someId",
		Name:        "someName",
		Description: "someDescription",
		Dimensions:  []int{204, 1},
		Type:        model.JS,
	}
	dataset := &model.Dataset{
		ID:    "someId",
		Name:  "someName",
		Label: "",
	}
	jobs := []*message.JobPayload{message.NewJob("export", []string{"someId", "or", "someId"})}
	suite.datasetService.On("GetDatasetByID", mock.Anything, "someId").Return(dataset, nil)
	suite.queueAdapter.On("Consume", config.EnvExportQueueName(), "ExportConsumer", true, false, false, false, mock.Anything).Return(suite.createChan(jobs), nil)
	suite.exportRunService.On("Save", mock.Anything, mock.Anything).Return(&model.ExportRun{
		ID:         "someRunId",
		ExporterId: exporter.ID,
		Name:       exporter.Name,
		Dataset:    "someId",
		Start:      time.Now(),
		End:        time.Now(),
	}, nil)
	suite.exporterService.On("FindByID", mock.Anything, "someId").Return(exporter, nil)
	suite.externalJob.On("Execute", exporter, "or", "").Return(nil)
	// when
	suite.exportConsumer.Consume()
	suite.exportConsumer.Wg.Wait()
	// then
	suite.queueAdapter.AssertCalled(suite.T(), "Consume", config.EnvExportQueueName(), "ExportConsumer", true, false, false, false, mock.Anything)
	suite.exporterService.AssertCalled(suite.T(), "FindByID", mock.Anything, "someId")
	suite.exportRunService.AssertNumberOfCalls(suite.T(), "Save", 2)
	suite.externalJob.AssertCalled(suite.T(), "Execute", exporter, "or", "")
}

func (suite *ExportConsumerTest) TestConsume_SuccessMultiple() {
	// given
	os.Setenv("EXPORT_QUEUE", "export")
	exporter1 := &model.Exporter{
		ID:          "someId1",
		Name:        "someName",
		Description: "someDescription",
		Dimensions:  []int{204, 1},
		Type:        model.IN_SERVICE,
	}
	exporter2 := &model.Exporter{
		ID:          "someId2",
		Name:        "someName",
		Description: "someDescription",
		Dimensions:  []int{204, 1},
		Type:        model.JS,
	}
	dataset := &model.Dataset{
		ID:    "someId",
		Name:  "someName",
		Label: "",
	}
	jobs := []*message.JobPayload{message.NewJob("export", []string{"someId1", "or", "someId"}),
		message.NewJob("export", []string{"someId2", "or", "someId"})}
	suite.datasetService.On("GetDatasetByID", mock.Anything, "someId").Return(dataset, nil)
	suite.exportRunService.On("Save", mock.Anything, mock.Anything).Return(&model.ExportRun{
		ID:         "someRunId",
		ExporterId: exporter1.ID,
		Name:       exporter1.Name,
		Dataset:    "someId",
		Start:      time.Now(),
		End:        time.Now(),
	}, nil)
	suite.queueAdapter.On("Consume", config.EnvExportQueueName(), "ExportConsumer", true, false, false, false, mock.Anything).Return(suite.createChan(jobs), nil)
	suite.exporterService.On("FindByID", mock.Anything, "someId1").Return(exporter1, nil)
	suite.exporterService.On("FindByID", mock.Anything, "someId2").Return(exporter2, nil)
	suite.internalJob.On("Execute", exporter1, "or", "").Return(nil)
	suite.externalJob.On("Execute", exporter2, "or", "").Return(nil)
	// when
	suite.exportConsumer.Consume()
	suite.exportConsumer.Wg.Wait()
	// then
	suite.queueAdapter.AssertCalled(suite.T(), "Consume", config.EnvExportQueueName(), "ExportConsumer", true, false, false, false, mock.Anything)
	suite.exporterService.AssertCalled(suite.T(), "FindByID", mock.Anything, "someId1")
	suite.exporterService.AssertCalled(suite.T(), "FindByID", mock.Anything, "someId2")
	suite.exportRunService.AssertNumberOfCalls(suite.T(), "Save", 4)
	suite.internalJob.AssertCalled(suite.T(), "Execute", exporter1, "or", "")
	suite.externalJob.AssertCalled(suite.T(), "Execute", exporter2, "or", "")
}

func (suite *ExportConsumerTest) createChan(jobs []*message.JobPayload) <-chan amqp.Delivery {
	jobCh := make(chan amqp.Delivery, len(jobs))
	for _, job := range jobs {
		ser, _ := job.Serialize()
		jobCh <- amqp.Delivery{
			ContentType:  "text/plain",
			Body:         []byte(ser),
			DeliveryMode: amqp.Persistent,
		}
	}
	defer close(jobCh)
	return jobCh
}
