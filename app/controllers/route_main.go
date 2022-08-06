package controllers

import (
	"Go_CRUD/app/models"
	"log"
	"net/http"
)

//ハンドラ
func top(w http.ResponseWriter, r *http.Request) {
	// //template.ParseFiles = file解析
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// //(レスポンスW,データをhtmlに渡す)
	// t.Execute(w, "hello")
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

//ログインしたユーザのみが閲覧できるハンドラ
func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		log.Println("------------------------------")
		log.Println(err)
		log.Println("------------------------------")

		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}

}

//todoを作成できるハンドラ
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_New")
	}

}

//TODO作成
func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()

		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}

//TODO更新のハンドラ
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		_, err := sess.GetUserBySession()

		if err != nil {
			log.Println(err)
		}
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}
}

//ユーザのTODOを更新する処理
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}

}
