package entity

type Book struct {
	Id       int64  `db:"id"`
	Title    string `db:"title"`
	Author   string `db:"author"`
	UserId   int64  `db:"user_id"`
	Donating bool   `db:"donating"`
}
