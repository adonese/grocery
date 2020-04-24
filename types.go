package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// User struct in the system
type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Mobile   string `db:"mobile"`
	Telegram string `db:"telegram"`
}

func (u *User) saveUser() error {
	db, _ := getDB("data.db")
	defer db.Close()

	db.Exec(stmt)

	tx := db.MustBegin()
	tx.NamedExec("insert into users(username, mobile, telegram) values(:username, :mobile, :telegram)", u)
	if err := tx.Commit(); err != nil {
		log.Printf("Error in DB: %v", err)
		return err
	}
	return nil
}

func (u *User) getUser(username string) error {
	db, _ := getDB("data.db")
	defer db.Close()

	db.Exec(stmt)

	tx := db.MustBegin()
	tx.Get(u, "select * from users where username = $1", username)
	if err := tx.Commit(); err != nil {
		log.Printf("Error in DB: %v", err)
		return err
	}
	return nil
}

func (u User) generateCookie() *http.Cookie {
	c := &http.Cookie{
		Name:       "grocery",
		Value:      u.Username,
		Path:       "/",
		Domain:     "",
		Expires:    time.Now().Add(1000 * time.Hour),
		RawExpires: "",
		MaxAge:     0,
		// Secure:     true,
		// HttpOnly:   true,
		SameSite: 0,
		Raw:      "",
		Unparsed: nil,
	}
	return c
}

type Cart struct {
	ID          int            `db:"id"`
	UserID      int            `db:"user_id"`
	CreatedAt   sql.NullString `db:"created_at"`
	Delivery    sql.NullString `db:"delivery"`
	IsCompleted sql.NullBool   `db:"is_completed"`
	ProductID   int            `db:"product_id"`
	Quantity    int            `db:"quantity"`
	Token       string         `db:"token"`
}

func newCart() *Cart {
	return &Cart{}
}

func (c *Cart) generateToken() {
	t := uuid.New().String()
	c.Token = t
}

func (c *Cart) save() (int, error) {
	var count int
	db, err := getDB("data.db")
	if err != nil {
		return count, err
	}
	defer db.Close()

	db.Exec(stmt)
	tx := db.MustBegin()
	tx.NamedExec("insert into carts(user_id, created_at, product_id, quantity, token) values(:user_id, :created_at, :product_id, :quantity, :token)", c)
	if err := tx.Commit(); err != nil {
		log.Printf("Error in cart.save: TX: %v", err)
		return count, err
	}
	db.Get(&count, "select id from carts order by id desc limit 1")

	return count, nil

}

func (c *Cart) get(id int) ([]Cart, error) {
	var carts []Cart

	db, err := getDB("data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.Exec(stmt)
	if err := db.Select(&carts, "select * from carts where user_id = $1", id); err != nil {
		return nil, err
	}
	return carts, nil
}

type CartItems struct {
	/*
			id integer primary key,
		    product_it integer,
		    cart_id integer,
			user_id integer
	*/
	ID        int `db:"id"`
	UserID    int `db:"user_id"`
	CartID    int `db:"cart_id"`
	ProductID int `db:"product_id"`
	Quantity  int `db:"quantity"`
}

func (c *CartItems) populate() error {
	db, err := getDB("data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	tmt := `create table cartitems (
    id integer primary key,
    product_id integer,
    cart_id integer,
	user_id integer,
	quantity integer
	)`

	db.Exec(tmt)
	log.Printf("Value of caritems is: %v", c)
	// tx := db.MustBegin()
	if _, err := db.NamedExec("insert into cartitems(user_id, cart_id, product_id, quantity) values(:user_id, :cart_id, :product_id, :quantity)", c); err != nil {
		log.Printf("Error in CartItems.populate: %v", err)
		return err
	}

	return nil
}

func (c *CartItems) all() ([]CartItems, error) {
	var items []CartItems

	db, err := getDB("data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tmtt := `create table cartitems (
    id integer primary key,
    product_it integer,
    cart_id integer,
	user_id integer
)`
	db.Exec(tmtt)

	tx := db.MustBegin()
	tx.Select(&items, "select * from cartitems")
	if err := tx.Commit(); err != nil {
		log.Printf("Error in CartItems.all: %v", err)
		return nil, err
	}
	return items, nil
}

type Product struct {
	Name string
	ID   int
}

type Price struct {
	ID        int
	UnitPrice float32
	ProductID int
}
