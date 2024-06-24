package entity

type DonatedBook struct {
	FromUserId int64 `db:"from_user_id"`
	ToUserId   int64 `db:"to_user_id"`
	BookId     int64 `db:"book_id"`
}
