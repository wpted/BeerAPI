package app

import (
	"BeerAPI/repositories"
	"context"
	"log"
)

var (
	Ctx       context.Context
	redisRepo *repositories.Redis
	psqlRepo  *repositories.PostgreSQL
)

func init() {
	var err error

	psqlRepo, err = repositories.NewPostgreSQL()

	Ctx = context.Background()
	redisRepo, err = repositories.NewRedisClient(Ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func Run() {

}
