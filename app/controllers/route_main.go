package controllers

import (
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

	generateHTML(w, "hello", "layout", "top")
}
