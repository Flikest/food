package services

import (
	"github.com/Flikest/food/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func (s Service) CreateGroup(ctx *fiber.Ctx) error {
	var request storage.CreateGroupRequest
	ctx.BodyParser(&request)

	var statusCodeChan = make(chan int)
	go s.Storage.CreateGroup(string(request.ID), statusCodeChan)

	result := <-statusCodeChan

	if result > 299 {
		return ctx.JSON("failed to create a group for shared meals 🍽️")
	}
	return ctx.JSON("a group for common meals has been created 🍽️")
}

func (s Service) JoinGroup(ctx *fiber.Ctx) error {
	var request storage.GroupRequest
	ctx.BodyParser(&request)

	var statusCodeChan = make(chan int)
	go s.Storage.JoinGroup(request.ID, request.UserID, statusCodeChan)

	result := <-statusCodeChan

	if result > 299 {
		return ctx.JSON("we couldn't connect you to this session 😞")
	}
	return ctx.JSON("Bon appetit 🥓")
}

func (s Service) LeaveGroup(ctx *fiber.Ctx) error {
	var request storage.GroupRequest
	ctx.BodyParser(&request)

	var statusCodeChan = make(chan int)
	go s.Storage.LeaveGroup(request.ID, request.UserID, statusCodeChan)

	result := <-statusCodeChan

	if result > 299 {
		return ctx.JSON("failed to leave group 😞")
	}
	return ctx.JSON("you left the group 🏃🚪")
}

func (s Service) DeleteGroup(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	var statusCodeChan = make(chan int)
	go s.Storage.DeleteGroup(ID, statusCodeChan)

	result := <-statusCodeChan

	if result > 299 {
		return ctx.JSON("the group has not been disbanded")
	}
	return ctx.JSON("the group has been disbanded 🗑️")
}

func (s Service) GetAllGroup(ctx *fiber.Ctx) error {
	ch := make(chan []string)

	go s.Storage.GetAllGroup(ch)

	result := <-ch

	return ctx.JSON(result[2:])
}

func (s Service) GetAllUserFromGroup(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	ch := make(chan []string)
	go s.Storage.GetAllUserFromGroup(ID, ch)

	result := <-ch
	if len(result)-1 == 0 {
		return ctx.JSON("There is no one in this group")
	}
	return ctx.JSON(result[0 : len(result)-1])

}
