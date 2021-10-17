package publisher

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/publisher"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service publisher.Service
}

func NewPresenter(publisherService publisher.Service) *Presenter {
	return &Presenter{
		service: publisherService,
	}
}

func (publisherPresenter *Presenter) Fetch(echoContext echo.Context) error {
	publisherDomains, err := publisherPresenter.service.Fetch()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var publisherResponses []Response
	for _, publisherDomain := range publisherDomains {
		publisherResponses = append(publisherResponses, domainToResponse(publisherDomain))
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": publisherResponses,
	})
}

func (publisherPresenter *Presenter) GetByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	publisherDomain, err := publisherPresenter.service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(publisherDomain),
	})
}

func (publisherPresenter *Presenter) Store(echoContext echo.Context) error {
	var publisherRequest Request

	err := echoContext.Bind(&publisherRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	publisherDomain := requestToDomain(publisherRequest)

	publisherDomainAfter, err := publisherPresenter.service.Store(publisherDomain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(publisherDomainAfter),
	})
}
