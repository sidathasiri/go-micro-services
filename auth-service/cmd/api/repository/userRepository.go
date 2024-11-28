package repository

import (
	"auth-service/cmd/api/db"
	"auth-service/cmd/api/models"
	"database/sql"
	"log"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		connection: db.GetConnection(),
	}
}

func (userRepo *UserRepository) SaveUser(user models.User) models.User{
	sqlStatement := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	statement, err := userRepo.connection.Prepare(sqlStatement)
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
	db := userRepo.connection
	user := models.User{} 
	query := "SELECT id, email, password FROM users where email = $1"
	err := db.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return models.User{}
	}

	defer db.Close()

	return user
}
