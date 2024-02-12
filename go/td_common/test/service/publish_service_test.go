package service_test

import (
	"os"

	"testing"

	"github.com/Tracking-Detector/td-backend/go/td_common/message"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/mocks"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/suite"
)

func TestPublish(t *testing.T) {
	suite.Run(t, &PublishServiceTest{})
}

type PublishServiceTest struct {
	suite.Suite
	publishService *service.PublishService
	queueAdapter   *mocks.IQueueChannelAdapter
}

func (suite *PublishServiceTest) SetupTest() {
	suite.queueAdapter = new(mocks.IQueueChannelAdapter)
	suite.publishService = service.NewPublishService(suite.queueAdapter)
}

func (suite *PublishServiceTest) TestEnqueueTrainingJob_Success() {
	// given
	os.Setenv("TRAIN_QUEUE", "training")
	modelId := "modelId"
	exporterId := "exporterId"
	reducer := "or"
	job := message.JobPayload{
		FunctionName: "train_model",
		Args:         []string{modelId, exporterId, reducer},
	}
	msg, _ := job.Serialize()
	suite.queueAdapter.On("Publish", "", "training", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(msg),
		DeliveryMode: amqp.Persistent,
	}).Return(nil)
	// when
	suite.publishService.EnqueueTrainingJob(modelId, exporterId, reducer)
	// then
	suite.queueAdapter.AssertCalled(suite.T(), "Publish", "", "training", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(msg),
		DeliveryMode: amqp.Persistent,
	})
}

func (suite *PublishServiceTest) TestEnqueueExportJob_Success() {
	// given
	os.Setenv("EXPORT_QUEUE", "exports")
	exporterId := "exporterId"
	reducer := "or"
	dataset := "human"
	job := message.JobPayload{
		FunctionName: "export",
		Args:         []string{exporterId, reducer, dataset},
	}
	msg, _ := job.Serialize()
	suite.queueAdapter.On("Publish", "", "exports", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(msg),
		DeliveryMode: amqp.Persistent,
	}).Return(nil)
	// when
	suite.publishService.EnqueueExportJob(exporterId, reducer, dataset)
	// then
	suite.queueAdapter.AssertCalled(suite.T(), "Publish", "", "exports", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(msg),
		DeliveryMode: amqp.Persistent,
	})
}
