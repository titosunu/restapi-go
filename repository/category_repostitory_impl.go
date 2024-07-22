package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest-api/helper"
	"rest-api/model/entities"
)

type CategoryRepositoryImpl struct {
}

func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category entities.Category) entities.Category {
	sql := "INSERT INTO customer(name) values (?)"

	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context,tx *sql.Tx, category entities.Category) entities.Category {
	sql := "UPDATE category SET name = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func(c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category entities.Category) {
	sql := "DELETE FROM category WHERE id = ?"

	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

func(c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (entities.Category, error) {
	sql := "SELECT id, name FROM category WHERE id = ?"

	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)

	category := entities.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	}

	return category, errors.New("category is not found")
	
}

func(c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entities.Category {
	sql := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	var categories []entities.Category
	for rows.Next() {
		category := entities.Category{}
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}
	
	return categories
}