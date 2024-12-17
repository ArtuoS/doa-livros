package entity

type DonatedBook struct {
	Id         int64  `db:"id"`
	FromUserId int64  `db:"from_user_id"`
	ToUserId   int64  `db:"to_user_id"`
	BookId     int64  `db:"book_id"`
	ToUserName string `db:"to_user_name"`
}
