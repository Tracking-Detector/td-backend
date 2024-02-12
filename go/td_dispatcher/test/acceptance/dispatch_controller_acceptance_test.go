package acceptance

import (
	"fmt"

	"testing"
	"time"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"

	"github.com/Tracking-Detector/td-backend/go/td_common/queue"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/testsupport"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_dispatcher/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/suite"
)

func TestDispatchControllerAcceptance(t *testing.T) {
	suite.Run(t, &DispatchControllerAcceptanceTest{})
}

type DispatchControllerAcceptanceTest struct {
	suite.Suite
	testsupport.AcceptanceTest
	app                *fiber.App
	publishController  *controller.DispatchController
	exporterService    service.IExporterService
	exportRunService   service.IExportRunService
	publishService     service.IPublishService
	exporterRepo       model.ExporterRepository
	requestRepo        model.RequestRepository
	exportRunRepo      model.ExportRunRepository
	trainingRunRepo    model.TrainingRunRepository
	modelRepo          model.ModelRepository
	datasetRepo        model.DatasetRepository
	datasetService     service.IDatasetService
	queueAdapter       queue.IQueueChannelAdapter
	modelService       service.IModelService
	trainingRunService service.ITrainingrunService
	testConsumer       *testsupport.TestQueueConsumer
	rabbitMq           *amqp.Channel
}

func (suite *DispatchControllerAcceptanceTest) SetupSuite() {
	suite.SetupIntegration()
	mongoClient := config.ConnectDB(suite.Ctx)
	suite.rabbitMq = config.ConnectRabbitMQ()
	suite.requestRepo = repository.NewMongoRequestRepository(config.GetDatabase(mongoClient))
	suite.queueAdapter = queue.NewRabbitMQChannelAdapter(suite.rabbitMq)
	suite.exporterRepo = repository.NewMongoExporterRepository(config.GetDatabase(mongoClient))
	suite.exportRunRepo = repository.NewMongoExportRunRunRepository(config.GetDatabase(mongoClient))
	suite.modelRepo = repository.NewMongoModelRepository(config.GetDatabase(mongoClient))
	suite.trainingRunRepo = repository.NewMongoTrainingRunRepository(config.GetDatabase(mongoClient))
	suite.datasetRepo = repository.NewMongoDatasetRepository(config.GetDatabase(mongoClient))

	suite.trainingRunService = service.NewTraingingrunService(suite.trainingRunRepo)
	suite.datasetService = service.NewDatasetService(suite.datasetRepo, suite.requestRepo)
	suite.exportRunService = service.NewExportRunService(suite.exportRunRepo)
	suite.exporterService = service.NewExporterService(suite.exporterRepo)
	suite.modelService = service.NewModelService(suite.modelRepo, suite.trainingRunService)
	suite.publishService = service.NewPublishService(suite.queueAdapter)

	suite.publishController = controller.NewDispatchController(suite.exporterService,
		suite.publishService, suite.modelService, suite.datasetService, suite.exportRunService)
	go func() {
		suite.app = util.DefaultFiberApp()
		suite.publishController.RegisterRoutes(suite.app).Listen(":8081")
	}()
	time.Sleep(5 * time.Second)
}

func (suite *DispatchControllerAcceptanceTest) SetupTest() {
	suite.exporterRepo.DeleteAll(suite.Ctx)
	suite.modelRepo.DeleteAll(suite.Ctx)
	suite.trainingRunRepo.DeleteAll(suite.Ctx)
	suite.datasetRepo.DeleteAll(suite.Ctx)
	suite.exportRunRepo.DeleteAll(suite.Ctx)
	suite.queueAdapter.PurgeQueue(config.EnvExportQueueName(), false)
	suite.queueAdapter.PurgeQueue(config.EnvTrainQueueName(), false)
	suite.testConsumer = testsupport.NewTestQueueConsumer(queue.NewRabbitMQChannelAdapter(suite.rabbitMq))
}

func (suite *DispatchControllerAcceptanceTest) TearDownTest() {
	suite.testConsumer.ClearMessages()
}

func (suite *DispatchControllerAcceptanceTest) TearDownSuite() {
	suite.app.Shutdown()
	suite.TeardownIntegration()
}

func (suite *DispatchControllerAcceptanceTest) TestHealth_Success() {
	// given
	// when
	resp, err := testsupport.Get("http://localhost:8081/dispatch/health")

	// then
	suite.NoError(err)
	suite.Equal(200, resp.StatusCode)
}

