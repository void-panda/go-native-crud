package postmodel

import (
	"go-native-crud/config"
	"go-native-crud/src/entities"
)

func GetAll() []entities.Post {
	rows, err := config.DB.Query(
		`SELECT 
			p.id,
			p.title,
			p.content,
			c.name
		FROM posts p
		JOIN categories c ON p.category_id = c.id
		ORDER BY p.created_at`,
	)
	if err != nil {
		panic(err)
	}

	var posts []entities.Post

	for rows.Next() {
		var post entities.Post
		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Category.Name,
		)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	defer rows.Close()

	return posts
}

func Create(post entities.Post) bool {
	result, err := config.DB.Exec(
		`INSERT INTO posts (title, category_id, content)
		VALUE (?, ?, ?)`,
		post.Title, post.Category.Id, post.Content,
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

func GetById(id int) entities.Post {
	rows := config.DB.QueryRow(
		`SELECT
			p.id,
			p.title,
			p.content,
			p.category_id
		FROM posts p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = ?`,
		id,
	)

	var post entities.Post
	err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category.Id)
	if err != nil {
		panic(err)
	}

	return post
}

func Update(id int, post entities.Post) bool {
	result, err := config.DB.Exec(
		`UPDATE posts set 
			title = ?, 
			content = ?, 
			category_id = ?, 
			updated_at = ?
		WHERE id = ?`,
		post.Title,
		post.Content,
		post.Category.Id,
		post.Updated_at,
		id,
	)
	if err != nil {
		panic(err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowAffected > 0
}

func Delete(id int) error {
	query, err := config.DB.Exec(`DELETE from posts WHERE id = ?`, id)
	if err != nil {
		return err
	}

	_, err = query.RowsAffected()

	return err
}
