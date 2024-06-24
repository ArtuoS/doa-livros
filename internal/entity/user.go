package entity

type User struct {
	Id           int64  `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	Books        []Book
	DonatedBooks []DonatedBook
}
