package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var temp, _ = template.ParseFiles("static/template.html")

var ordersTemp, _ = template.ParseFiles("static/orders.html")

var stmt = `
create table users (
	id integer primary key,
	username text unique,
	mobile text unique,
	telegram text
);
	
create table carts (
	id integer primary key,
	user_id integer,
	created_at time,
	delivery time,
	is_completed bool,
	product_id integer,
	quantity integer,
	token text
);

create table products (
    name text,
    product_id integer
);

create table prices (
    id integer primary key,
    unit_price real
);

`

func form(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title": "Hello Solus",
	}

	// get user
	var u User
	c, err := getCookie(r, "grocery")
	if err != nil {
		data["error"] = err.Error()
		u.Username = "anon"
		temp.Execute(w, data)
		return
	}
	// log.Printf("Cookie is: %v", c.Value)
	if err := u.getUser(c.Value); err != nil {
		log.Printf("Error in getting user: %v", err)
		data["error"] = err.Error()
		u.Username = "anon"
	}

	log.Printf("Loaded user profile is: %#v", u)
	data["username"] = u.Username
	if r.Method == "GET" {
		temp.Execute(w, data)
		return
	}

	// for submitting part
	item := r.PostFormValue("item")
	quantity := r.PostFormValue("quantity")
	// notes := r.PostFormValue("notes")
	// email := r.PostFormValue("email")

	var cart Cart
	cart.UserID = u.ID

	cart.CreatedAt.Scan(time.Now())
	cart.generateToken()
	cart.Quantity = toInt(quantity)
	cart.ProductID = toInt(item)

	if err := cart.save(); err != nil {
		log.Printf("Error in cart.save: %v", err)
		data["error"] = err.Error()
		temp.Execute(w, data)
		return
	}

}

func register(w http.ResponseWriter, r *http.Request) {
	var registerPage, err = template.ParseFiles("static/register.html")
	if err != nil {
		log.Fatalf("Template error: Error in parsing register page: %v", err)
	}

	errors := make(map[string]string)
	errors["title"] = "Register - Grocery"
	if r.Method == "GET" {
		registerPage.Execute(w, errors)
		return
	}

	// it is a post method
	var u User
	defer r.Body.Close()

	// req, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	errors["register"] = err.Error()
	// }

	u.Username = r.PostFormValue("username")
	u.Mobile = r.PostFormValue("mobile")
	u.Telegram = r.PostFormValue("telegram")
	if u.Username == "" || u.Mobile == "" {
		errors["error"] = "empty username or mobile"
		registerPage.Execute(w, nil)
		return
	}

	// save the result to database
	log.Printf("User profile in registration: %v", u)
	cookie := u.generateCookie()
	if err := u.saveUser(); err != nil {
		errors["error"] = err.Error()
	}
	log.Printf("The cookie is: %#v", cookie)
	http.SetCookie(w, cookie)
	registerPage.Execute(w, errors)
}

func orders(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	// get user
	var u User
	c, err := getCookie(r, "grocery")
	if err != nil {
		log.Printf("Error in getting cookie: %v", err)
		data["error"] = err.Error()
		u.Username = "anon"
		ordersTemp.Execute(w, data)
		return
	}
	// log.Printf("Cookie is: %v", c.Value)
	if err := u.getUser(c.Value); err != nil {
		log.Printf("Error in getting user: %v", err)
		data["error"] = err.Error()
		u.Username = "anon"
	}
	log.Printf("Loaded user is: %v", u)
	var cart Cart

	if carts, err := cart.get(u.ID); err != nil {
		log.Printf("Error in getting cart: %v", err)
		data["error"] = err.Error()
	} else {
		log.Printf("Loaded cart is: %#v", carts)
		ordersTemp.Execute(w, carts)
	}

}

func main() {
	http.HandleFunc("/", form)
	http.HandleFunc("/register", register)
	http.HandleFunc("/orders", orders)
	http.ListenAndServe(":8080", nil)
}
