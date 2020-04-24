package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func getDB(filename string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getCookie(r *http.Request, name string) (*http.Cookie, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func toInt(v string) int {
	d, _ := strconv.Atoi(v)
	return d
}

func getItems(r *http.Request) []CartItems {
	var c []CartItems

	r.ParseForm()
	items := r.Form["items"]
	quantities := r.Form["quantities"]
	log.Printf("Items are: %v", items)

	// make sure length of items == length = quarantines
	var cart CartItems
	for k := range items {
		cart.ProductID = toInt(items[k])
		cart.Quantity = toInt(quantities[k])
		c = append(c, cart)

	}
	return c
}
