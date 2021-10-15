package route

import (
	"github.com/daniel5u/suisei/presenter/user"
	"github.com/labstack/echo/v4"
)

type PresenterList struct {
	UserPresenter user.Presenter
}

func (presenterList *PresenterList) RegisterRoute(e *echo.Echo) {
	user := e.Group("users")
	user.GET("", presenterList.UserPresenter.Fetch)
	user.GET("/:id", presenterList.UserPresenter.GetByID)
	user.PUT("/:id", presenterList.UserPresenter.Update)
	user.POST("/register", presenterList.UserPresenter.Store)
	user.DELETE("/:id", presenterList.UserPresenter.Delete)
}
