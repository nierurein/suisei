package category

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/category"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service category.Service
}

func NewPresenter(categoryService category.Service) *Presenter {
	return &Presenter{
		service: categoryService,
	}
}

func (categoryPresenter *Presenter) Fetch(echoContext echo.Context) error {
	categoryDomains, err := categoryPresenter.service.Fetch()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var categoryResponses []Response
	for _, categoryDomain := range categoryDomains {
		categoryResponses = append(categoryResponses, domainToResponse(categoryDomain))
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": categoryResponses,
	})
}

func (categoryPresenter *Presenter) GetByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	categoryDomain, err := categoryPresenter.service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(categoryDomain),
	})
}

func (categoryPresenter *Presenter) Store(echoContext echo.Context) error {
	var categoryRequest Request

	err := echoContext.Bind(&categoryRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	categoryDomain := requestToDomain(categoryRequest)

	categoryDomainAfter, err := categoryPresenter.service.Store(categoryDomain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(categoryDomainAfter),
	})
}
