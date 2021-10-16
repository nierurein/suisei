package user

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/user"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service user.Service
}

func NewPresenter(userService user.Service) *Presenter {
	return &Presenter{
		service: userService,
	}
}

func (userPresenter *Presenter) Fetch(echoContext echo.Context) error {
	userDomains, err := userPresenter.service.Fetch()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var userResponses []Response
	for _, userDomain := range userDomains {
		userResponses = append(userResponses, domainToResponse(userDomain))
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": userResponses,
	})
}

func (userPresenter *Presenter) GetByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	userDomain, err := userPresenter.service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(userDomain),
	})
}

func (userPresenter *Presenter) Update(echoContext echo.Context) error {
	var userRequest Request

	err := echoContext.Bind(&userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, _ := strconv.Atoi(echoContext.Param("id"))

	userDomain := requestToDomain(userRequest)

	userDomainAfter, err := userPresenter.service.Update(userDomain, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(userDomainAfter),
	})
}

func (userPresenter *Presenter) Store(echoContext echo.Context) error {
	var userRequest Request

	err := echoContext.Bind(&userRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userDomain := requestToDomain(userRequest)

	userDomainAfter, err := userPresenter.service.Store(userDomain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(userDomainAfter),
	})
}

func (userPresenter *Presenter) Delete(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	err := userPresenter.service.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
