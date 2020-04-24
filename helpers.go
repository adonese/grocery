package main

import (
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

	// make sure length of items == length = quarantines
	for k := range items {
		c[k].ProductID = toInt(items[k])
		c[k].Quantity = toInt(quantities[k]) // FIXME this could be a float

	}
	return c
}
