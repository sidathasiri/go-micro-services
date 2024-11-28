package repository

import (
	"auth-service/cmd/api/models"
	"database/sql"
	"log"
)

type UserRepository struct {
	Connection *sql.DB
}

func (userRepo *UserRepository) SaveUser(user models.User) models.User{
	sqlStatement := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	statement, err := userRepo.Connection.Prepare(sqlStatement)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Id, user.Email, user.Password); err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	log.Printf("User %s created\n", user.Email)

	return user
}

func (userRepo *UserRepository) FindUserByEmail(email string) models.User {
	db := userRepo.Connection
	user := models.User{} 
	query := "SELECT id, email, password FROM users where email = $1"
	err := db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return models.User{}
	}

	return user
}

func (userRepo *UserRepository) UpdateUser (user models.User) models.User {
	sqlStatement := "UPDATE users SET email = $1, password = $2 WHERE id = $3"
	statement, err := userRepo.Connection.Prepare(sqlStatement)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Email, user.Password, user.Id); err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	log.Printf("User %s updated\n", user.Email)

	return user
}