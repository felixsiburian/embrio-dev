// main function here and Start func is for calling all Init Config and start server

package main

import (
	"embrio-dev/lib/migration"
	econfig "embrio-dev/service/config"
	"embrio-dev/service/delivery/router"
	"embrio-dev/service/repository"
	"embrio-dev/service/tools"
	"embrio-dev/service/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	app := econfig.Config{}
	econfig.CatchEror(app.InitEnv())
	econfig.CatchEror(migration.InitTable())
	Start()
}

func Start() {
	e := echo.New()

	toolRepo := tools.NewToolRepository()
	tokenRepo := repository.NewTokenRepository(toolRepo)
	nasabahRepo := repository.NewNasabahRepository(toolRepo, tokenRepo)
	nasabahCase := usecase.NewNasabahUsecase(nasabahRepo, toolRepo)
	tokenCase := usecase.NewTokenUsecase(tokenRepo)
	router.NewRouter(e, nasabahCase, tokenCase)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s%s%v", os.Getenv("APP_HOST"), ":", os.Getenv("APP_PORT"))))
}
