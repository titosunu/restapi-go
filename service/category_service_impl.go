package service

import (
	"context"
	"database/sql"
	"rest-api/helper"
	"rest-api/model/entities"
	"rest-api/model/web"
	"rest-api/repository"
)

type CategoryServiceImpl struct {
	repository 	repository.CategoryRepository
	DB 					*sql.DB
}

func(service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category := entities.Category{
		Name: request.Name,
	}

	category = service.repository.Create(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func(service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.repository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func(service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.repository.Delete(ctx, tx, category)
}

func(service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	category, err := service.repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func(service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	categories := service.repository.FindAll(ctx, tx)

	var categoryResponses []web.CategoryResponse
	for _,  category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}