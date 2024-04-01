package models

type User struct {
	ID       int
	Username string
	Email    string
	Age      int
}

var First = User{
	ID:       1,
	Username: "john_doe",
	Email:    "john@example.com",
	Age:      30,
}