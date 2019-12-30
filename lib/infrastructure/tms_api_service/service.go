package tms_api_service

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/uasouz/tms/config"
	"github.com/uasouz/tms/lib/infrastructure"
	"github.com/uasouz/tms/lib/interface_adapters/controllers"
	"github.com/uasouz/tms/lib/util"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type TMSApiService struct {
	infrastructure.IService
	Echo       *echo.Echo
	Config     config.Config
	controller controllers.TMSApiController
	shutDown   chan os.Signal
}

func (service *TMSApiService) Start() {
	service.Echo = echo.New()

	service.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.POST, echo.GET, echo.OPTIONS, echo.PUT, echo.DELETE, echo.HEAD},
	}))

	go func() {
		service.Echo.Logger.Fatal(service.Echo.Start(service.Config.TesAPIPort))
	}()

	service.shutDown = make(chan os.Signal)
	<-service.shutDown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := service.Echo.Shutdown(ctx); err != nil {
		service.Echo.Logger.Fatal(err)
	}

}

func (service *TMSApiService) Stop() {
	signal.Notify(service.shutDown, os.Interrupt)
}

func (service *TMSApiService) initializeRoutes() {

	service.Echo.Use(middleware.Logger())

	service.Echo.GET("/", Index)
	service.Echo.POST("/message", service.controller.SendMessage)
}

func Index(ctx echo.Context) error {
	var response = util.BaseResponseCompat{
		BaseResponse: util.BaseResponse{
			Status:  false,
			Data:    map[string]interface{}{},
			Message: "",
		},
		Code: 401,
	}
	response.Success(map[string]interface{}{"Out": "Nada para ver aqui!!"})
	return ctx.JSON(http.StatusOK, response)
}
