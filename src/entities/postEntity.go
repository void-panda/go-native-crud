package entities

import "time"

type Post struct {
	Id             uint
	Title, Content string
	Category       Category
	Created_at     time.Time
	Updated_at     time.Time
}
