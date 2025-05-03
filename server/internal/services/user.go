package services

import (
	"log/slog"
	"strconv"

	"github.com/Flikest/food/internal/storage"
	"github.com/gofiber/fiber/v2"
)

var (
	statusCodeChan = make(chan int)
	userChan       = make(chan *storage.User)
)

func (s Service) CreateUser(ctx *fiber.Ctx) error {
	var user storage.User
	ctx.BodyParser(&user)

	go s.Storage.CreateUser(user, statusCodeChan)

	result := <-statusCodeChan
	if result > 299 {
		return ctx.JSON("failed to create user 🤔")
	}
	return ctx.JSON("Welcome to our bot 👋")
}

func (s Service) GetUserById(ctx *fiber.Ctx) error {
	StringID := ctx.Params("id")
	ID, err := strconv.Atoi(StringID)
	if err != nil {
		slog.Error("error while converting data:", err)
	}

	go s.Storage.GetUserById(ID, userChan)

	result := <-userChan

	if result != nil {
		return ctx.JSON("user with such id not found 👀")
	}
	return ctx.JSON(&result)
}

func (s Service) UpdateUser(ctx *fiber.Ctx) error {
	var user storage.User
	ctx.BodyParser(&user)

	go s.Storage.UpdateUser(user, statusCodeChan)

	result := <-statusCodeChan
	if result > 299 {
		return ctx.JSON("failed to updating user 🆙")
	}
	return ctx.JSON(user)
}

func (s Service) DeleteUser(ctx *fiber.Ctx) error {
	stringID := ctx.Params("id")
	ID, err := strconv.Atoi(stringID)
	if err != nil {
		slog.Error("error while converting data:", err)
	}

	go s.Storage.DeleteUser(ID, statusCodeChan)
	result := <-statusCodeChan
	if result > 299 {
		return ctx.JSON("failed to deleting user 🗑️")
	}
	return ctx.JSON("user with id: %s deleted 🗑️", stringID)
}
