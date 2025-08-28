package categorymodel

import (
	"go-native-crud/config"
	"go-native-crud/src/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT id, name FROM categories ORDER BY created_at`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	// parse setiap baris data kedalam struct
	for rows.Next() {
		var category entities.Category
		err = rows.Scan(&category.Id, &category.Name)
		if err != nil {
			panic(err)
		}
		// tampung semua data dalam 1 slice
		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(
		`INSERT INTO categories (name, created_at, updated_at)
		VALUE (?, ?, ?)`,
		category.Name, category.Created_at, category.Updated_at,
	)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0

}

func GetById(id int) entities.Category {
	row := config.DB.QueryRow(
		`SELECT id, name FROM categories WHERE id = ?`,
		id,
	)

	var category entities.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		panic(err)
	}

	return category
}

func Update(id int, data entities.Category) bool {
	query, err := config.DB.Exec(
		`UPDATE categories SET name = ?, updated_at = ? WHERE id = ?`,
		data.Name,
		data.Updated_at,
		id,
	)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	query, err := config.DB.Exec(
		`DELETE FROM categories WHERE id = ?`,
		id,
	)
	if err != nil {
		panic(err)
	}

	_, err = query.RowsAffected()

	return err
}
