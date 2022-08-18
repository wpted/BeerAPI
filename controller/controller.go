package controller

import (
	"BeerAPI/repositories"
	"BeerAPI/util"
	"fmt"
	"net/http"
	"strconv"
)

type BeerHandler struct {
	*repositories.PostgreSQL
	*repositories.Redis
}

func (b *BeerHandler) CreateBeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//id := r.FormValue("id")

	//b.InsertBeer(id)

}

func (b *BeerHandler) GetBeerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	idString := r.FormValue("id")

	// check if data exist in Redis
	fmt.Println(idString)
	val, err := b.Redis.GetFromRedis(r.Context(), idString)
	fmt.Println(val)
	if err == nil {
		// if data exist return the cached value to the response
		util.RenderJson(w, val, http.StatusOK)
		return
	}
	// if doesn't exist convert the id string to id int
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, http.StatusText(400)+", please check your input", http.StatusBadRequest)
	}
	// get value corresponding to the id from the PostgreSQL server
	result := b.GetBeer(id)

	// after getting from psql, add to the cache pool
	b.Redis.AddToRedis(r.Context(), idString, result, 0)
	// send the data back to the response in json form
	util.RenderJson(w, result, http.StatusOK)
}

func (b *BeerHandler) GetAllBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	requestedBeers := b.GetBeers()

	fmt.Fprint(w, requestedBeers)

}
