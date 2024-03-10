package categorymodel

import (
	"Golang_CRUD_Native/config"
	"Golang_CRUD_Native/entities"
)

func Getall() []entities.Category {
	rows, err := config.DB.Query("SELECT * FROM categories")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)

	}
	return categories
}