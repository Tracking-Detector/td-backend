package acceptance

import (
	"encoding/json"
	"fmt"
	"net/http"

	"testing"
	"time"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/test/testsupport"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_dataset/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/suite"
)

func TestDatasetControllerAcceptance(t *testing.T) {
	suite.Run(t, &DatasetControllerAcceptanceTest{})
}

type DatasetControllerAcceptanceTest struct {
	testsupport.AcceptanceTest
	suite.Suite
	app               *fiber.App
	requestRepo       model.RequestRepository
	requestService    service.IRequestService
	datasetRepo       model.DatasetRepository
	datasetService    service.IDatasetService
	datasetController *controller.DatasetController
}

func (suite *DatasetControllerAcceptanceTest) SetupSuite() {
	suite.SetupIntegration()
	suite.requestRepo = repository.NewMongoRequestRepository(config.GetDatabase(config.ConnectDB(suite.AcceptanceTest.Ctx)))
	suite.requestService = service.NewRequestService(suite.requestRepo)
	suite.datasetRepo = repository.NewMongoDatasetRepository(config.GetDatabase(config.ConnectDB(suite.AcceptanceTest.Ctx)))
	suite.datasetService = service.NewDatasetService(suite.datasetRepo, suite.requestRepo)
	suite.datasetController = controller.NewDatasetController(suite.datasetService)
	go func() {
		suite.app = util.DefaultFiberApp()
		suite.datasetController.RegisterRoutes(suite.app).Listen(":8081")
	}()
	time.Sleep(5 * time.Second)
}

func (suite *DatasetControllerAcceptanceTest) SetupTest() {
	suite.datasetRepo.DeleteAll(suite.AcceptanceTest.Ctx)
}

func (suite *DatasetControllerAcceptanceTest) TearDownSuite() {
	suite.app.Shutdown()
	suite.TeardownIntegration()
}

func (suite *DatasetControllerAcceptanceTest) Test_FindByID() {
	// given
	dataset := &model.Dataset{
		Name:        "test",
		Description: "test",
		Label:       "test",
	}
	dataset, err := suite.datasetRepo.Save(suite.AcceptanceTest.Ctx, dataset)
	// when
	newDataSet, err := suite.datasetRepo.FindByID(suite.AcceptanceTest.Ctx, dataset.ID)
	// then
	fmt.Println(newDataSet, err)

}

func (suite *DatasetControllerAcceptanceTest) TestHealth_Success() {
	// given

	// when
	resp, err := testsupport.Get("http://localhost:8081/datasets/health")

	// then
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
	suite.Equal(`{"success":true,"data":"System is running correct."}`, resp.Body)
}

func (suite *DatasetControllerAcceptanceTest) TestCreateDataset() {
	// given
	datasetPayload := &model.Dataset{
		Name:        "test",
		Description: "test",
		Label:       "test",
	}
	body, _ := json.Marshal(datasetPayload)
	// when
	resp, err := testsupport.Post("http://localhost:8081/datasets", string(body), "application/json")
	// then
	count, _ := suite.datasetRepo.Count(suite.AcceptanceTest.Ctx)
	dataset, _ := suite.datasetRepo.FindByLabel(suite.AcceptanceTest.Ctx, "test")
	suite.Equal(int64(1), count)
	suite.Equal("test", dataset.Label)
	suite.Equal("test", dataset.Name)
	suite.Equal("test", dataset.Description)
	suite.Equal(http.StatusCreated, resp.StatusCode)
	suite.NoError(err)
}

func (suite *DatasetControllerAcceptanceTest) TestGetAllDatasets() {
	// given
	datasetPayload := &model.Dataset{
		Name:        "test",
		Description: "test",
		Label:       "test",
	}
	body, _ := json.Marshal(datasetPayload)
	testsupport.Post("http://localhost:8081/datasets", string(body), "application/json")
	// when
	resp, err := testsupport.Get("http://localhost:8081/datasets")

	// then
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}
