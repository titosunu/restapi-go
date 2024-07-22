package helper

import (
	"rest-api/model/entities"
	"rest-api/model/web"
)

func ToCategoryResponse(category entities.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}