package controllers

import (
	"Go_CRUD/app/models"
	"log"
	"net/http"
)

//ハンドラ作成
func signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {

		//POSTリクエスト受け取れる(ParseForm)
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := &models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}
		//リダイレクト先のURL指定
		http.Redirect(w, r, "/", 302)
	}
}

//ログインハンドラ
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}

	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		cookie := http.Cookie{
			Name: "_cookie",

			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
