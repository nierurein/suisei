package route

import (
	"github.com/daniel5u/suisei/presenter/author"
	"github.com/daniel5u/suisei/presenter/book"
	"github.com/daniel5u/suisei/presenter/bookauthor"
	"github.com/daniel5u/suisei/presenter/booktransaction"
	"github.com/daniel5u/suisei/presenter/category"
	"github.com/daniel5u/suisei/presenter/publisher"
	"github.com/daniel5u/suisei/presenter/transaction"
	"github.com/daniel5u/suisei/presenter/user"
	"github.com/labstack/echo/v4"
)

type PresenterList struct {
	UserPresenter            user.Presenter
	CategoryPresenter        category.Presenter
	PublisherPresenter       publisher.Presenter
	AuthorPresenter          author.Presenter
	TransactionPresenter     transaction.Presenter
	BookPresenter            book.Presenter
	BooktransactionPresenter booktransaction.Presenter
	BookauthorPresenter      bookauthor.Presenter
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

	author := e.Group("authors")
	author.GET("", presenterList.AuthorPresenter.Fetch)
	author.GET("/:id", presenterList.AuthorPresenter.GetByID)
	author.POST("/register", presenterList.AuthorPresenter.Store)

	transaction := e.Group("transactions")
	transaction.GET("", presenterList.TransactionPresenter.Fetch)
	transaction.GET("/:id", presenterList.TransactionPresenter.GetByID)
	transaction.PUT("/:id", presenterList.TransactionPresenter.Update)
	transaction.POST("/register", presenterList.TransactionPresenter.Store)
	transaction.POST("/:id/items", presenterList.BooktransactionPresenter.StoreBatch)

	book := e.Group("books")
	book.GET("", presenterList.BookPresenter.Fetch)
	book.GET("/:id", presenterList.BookPresenter.GetByID)
	book.PUT("/:id", presenterList.BookPresenter.Update)
	book.POST("/register", presenterList.BookPresenter.Store)
	book.POST("/:id/authors", presenterList.BookauthorPresenter.StoreBatch)
	book.DELETE("/:id", presenterList.BookPresenter.Delete)
}