// func (suite *DispatchControllerAcceptanceTest) TestDispatchExportJob_Success() {
// 	// given
// 	loc := "exporter.js"
// 	exporter, _ := suite.exporterRepo.Save(suite.Ctx, &model.Exporter{
// 		Name:                 "ExporterJs204",
// 		Description:          "ExporterJs204",
// 		Dimensions:           []int{204, 1},
// 		Type:                 model.JS,
// 		ExportScriptLocation: &loc,
// 	})
// 	dataset, _ := suite.datasetRepo.Save(suite.Ctx, &model.Dataset{
// 		Name:        "Verification",
// 		Description: "Can be used for verifaction",
// 		Label:       "verifiaction",
// 	})
// 	reducer := "or"
// 	// when
// 	go suite.testConsumer.Consume(config.EnvExportQueueName(), 1)
// 	time.Sleep(5 * time.Second)
// 	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/export/%s/%s/%s", exporter.ID, reducer, dataset.ID), "", "")
// 	suite.testConsumer.WaitForMessages(config.EnvExportQueueName(), 1)
// 	// then
// 	suite.NoError(err)
// 	suite.Equal(201, resp.StatusCode)
// 	suite.Equal(1, len(suite.testConsumer.QueueMessages[config.EnvExportQueueName()]))
// 	suite.Equal(fmt.Sprintf(`{"functionName":"export","args":["%s","%s","%s"]}`, exporter.ID, reducer, dataset.ID), suite.testConsumer.QueueMessages[config.EnvExportQueueName()][0])
// }

func (suite *DispatchControllerAcceptanceTest) TestDispatchExportJob_ErrorReducerNotFound() {
	// given
	loc := "exporter.js"
	exporter, _ := suite.exporterRepo.Save(suite.Ctx, &model.Exporter{
		Name:                 "ExporterJs204",
		Description:          "ExporterJs204",
		Dimensions:           []int{204, 1},
		Type:                 model.JS,
		ExportScriptLocation: &loc,
	})
	dataset, _ := suite.datasetRepo.Save(suite.Ctx, &model.Dataset{
		Name:        "Verification",
		Description: "Can be used for verifaction",
		Label:       "verifiaction",
	})
	reducer := "notKnown"
	// when
	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/export/%s/%s/%s", exporter.ID, reducer, dataset.ID), "", "")
	// then
	suite.NoError(err)
	suite.Equal(400, resp.StatusCode)
	suite.Equal(`{"success":false,"message":"The reducer type is not valid"}`, resp.Body)
}

func (suite *DispatchControllerAcceptanceTest) TestDispatchExportJob_ErrorDatasetNotFound() {
	loc := "exporter.js"
	exporter, _ := suite.exporterRepo.Save(suite.Ctx, &model.Exporter{
		Name:                 "ExporterJs204",
		Description:          "ExporterJs204",
		Dimensions:           []int{204, 1},
		Type:                 model.JS,
		ExportScriptLocation: &loc,
	})
	datasetNotInDbId := "5f5e7e3e3e3e3e3e3e3e3e3e"
	reducer := "or"
	// when
	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/export/%s/%s/%s", exporter.ID, reducer, datasetNotInDbId), "", "")
	// then
	suite.NoError(err)
	suite.Equal(404, resp.StatusCode)
	suite.Equal(`{"success":false,"message":"The dataset for the given id does not exist."}`, resp.Body)
}

func (suite *DispatchControllerAcceptanceTest) TestDispatchExportJob_ErrorExporterNotFound() {
	// given
	randomExporterId := "5f5e7e3e3e3e3e3e3e3e3e3e"
	dataset, _ := suite.datasetRepo.Save(suite.Ctx, &model.Dataset{
		Name:        "Verification",
		Description: "Can be used for verifaction",
		Label:       "verifiaction",
	})
	reducer := "or"
	// when
	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/export/%s/%s/%s", randomExporterId, reducer, dataset.ID), "", "")
	// then
	suite.NoError(err)
	suite.Equal(404, resp.StatusCode)
	suite.Equal(`{"success":false,"message":"The extractor for the given id does not exist."}`, resp.Body)
}

// func (suite *DispatchControllerAcceptanceTest) TestDispatchTrainingJob_Success() {
// 	// given
// 	reducer := "or"
// 	exporter, _ := suite.exporterRepo.Save(suite.Ctx, &model.Exporter{
// 		Name:                 "ExporterJs204",
// 		Description:          "ExporterJs204",
// 		Dimensions:           []int{204, 1},
// 		Type:                 model.JS,
// 		ExportScriptLocation: nil,
// 	})
// 	suite.exportRunRepo.Save(suite.Ctx, &model.ExportRun{
// 		ExporterId: exporter.ID,
// 		Reducer:    reducer,
// 		Start:      time.Now(),
// 		End:        time.Now(),
// 	})
// 	model, _ := suite.modelRepo.Save(suite.Ctx, &model.Model{
// 		Name:        "Model1",
// 		Description: "Model1",
// 		Dims:        []int{204, 1},
// 	})

