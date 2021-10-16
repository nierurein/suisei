package book

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/book"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service book.Service
}

func NewPresenter(bookService book.Service) *Presenter {
	return &Presenter{
		service: bookService,
	}
}

func (bookPresenter *Presenter) Fetch(echoContext echo.Context) error {
	bookDomains, err := bookPresenter.service.Fetch()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var bookResponses []Response
	for _, bookDomain := range bookDomains {
		bookResponses = append(bookResponses, domainToResponse(bookDomain))
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": bookResponses,
	})
}

func (bookPresenter *Presenter) GetByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	bookDomain, err := bookPresenter.service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(bookDomain),
	})
}

func (bookPresenter *Presenter) Update(echoContext echo.Context) error {
	var bookRequest Request

	err := echoContext.Bind(&bookRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, _ := strconv.Atoi(echoContext.Param("id"))

	bookDomain := requestToDomain(bookRequest)

	bookDomainAfter, err := bookPresenter.service.Update(bookDomain, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(bookDomainAfter),
	})
}

func (bookPresenter *Presenter) Store(echoContext echo.Context) error {
	var bookRequest Request

	err := echoContext.Bind(&bookRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bookDomain := requestToDomain(bookRequest)

	bookDomainAfter, err := bookPresenter.service.Store(bookDomain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(bookDomainAfter),
	})
}

func (bookPresenter *Presenter) Delete(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	err := bookPresenter.service.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
