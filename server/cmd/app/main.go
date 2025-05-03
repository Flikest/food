package main

import (
	"context"
	"flag"
	"os"

	"github.com/Flikest/food/internal/database"
	"github.com/Flikest/food/internal/handler"
	"github.com/Flikest/food/internal/services"
	"github.com/Flikest/food/internal/storage"
	"github.com/Flikest/food/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "", "Port for the application")

	flag.Parse()

	godotenv.Load()

	log := logger.InitLogger(os.Getenv("LVL_DEPLOYMENT"))

	db, err := database.NewDataBase(os.Getenv("DB_PATH"))
	if err != nil {
		log.Error("error with connecting database", err)
	}
	log.Info("the database is connected 😃")

	rdb := database.NewRedisClient()
	log.Info("the redis is connected 🤩")

	storage := storage.InitStorage(storage.Storage{
		DB:      db,
		Context: context.Background(),
		RDB:     rdb,
	})
	services := services.InitService(storage)
	handler := handler.InitHandler(services)
	router := handler.NewRouter()

	router.Listen(":" + port)
}
