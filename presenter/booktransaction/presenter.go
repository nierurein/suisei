package booktransaction

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/booktransaction"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service booktransaction.Service
}

func NewPresenter(booktransactionService booktransaction.Service) *Presenter {
	return &Presenter{
		service: booktransactionService,
	}
}

func (booktransactionPresenter *Presenter) StoreBatch(echoContext echo.Context) error {
	var booktransactionRequest Request
	var booktransactionDomains []booktransaction.Domain

	id, _ := strconv.Atoi(echoContext.Param("id"))

	err := echoContext.Bind(&booktransactionRequest)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, booktransactionRequestItem := range booktransactionRequest.Items {
		booktransactionDomains = append(booktransactionDomains, requestItemToDomain(booktransactionRequestItem))
	}

	err = booktransactionPresenter.service.StoreBatch(booktransactionDomains, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
