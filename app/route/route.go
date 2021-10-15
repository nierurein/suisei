package route

import (
	"github.com/daniel5u/suisei/presenter/category"
	"github.com/daniel5u/suisei/presenter/publisher"
	"github.com/daniel5u/suisei/presenter/user"
	"github.com/labstack/echo/v4"
)

type PresenterList struct {
	UserPresenter      user.Presenter
	CategoryPresenter  category.Presenter
	PublisherPresenter publisher.Presenter
}

func (presenterList *PresenterList) RegisterRoute(e *echo.Echo) {
	user := e.Group("users")
	user.GET("", presenterList.UserPresenter.Fetch)
	user.GET("/:id", presenterList.UserPresenter.GetByID)
	user.PUT("/:id", presenterList.UserPresenter.Update)
	user.POST("/register", presenterList.UserPresenter.Store)
	user.DELETE("/:id", presenterList.UserPresenter.Delete)

	category := e.Group("categories")
	category.GET("", presenterList.CategoryPresenter.Fetch)
	category.GET("/:id", presenterList.CategoryPresenter.GetByID)
	category.POST("/register", presenterList.CategoryPresenter.Store)

	publisher := e.Group("publishers")
	publisher.GET("", presenterList.PublisherPresenter.Fetch)
	publisher.GET("/:id", presenterList.PublisherPresenter.GetByID)
	publisher.POST("/register", presenterList.PublisherPresenter.Store)
}
