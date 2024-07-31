package main

import (
	"net/http"

	"github.com/MrBista/crud-go/app"
	"github.com/MrBista/crud-go/controller"
	"github.com/MrBista/crud-go/helper"
	"github.com/MrBista/crud-go/middleware"
	"github.com/MrBista/crud-go/repository"
	"github.com/MrBista/crud-go/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// pooling db

	db := app.NewDB()

	validate := validator.New()

	// repository -> untuk db logic
	categoryRepository := repository.NewCategoryRepository()

	// service -> untuk bisnis logic
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// controller -> untuk menampung request http
	categoryController := controller.NewCategoryController(categoryService)

	// app router
	router := app.NewRouter(categoryController)

	server := http.Server{
		// Addr:    "localhost:8000",
		Addr: "127.0.0.1:3000",

		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
