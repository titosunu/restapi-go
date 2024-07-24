package main

import (
	"rest-api/app"
	"rest-api/controller"
	"rest-api/repository"
	"rest-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/categories", categoryController.FindAll)
	router.GET("/categories:categoryId", categoryController.FindById)
	router.POST("/categories", categoryController.Create)
	router.PUT("/categories:categoryId", categoryController.Update)
	router.DELETE("/categories:categoryId", categoryController.Delete)

}