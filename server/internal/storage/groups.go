package storage

import (
	"net/http"
)

type CreateGroupRequest struct {
	ID int `json:"id"`
}

type GroupRequest struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type Group struct {
	ID           int   `json:"id"` // id ресторана
	Participants []int `json:"participants"`
}

func (s Storage) CreateGroup(ID int, ch chan int) {
	err := s.RDB.LPush(s.Context, string(ID)).Err()
	if err != nil {
		s.Log.Error("error with creating group")
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusCreated
}

func (s Storage) JoinGroup(ID int, userID int, ch chan int) {
	err := s.RDB.LPush(s.Context, string(ID), userID).Err()
	if err != nil {
		s.Log.Error("error with joinig group", err)
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusOK
}

func (s Storage) LeaveGroup(ID int, userID int, ch chan int) {
	err := s.RDB.LRem(s.Context, string(ID), 1, userID).Err()
	if err != nil {
		s.Log.Error("error with leaving group: ", err)
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusOK
}

func (s Storage) DeleteGroup(ID int, ch chan int) {
	err := s.RDB.Del(s.Context, string(ID)).Err()
	if err != nil {
		s.Log.Error("error with deleting group: ", err)
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusOK
}

func (s Storage) GetAllGroup(ch chan []string) {
	slice, err := s.RDB.Keys(s.Context, "*").Result()
	if err != nil {
		s.Log.Error("error with getting groups: ", err)
	}
	ch <- slice
}

func (s Storage) GetAllUserFromGroup(ID int, ch chan []string) {
	slice, err := s.RDB.LRange(s.Context, string(ID), 0, -1).Result()
	if err != nil {
		s.Log.Error("error with getting users from group: ", err)
	}

	ch <- slice
}
