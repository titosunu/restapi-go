package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"rest-api/app"
	"rest-api/controller"
	"rest-api/helper"
	"rest-api/middleware"
	"rest-api/model/entities"
	"rest-api/repository"
	"rest-api/service"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
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

func setUpRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB)  {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setUpRouter(db)
	
	requestBody := strings.NewReader(`{"name": "Jacket"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
	assert.Equal(t, "Jacket", responBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setUpRouter(db)
	
	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 400, int(responBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entities.Category{
		Name: "Jacket",
	})
	tx.Commit()

	router := setUpRouter(db)
	
	requestBody := strings.NewReader(`{"name": "Jacket"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
	assert.Equal(t, category.Id, int(responBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Jacket", responBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entities.Category{
		Name: "Jacket",
	})
	tx.Commit()

	router := setUpRouter(db)
	
	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/categories/" + strconv.Itoa(category.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 400, int(responBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entities.Category{
		Name: "Jacket",
	})
	tx.Commit()

	router := setUpRouter(db)
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
	assert.Equal(t, category.Id, int(responBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setUpRouter(db)
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories/404", nil)
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 404, int(responBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entities.Category{
		Name: "Jacket",
	})
	tx.Commit()

	router := setUpRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/categories/" + strconv.Itoa(category.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	router := setUpRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/categories/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 404, int(responBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responBody["status"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(), tx, entities.Category{
		Name: "Jacket",
	})
	tx.Commit()

	router := setUpRouter(db)
	
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/categories", nil)
	request.Header.Add("X-API-Key", "SECRET")

	recoder := httptest.NewRecorder()

	router.ServeHTTP(recoder, request)

	response := recoder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responBody map[string]interface{}
	json.Unmarshal(body, &responBody)
	
	assert.Equal(t, 200, int(responBody["code"].(float64)))
	assert.Equal(t, "OK", responBody["status"])

	var categories = responBody["data"].([]interface{})

	categoriesResponse := categories[0].(map[string]interface{})

	assert.Equal(t, category.Id, int(categoriesResponse["id"].(float64)))
	assert.Equal(t, category.Name, categoriesResponse["name"])
}

func TestUnauthorized(t *testing.T) {
	
}