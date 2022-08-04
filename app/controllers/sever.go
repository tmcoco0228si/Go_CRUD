package controllers

import (
	"Go_CRUD/config"
	"fmt"
	"html/template"
	"net/http"
)

//テンプレートを渡して表示するハンドラ関数を共通化する関数
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		//sプリントでfileをsに入れる。
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	//エラーの際にパニック状態になるのが[Must]
	templates := template.Must(template.ParseFiles(files...))
	//(レスポンス, 実行するテンプレート, data)
	templates.ExecuteTemplate(w, "layout", data)
}

//サーバ起動
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//(ポート番号, デフォルトマルチプレクサ)
	//登録されていないURLにアクセスすると404を返す設定がnil
	//URL登録(url,template)
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
