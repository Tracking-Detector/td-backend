package controller

import (
	"net/http"

	"github.com/Tracking-Detector/td-backend/go/td_common/payload"
	"github.com/Tracking-Detector/td-backend/go/td_common/response"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_user/representation"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	app         *fiber.App
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	users, err := uc.userService.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrorResponse(err.Error()))
	}
	return c.Status(http.StatusOK).JSON(representation.ConvertUserDatasToUserDataRepresentations(users))
}

func (uc *UserController) CreateApiUser(c *fiber.Ctx) error {
	var createUserData payload.CreateUserData
	if err := c.BodyParser(&createUserData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrorResponse(err.Error()))
	}
	key, err := uc.userService.CreateApiUser(c.Context(), createUserData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrorResponse(err.Error()))
	}
	return c.Status(http.StatusCreated).JSON("User created with key '" + key + "'.")
}

func (uc *UserController) DeleteUserByID(c *fiber.Ctx) error {
	userId := c.Params("id")

	err := uc.userService.DeleteUserByID(c.Context(), userId)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.NewErrorResponse(err.Error()))
	}

	return c.Status(http.StatusOK).JSON(response.NewSuccessResponse("Deleted user successful."))
}

func (uc *UserController) RegisterRoutes(app *fiber.App) *fiber.App {
	app.Get("/users/health", util.GetHealth)
	app.Get("/users", uc.GetUsers)
	app.Post("/users", uc.CreateApiUser)
	app.Delete("/users/:Id", uc.DeleteUserByID)
	return app
}
