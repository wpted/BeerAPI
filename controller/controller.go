package controller

import (
	"BeerAPI/model"
	"BeerAPI/repositories"
	"BeerAPI/util"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

const CacheExpirationTime = 180 * time.Second

// BeerHandler is a struct that contains a pool of the PSQL connection and the Redis connection
type BeerHandler struct {
	*repositories.PostgreSQL
	*repositories.Redis
}

//func (b *BeerHandler) CreateBeer(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("content-type", "application/json")
//	//id := r.FormValue("id")
//
//	//b.InsertBeer(id)
//
//}

// GetBeerByID is a handler that gets beer info from the database with the given id and writes is as json to the http response
func (b *BeerHandler) GetBeerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	idString := r.FormValue("id")
	// check if data exist in Redis
	val, err := b.Redis.GetFromRedis(context.Background(), idString)
	if err == nil {
		cachedBeer := model.Beer{}
		// decode the redis string value back into a struct
		err = json.Unmarshal([]byte(val), &cachedBeer)
		if err != nil {
			// if anything wrong it must be the retrieved value doesn't bind with the beer struct
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// if data exist return the cached value to the response
		util.RenderJson(w, cachedBeer, http.StatusOK)
		return
	}
	// if the data doesn't exist convert the id string to id int
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, http.StatusText(400)+", please check your input", http.StatusBadRequest)
	}
	// get value corresponding to the id from the PostgreSQL server
	result := b.GetBeer(id)
	resultJson, err := result.MarshalBinary()
	// after getting from psql, add to the cache pool
	err = b.Redis.AddToRedis(context.Background(), idString, resultJson, CacheExpirationTime)
	if err != nil {
		log.Fatal(err.Error())
	}
	//// send the data back to the response in json form
	util.RenderJson(w, result, http.StatusOK)
}

// GetAllBeers is a handler that gets all existing beer info within the database and writes it as a json to the http response
func (b *BeerHandler) GetAllBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	result := b.GetBeers()
	util.RenderJson(w, result, http.StatusOK)
}
