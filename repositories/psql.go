package repositories

import (
	"BeerAPI/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type PostgreSQL struct {
	DB *sql.DB
}

// -------------------- Connection Logic ---------------------

func NewPostgreSQL() (*PostgreSQL, error) {
	db, err := sql.Open("postgres", "postgres://beer_fellow:lovebeer@localhost/beer_server?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("You've successfully connected to PostgreSQL")
	return &PostgreSQL{DB: db}, nil

}

func (p *PostgreSQL) Close() {
	p.DB.Close()
}

// -------------------- Access Database ---------------------

func (p *PostgreSQL) InsertBeer(b model.Beer) {
	insertQuery := " INSERT INTO beers (name, price, company) VALUES($1, $2, $3)"
	_, err := p.DB.Exec(insertQuery, b.Name, b.Price, b.Company)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(b)
}

func (p *PostgreSQL) GetBeers() []model.Beer {
	beers := make([]model.Beer, 0)
	getBeersQuery := "SELECT * FROM beers"
	rows, err := p.DB.Query(getBeersQuery)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		beer := model.Beer{}
		rows.Scan(&beer.ID, &beer.Name, &beer.Price, &beer.Company)
		beers = append(beers, beer)
	}
	return beers
}

func (p *PostgreSQL) GetBeer(id int) model.Beer {
	beer := model.Beer{}
	getBeerQuery := "SELECT * FROM beers WHERE ID=$1"
	row := p.DB.QueryRow(getBeerQuery, id)
	err := row.Scan(&beer.ID, &beer.Name, &beer.Price, &beer.Company)
	switch {
	case err == sql.ErrNoRows:
		return model.Beer{}
	case err != nil:
		return model.Beer{}
	default:
		return beer
	}
}
