package handler

import (
	"github.com/Flikest/food/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *services.Service
}

func InitHandler(s *services.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h Handler) NewRouter() *fiber.App {
	router := fiber.New()

	v1 := router.Group("/v1")
	{
		userRouter := v1.Group("/user")
		{
			userRouter.Post("/", h.Service.CreateUser)
			userRouter.Get("/:id", h.Service.GetUserById)
			userRouter.Patch("/", h.Service.UpdateUser)
			userRouter.Delete("/:id", h.Service.DeleteUser)
		}

		groupsRouter := v1.Group("/room")
		{
			groupsRouter.Get("/:id", h.Service.GetAllUserFromGroup)
			groupsRouter.Get("/", h.Service.GetAllGroup)
			groupsRouter.Post("/", h.Service.CreateGroup)
			groupsRouter.Post("/join", h.Service.JoinGroup)
			groupsRouter.Delete("/leave/", h.Service.LeaveGroup)
			groupsRouter.Delete("/:id", h.Service.DeleteGroup)

		}

		raitingRouter := v1.Group("/raiting")
		{
			raitingRouter.Patch("/", h.Service.UpdateRating)
		}
	}
	return router
}
