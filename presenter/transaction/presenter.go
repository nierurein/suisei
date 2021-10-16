package transaction

import (
	"net/http"
	"strconv"

	"github.com/daniel5u/suisei/domain/transaction"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	service transaction.Service
}

func NewPresenter(transactionService transaction.Service) *Presenter {
	return &Presenter{
		service: transactionService,
	}
}

func (transactionPresenter *Presenter) Fetch(echoContext echo.Context) error {
	transactionDomains, err := transactionPresenter.service.Fetch()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var transactionResponses []Response
	for _, transactionDomain := range transactionDomains {
		transactionResponses = append(transactionResponses, domainToResponse(transactionDomain))
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": transactionResponses,
	})
}

func (transactionPresenter *Presenter) GetByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	transactionDomain, err := transactionPresenter.service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(transactionDomain),
	})
}

func (transactionPresenter *Presenter) Update(echoContext echo.Context) error {
	var transactionRequest Request

	echoContext.Bind(&transactionRequest)
	id, _ := strconv.Atoi(echoContext.Param("id"))

	transactionDomain := requestToDomain(transactionRequest)

	transactionDomainAfter, err := transactionPresenter.service.Update(transactionDomain, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(transactionDomainAfter),
	})
}

func (transactionPresenter *Presenter) Store(echoContext echo.Context) error {
	var transactionRequest Request

	echoContext.Bind(&transactionRequest)

	transactionDomain := requestToDomain(transactionRequest)

	transactionDomainAfter, err := transactionPresenter.service.Store(transactionDomain)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"result": domainToResponse(transactionDomainAfter),
	})
}