// 	// when
// 	go suite.testConsumer.Consume(config.EnvTrainQueueName(), 1)
// 	time.Sleep(5 * time.Second)
// 	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/train/%s/run/%s/%s", model.ID, exporter.ID, reducer), "", "")
// 	suite.testConsumer.WaitForMessages(config.EnvTrainQueueName(), 1)
// 	// then
// 	suite.NoError(err)
// 	suite.Equal(201, resp.StatusCode)
// 	suite.Equal(1, len(suite.testConsumer.QueueMessages[config.EnvTrainQueueName()]))
// 	suite.Equal(fmt.Sprintf(`{"functionName":"train_model","args":["%s","%s","%s"]}`, model.ID, exporter.ID, reducer), suite.testConsumer.QueueMessages[config.EnvTrainQueueName()][0])
// }

func (suite *DispatchControllerAcceptanceTest) TestDispatchTrainingJob_ErrorNoRunFound() {
	// given

	exporter, _ := suite.exporterRepo.Save(suite.Ctx, &model.Exporter{
		Name:                 "ExporterJs204",
		Description:          "ExporterJs204",
		Dimensions:           []int{204, 1},
		Type:                 model.JS,
		ExportScriptLocation: nil,
	})
	reducer := "or"
	model, _ := suite.modelRepo.Save(suite.Ctx, &model.Model{
		Name:        "Model1",
		Description: "Model1",
		Dims:        []int{204, 1},
	})
	// when
	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/train/%s/run/%s/%s", model.ID, exporter.ID, reducer), "", "")

	// then
	suite.NoError(err)
	suite.Equal(404, resp.StatusCode)
	suite.Equal(`{"success":false,"message":"The export for the given id and reducer does not exist."}`, resp.Body)
}

func (suite *DispatchControllerAcceptanceTest) TestDispatchTrainingJob_ErrorModelNotFound() {
	// given
	random := "5f5e7e3e3e3e3e3e3e3e3e3e"
	exporter, _ := suite.exporterRepo.Save(suite.Ctx, &model.Exporter{
		Name:                 "ExporterJs204",
		Description:          "ExporterJs204",
		Dimensions:           []int{204, 1},
		Type:                 model.JS,
		ExportScriptLocation: nil,
	})
	reducer := "or"
	suite.exportRunRepo.Save(suite.Ctx, &model.ExportRun{
		ExporterId: exporter.ID,
		Reducer:    reducer,
		Start:      time.Now(),
		End:        time.Now(),
	})
	// when
	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/train/%s/run/%s/%s", random, exporter.ID, reducer), "", "")
	// then
	suite.NoError(err)
	suite.Equal(404, resp.StatusCode)
	suite.Equal(`{"success":false,"message":"The model for the given id does not exist."}`, resp.Body)
}

func (suite *DispatchControllerAcceptanceTest) TestDispatchTrainingJob_ErrorDimensionMismatch() {
	// given

	exporter, _ := suite.exporterRepo.Save(suite.Ctx, &model.Exporter{
		Name:                 "ExporterJs204",
		Description:          "ExporterJs204",
		Dimensions:           []int{204, 2},
		Type:                 model.JS,
		ExportScriptLocation: nil,
	})
	reducer := "or"
	suite.exportRunRepo.Save(suite.Ctx, &model.ExportRun{
		ExporterId: exporter.ID,
		Reducer:    reducer,
		Start:      time.Now(),
		End:        time.Now(),
	})
	model, _ := suite.modelRepo.Save(suite.Ctx, &model.Model{
		Name:        "Model1",
		Description: "Model1",
		Dims:        []int{204, 1},
	})
	// when
	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/train/%s/run/%s/%s", model.ID, exporter.ID, reducer), "", "")
	// then
	suite.NoError(err)
	suite.Equal(400, resp.StatusCode)
	suite.Equal(`{"success":false,"message":"There is a dimension mismatch for the dataset and the model."}`, resp.Body)
}

func (suite *DispatchControllerAcceptanceTest) TestDispatchTrainingJob_ErrorExporterNotFound() {
	// given
	random := "5f5e7e3e3e3e3e3e3e3e3e3e"
	model, _ := suite.modelRepo.Save(suite.Ctx, &model.Model{
		Name:        "Model1",
		Description: "Model1",
		Dims:        []int{204, 1},
	})
	reducer := "or"
	// when
	resp, err := testsupport.Post(fmt.Sprintf("http://localhost:8081/dispatch/train/%s/run/%s/%s", model.ID, random, reducer), "", "")
	// then
	suite.NoError(err)
	suite.Equal(404, resp.StatusCode)
	suite.Equal(`{"success":false,"message":"The extractor for the given id does not exist."}`, resp.Body)
}
