package app

import (
	"BeerAPI/controller"
	"BeerAPI/repositories"
	"context"
	"log"
	"net/http"
)

var (
	Ctx       context.Context
	redisRepo *repositories.Redis
	psqlRepo  *repositories.PostgreSQL
	handler   *controller.BeerHandler
)

func init() {
	var err error

	psqlRepo, err = repositories.NewPostgreSQL()

	Ctx = context.Background()
	redisRepo, err = repositories.NewRedisClient(Ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	handler = &controller.BeerHandler{PostgreSQL: psqlRepo}
}

func Run() {
	defer psqlRepo.Close()
	defer redisRepo.Close()
	http.HandleFunc("/beer", handler.GetBeerByID)
	http.HandleFunc("/beers", handler.GetAllBeers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
