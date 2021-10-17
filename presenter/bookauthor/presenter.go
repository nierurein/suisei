package bookauthor

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/bookauthor"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service bookauthor.Service
}

func NewPresenter(bookauthorService bookauthor.Service) *Presenter {
	return &Presenter{
		service: bookauthorService,
	}
}

func (bookauthorPresenter *Presenter) StoreBatch(echoContext echo.Context) error {
	var bookauthorRequest Request
	var bookauthorDomains []bookauthor.Domain

	id, _ := strconv.Atoi(echoContext.Param("id"))

	err := echoContext.Bind(&bookauthorRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, bookauthorRequestItem := range bookauthorRequest.Items {
		bookauthorRequestItem.BookID = id
		bookauthorDomains = append(bookauthorDomains, requestItemToDomain(bookauthorRequestItem))
	}

	err = bookauthorPresenter.service.StoreBatch(bookauthorDomains)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
