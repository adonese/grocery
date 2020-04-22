package main

import (
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
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
	Delivery    time.Time `db:"delivery"`
	IsCompleted bool      `db:"is_completed"`
	ProductID   int       `db:"product_id"`
	Quantity    int       `db:"quantity"`
	Token       string    `db:"token"`
}

func (c *Cart) generateToken() {
	t := uuid.New().String()
	c.Token = t
}

func (c *Cart) save() error {
	db, err := getDB("data.db")
	if err != nil {
		return err
	}

	db.Exec(stmt)
	tx := db.MustBegin()
	tx.NamedExec("insert into carts(user_id, created_at, product_id, quantity, token) values(:user_id, :created_at, :product_id, :quantity, :token)", c)
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}

type Product struct {
	Name string
	ID   int
}

type Price struct {
	ID        int
	UnitPrice float32
}
