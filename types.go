package main

import (
	"net/http"
	"time"
)

// User struct in the system
type User struct {
	Username string
	Mobile   string
	Telegram string
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
