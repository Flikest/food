package storage

import (
	"net/http"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID          int    `json:"id"`
	Use         string `json:"use"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}

func (s Storage) CreateUser(u User, ch chan int) {
	createUserQuery := "INSERT INTO users(id, use, avatar, description) VALUES ($1, $2, $3, $4)"

	_, err := s.DB.Exec(s.Context, createUserQuery, u.ID, u.Use, u.Avatar, u.Description)
	if err != nil {
		s.Log.Error("error with creating user:", err)
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusCreated
}

func (s Storage) GetUserById(ID int, ch chan *User) {
	getUserQuery := "SELECT * FROM users WHERE id=$1"
	row := s.DB.QueryRow(s.Context, getUserQuery, ID)

	var user User
	err := row.Scan(&user.ID, &user.Use, &user.Avatar, &user.Description, &user.Rating)
	if err != nil {
		if err == pgx.ErrNoRows {
			s.Log.Info("user not found with id:", ID)
			ch <- nil
			return
		}
		s.Log.Error("error with getting user: ", err)
		ch <- nil
		return
	}

	ch <- &user
}

func (s Storage) UpdateUser(u User, ch chan int) {
	updateUserQuery := "UPDATE users SET use=$2, avatar=$3, description=$4 WHERE id=$1"
	_, err := s.DB.Exec(s.Context, updateUserQuery, u.ID, u.Use, u.Avatar, u.Description)
	if err != nil {
		s.Log.Error("error with updating user: ", err)
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusOK
}

func (s Storage) DeleteUser(ID int, ch chan int) {
	deleteUserQuery := "DELETE FROM users WHERE id=$1"

	_, err := s.DB.Exec(s.Context, deleteUserQuery, ID)
	if err != nil {
		s.Log.Error("error with deleting user: ", err)
		ch <- http.StatusBadRequest
	}
	ch <- http.StatusOK
}
