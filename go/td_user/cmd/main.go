package main

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/repository"
	"github.com/Tracking-Detector/td-backend/go/td_common/service"
	"github.com/Tracking-Detector/td-backend/go/td_common/util"
	"github.com/Tracking-Detector/td-backend/go/td_user/controller"
)

func main() {
	ctx := context.Background()
	userRepo := repository.NewMongoUserRepository(config.GetDatabase(config.ConnectDB(ctx)))
	encryptionService := service.NewEncryptionService()
	userService := service.NewUserService(userRepo, encryptionService)
	userController := controller.NewUserController(userService)

	app := util.DefaultFiberApp()

	userController.RegisterRoutes(app).Listen(":8081")
}
