package domain

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
