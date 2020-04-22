package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp, _ = template.ParseFiles("static/template.html")

func form(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title": "Hello Solus",
	}
	temp.Execute(w, data)
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
	cookie := u.generateCookie()
	log.Printf("The cookie is: %#v", cookie)
	http.SetCookie(w, cookie)
	registerPage.Execute(w, errors)
}

func main() {
	http.HandleFunc("/", form)
	http.HandleFunc("/register", register)
	http.ListenAndServe(":8080", nil)
}
