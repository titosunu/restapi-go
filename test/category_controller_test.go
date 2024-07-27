package test

import (
	"database/sql"
	"net/http"
	"rest-api/app"
	"rest-api/controller"
	"rest-api/helper"
	"rest-api/middleware"
	"rest-api/repository"
	"rest-api/service"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/rest_api_test")
	helper.PanicIfError(err)
	
	db.SetConnMaxIdleTime(5)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setUpRouter() http.Handler {
	validate := validator.New()
	db := setupTestDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	
}

func TestCreateCategoryFailed(t *testing.T) {
	
}

func TestUpdateCategorySuccess(t *testing.T) {
	
}

func TestUpdateCategoryFailed(t *testing.T) {
	
}

func TestGetCategorySuccess(t *testing.T) {
	
}

func TestGetCategoryFailed(t *testing.T) {
	
}

func TestDeleteCategorySuccess(t *testing.T) {
	
}

func TestDeleteCategoryFailed(t *testing.T) {
	
}

func TestListCategoriesSuccess(t *testing.T) {
	
}

func TestUnauthorized(t *testing.T) {
	
}