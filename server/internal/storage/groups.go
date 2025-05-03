package storage

import (
	"net/http"
)

type CreateGroupRequest struct {
	ID string `json:"id"`
}

type GroupRequest struct {
	ID     string `json:"id"`
	UserID int    `json:"user_id"`
}

type Group struct {
	ID           string `json:"id"` // id ресторана
	Participants []int  `json:"participants"`
}

func (s Storage) CreateGroup(ID string, ch chan int) {
	err := s.RDB.LPush(s.Context, ID, "").Err()
	if err != nil {
		s.Log.Error("error with creating group")
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusCreated
}

func (s Storage) JoinGroup(ID string, userID int, ch chan int) {
	err := s.RDB.LPush(s.Context, string(ID), userID).Err()
	if err != nil {
		s.Log.Error("error with joinig group", err)
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusOK
}

func (s Storage) LeaveGroup(ID string, userID int, ch chan int) {
	err := s.RDB.LRem(s.Context, string(ID), 1, userID).Err()
	if err != nil {
		s.Log.Error("error with leaving group: ", err)
		ch <- http.StatusInternalServerError
		return
	}

	ch <- http.StatusOK
}

func (s Storage) DeleteGroup(ID string, ch chan int) {
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
		ch <- nil
		return
	}
	ch <- slice
}

func (s Storage) GetAllUserFromGroup(ID string, ch chan []string) {
	slice, err := s.RDB.LRange(s.Context, ID, 0, -1).Result()
	if err != nil {
		s.Log.Error("error with getting users from group: ", err)
		ch <- nil
		return
	}

	ch <- slice
}
