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

	routes := _route.PresenterList{
		UserPresenter:     *userPresenter,
		CategoryPresenter: *categoryPresenter,
	}
	routes.RegisterRoute(e)

	log.Fatal(e.Start(viper.GetString("server.port")))
}
