package services

import (
	"github.com/Flikest/food/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func (s Service) UpdateRating(ctx *fiber.Ctx) error {
	var request storage.UpdateRatingRequest
	ctx.BodyParser(request)

	var statusCodeChan = make(chan int)

	go s.Storage.UpdateRating(request, statusCodeChan)

	result := <-statusCodeChan

	if result > 299 {
		return ctx.JSON("failed to update rating for user under id %s", string(request.User_id))
	}

	return ctx.JSON("you updated the user rating under id 🆙")
}
