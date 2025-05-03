package storage

import (
	"fmt"
	"log"
	"net/http"
)

type UpdateRatingRequest struct {
	Operation string `json:operation`
	User_id   int    `json:"user_id"`
}

func (s Storage) UpdateRating(urq UpdateRatingRequest, ch chan int) {

	updateRatingQuery := "UPDATE users SET rating = rating %s 20 WHERE id=$1"

	switch urq.Operation {
	case "+":
		query := fmt.Sprintf(updateRatingQuery, "+")
		log.Println(query)
		_, err := s.DB.Exec(s.Context, query, urq.User_id)
		if err != nil {
			s.Log.Error("error when increasing rating: ", err)
			ch <- http.StatusInternalServerError
			return
		}
		ch <- http.StatusOK
	case "-":
		query := fmt.Sprintf(updateRatingQuery, "-")
		_, err := s.DB.Exec(s.Context, query, urq.User_id)
		if err != nil {
			s.Log.Error("error when decreasing rating: ", err)
			ch <- http.StatusInternalServerError
			return
		}
		ch <- http.StatusOK
	default:
		ch <- http.StatusBadRequest
	}

}
