package controller

import (
	"fmt"

	"github.com/Tracking-Detector/td-backend/go/td_common/payload"
	"github.com/Tracking-Detector/td-backend/go/td_common/response"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/gofiber/fiber/v2"
)

type DatasetController struct {
	datasetService service.IDatasetService
}

func NewDatasetController(datasetService service.IDatasetService) *DatasetController {
	return &DatasetController{
		datasetService: datasetService,
	}
}

func (dc *DatasetController) GetAllDatasets(c *fiber.Ctx) error {
	datasets := dc.datasetService.GetAllDatasets()
	return c.Status(fiber.StatusOK).JSON(response.NewSuccessResponse(datasets))
}

func (dc *DatasetController) CreateDataset(c *fiber.Ctx) error {
	datasetPayload := new(payload.CreateDatasetPayload)

	if err := c.BodyParser(datasetPayload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(err.Error()))
	}
	fmt.Println(*datasetPayload)
	dataset, err := dc.datasetService.CreateDataset(c.Context(), datasetPayload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(response.NewSuccessResponse(dataset))
}

func (dc *DatasetController) DeleteDataset(c *fiber.Ctx) error {
	id := c.Params("id")
	err := dc.datasetService.DeleteDataset(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.NewErrorResponse(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(response.NewSuccessResponse(nil))
}

func (dc *DatasetController) RegisterRoutes(app *fiber.App) *fiber.App {
	app.Get("/datasets/health", util.GetHealth)
	app.Delete("/datasets/:id", dc.DeleteDataset)
	app.Get("/datasets", dc.GetAllDatasets)
	app.Post("/datasets", dc.CreateDataset)
	return app
}
