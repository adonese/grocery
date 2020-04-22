package main

import (
	"log"
	"net/http"
	"time"

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

type Card struct {
	ID          int
	UserID      int
	CreatedAT   time.Time
	Delivery    time.Time
	IsCompleted bool
	ProductID   int
	Quantity    int
	Token       string
}

type Product struct {
	Name string
	ID   int
}

type Price struct {
	ID        int
	UnitPrice float32
}
