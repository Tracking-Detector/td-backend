package controller

import (
	"github.com/Tracking-Detector/td-backend/go/td_common/response"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/gofiber/fiber/v2"
)

type ExportController struct {
	extractorService service.IExporterService
}

func NewExportController(extractorService service.IExporterService) *ExportController {
	return &ExportController{
		extractorService: extractorService,
	}
}

func (con *ExportController) GetAllExporter(c *fiber.Ctx) error {
	extractors, err := con.extractorService.GetAllExporter(c.Context())
	if err != nil {
		errorResponse := response.NewErrorResponse(err.Error())
		return c.Status(500).JSON(errorResponse)
	}
	successResponse := response.NewSuccessResponse(extractors)
	return c.Status(200).JSON(successResponse)
}

func (con *ExportController) RegisterRoutes(app *fiber.App) *fiber.App {

	app.Get("/export/health", util.GetHealth)
	app.Get("/export", con.GetAllExporter)
	return app
}
