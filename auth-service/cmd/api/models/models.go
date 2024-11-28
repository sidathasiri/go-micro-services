package models

type User struct {
	Id  string `db:"id"`
	Email string `db:"email"`
	Password string `db:"password"`
   }