package repository

import (
	"context"
	"database/sql"
	"rest-api/model/entities"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category entities.Category) entities.Category
	Update(ctx context.Context, tx *sql.Tx, category entities.Category) entities.Category
	Delete(ctx context.Context, tx *sql.Tx, category entities.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int ) (entities.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entities.Category
}