package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_middleware "github.com/daniel5u/suisei/app/middleware"
	_route "github.com/daniel5u/suisei/app/route"

	_userPresenter "github.com/daniel5u/suisei/presenter/user"
	_userRepository "github.com/daniel5u/suisei/repository/postgresql/user"
	_userService "github.com/daniel5u/suisei/service/user"

	_categoryPresenter "github.com/daniel5u/suisei/presenter/category"
	_categoryRepository "github.com/daniel5u/suisei/repository/postgresql/category"
	_categoryService "github.com/daniel5u/suisei/service/category"

	_publisherPresenter "github.com/daniel5u/suisei/presenter/publisher"
	_publisherRepository "github.com/daniel5u/suisei/repository/postgresql/publisher"
	_publisherService "github.com/daniel5u/suisei/service/publisher"

	_authorPresenter "github.com/daniel5u/suisei/presenter/author"
	_authorRepository "github.com/daniel5u/suisei/repository/postgresql/author"
	_authorService "github.com/daniel5u/suisei/service/author"

	_transactionPresenter "github.com/daniel5u/suisei/presenter/transaction"
	_transactionRepository "github.com/daniel5u/suisei/repository/postgresql/transaction"
	_transactionService "github.com/daniel5u/suisei/service/transaction"

	_bookPresenter "github.com/daniel5u/suisei/presenter/book"
	_bookRepository "github.com/daniel5u/suisei/repository/postgresql/book"
	_bookService "github.com/daniel5u/suisei/service/book"

	_bookauthorRepository "github.com/daniel5u/suisei/repository/postgresql/bookauthor"

	_booktransactionRepository "github.com/daniel5u/suisei/repository/postgresql/booktransaction"
)

func initConfig() {
	viper.SetConfigFile(`app/config/config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetString("database.port"),
		viper.GetString("database.sslmode"),
		viper.GetString("database.TimeZone"),
	)

	DB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(
		&_userRepository.User{},
		&_categoryRepository.Category{},
		&_publisherRepository.Publisher{},
		&_authorRepository.Author{},
		&_transactionRepository.Transaction{},
		&_bookRepository.Book{},
		&_bookauthorRepository.BookAuthor{},
		&_booktransactionRepository.BookTransaction{},
	)

	return DB
}

func main() {
	initConfig()
	db := initDB()
	e := echo.New()

	_middleware.UseTrailingSlash(e)
	_middleware.UseLogger(e)

	userRepository := _userRepository.NewRepository(db)
	userService := _userService.NewService(userRepository)
	userPresenter := _userPresenter.NewPresenter(userService)

	categoryRepository := _categoryRepository.NewRepository(db)
	categoryService := _categoryService.NewService(categoryRepository)
	categoryPresenter := _categoryPresenter.NewPresenter(categoryService)

	publisherRepository := _publisherRepository.NewRepository(db)
	publisherService := _publisherService.NewService(publisherRepository)
	publisherPresenter := _publisherPresenter.NewPresenter(publisherService)

	authorRepository := _authorRepository.NewRepository(db)
	authorService := _authorService.NewService(authorRepository)
	authorPresenter := _authorPresenter.NewPresenter(authorService)

	transactionRepository := _transactionRepository.NewRepository(db)
	transactionService := _transactionService.NewService(transactionRepository)
	transactionPresenter := _transactionPresenter.NewPresenter(transactionService)

	bookRepository := _bookRepository.NewRepository(db)
	bookService := _bookService.NewService(bookRepository)
	bookPresenter := _bookPresenter.NewPresenter(bookService)

	routes := _route.PresenterList{
		UserPresenter:        *userPresenter,
		CategoryPresenter:    *categoryPresenter,
		PublisherPresenter:   *publisherPresenter,
		AuthorPresenter:      *authorPresenter,
		TransactionPresenter: *transactionPresenter,
		BookPresenter:        *bookPresenter,
	}
	routes.RegisterRoute(e)

	log.Fatal(e.Start(viper.GetString("server.port")))
}
