package controllers

import (
	"Go_CRUD/app/models"
	"log"
	"net/http"
)

//ハンドラ作成
func signup(w http.ResponseWriter, r *http.Request) {
	
	if r.Method == "GET" {
		_, err := session(w, r)

		//セッションが存在しない場合
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			//セッションが存在する場合
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {

		//POSTリクエスト受け取れる(ParseForm)
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := &models.User{
			//<< PostFormValue >> keyを指定してformの値取得
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

//ログイン画面を返すハンドラ
func login(w http.ResponseWriter, r *http.Request) {
	//セッション情報があるかチェック
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

//ログイン セッション確認ハンドラ
func authenticate(w http.ResponseWriter, r *http.Request) {
	//パラメータを全て取得(ParseForm)
	err := r.ParseForm()

	//userが入力したメールアドレスの認証
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
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		//レスポンスヘッダーにセッション設定
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//ログアウト処理（付与したセッション削除）
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}
