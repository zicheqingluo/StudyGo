package model

import (
	"time"
)

//留言
type Leave struct{
	Id	int64	`db:"id"`
	Username string `db:"username"`
	Content	string	`db:"content"`
	CreateTime	time.Time `db:"create_time"`
	Email	string	`db:"email"`
	UpdateTime	time.Time	`db:"update_time"`
}