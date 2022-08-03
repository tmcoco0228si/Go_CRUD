package controllers

import (
	"Go_CRUD/config"
	"fmt"
	"html/template"
	"net/http"
)

//
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		//sプリントでfileをsに入れる。
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//(ポート番号, デフォルトマルチプレクサ)
	//登録されていないURLにアクセスすると404を返す設定がnil
	//URL登録
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
