package author

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/author"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service author.Service
}

func NewPresenter(authorService author.Service) *Presenter {
	return &Presenter{
		service: authorService,
	}
}

func (authorPresenter *Presenter) Fetch(echoContext echo.Context) error {
	authorDomains, err := authorPresenter.service.Fetch()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var authorResponses []Response
	for _, authorDomain := range authorDomains {
		authorResponses = append(authorResponses, domainToResponse(authorDomain))
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": authorResponses,
	})
}

func (authorPresenter *Presenter) GetByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	authorDomain, err := authorPresenter.service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(authorDomain),
	})
}

func (authorPresenter *Presenter) Store(echoContext echo.Context) error {
	var authorRequest Request

	err := echoContext.Bind(&authorRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	authorDomain := requestToDomain(authorRequest)

	authorDomainAfter, err := authorPresenter.service.Store(authorDomain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(authorDomainAfter),
	})
}
