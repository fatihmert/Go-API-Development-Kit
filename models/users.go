package models

type User struct {
	ID       uint   `db:"id"`
	Mail     string `db:"mail"`
	Password string `db:"password"`
	Token    string `db:"token"`
}
