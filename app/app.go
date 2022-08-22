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
	if err != nil {
		log.Fatalf("Couldn't intialize PostgreSQL connection: %s", err.Error())
	}
	// initialize Ctx as an empty context for demo, shouldn't be empty
	Ctx = context.Background()
	redisRepo, err = repositories.NewRedisClient(Ctx)
	if err != nil {
		log.Fatalf("Couldn't intialize Redis connection: %s", err.Error())
	}
	// connect initialized repos with the handler
	handler = &controller.BeerHandler{PostgreSQL: psqlRepo, Redis: redisRepo}
}

func Run() {
	defer psqlRepo.Close()
	defer redisRepo.Close()
	http.HandleFunc("/beer", handler.GetBeerByID)
	http.HandleFunc("/beers", handler.GetAllBeers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
