package controllers

import (
	"github.com/labstack/echo"
	"github.com/uasouz/tms/lib/domain/models"
	"github.com/uasouz/tms/lib/repositories"
	"github.com/uasouz/tms/lib/util"
	"strings"
)

type TMSApiController struct {
	emailTemplateRepository repositories.EmailTemplatesRepository
	destinationRepository   repositories.DestinationRepository
}

func (controller *TMSApiController) SendMessage(c echo.Context) error {
	request := new(models.MessageData)

	if err := c.Bind(request); err != nil {
		return err
	}

	//Send Message on each informed platform of each destination
	//Code can be improved
	for _,destination := range request.To {
		parseDestination(destination)
	}

	return c.JSON(util.SuccessResponse(""))
}

func parseDestination(destinationURN string) *models.Destination {
	splitedParts := strings.Split(destinationURN,"://")
	return &models.Destination{
		Platform:   splitedParts[0],
		Identifier: splitedParts[1],
	}
}