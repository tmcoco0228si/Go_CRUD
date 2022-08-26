package controllers

import (
	"Go_CRUD/app/models"
	"Go_CRUD/config"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)


//テンプレートを渡して表示するハンドラ関数を共通化する関数
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	//<< template.ParseFiles >> 外部ファイルを取り込む
	templates := template.Must(template.ParseFiles(files...))
	//<< ExecuteTemplate >>テンプレートへの値の埋め込み
	templates.ExecuteTemplate(writer, "layout", data)
}

//クッキー取得
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

//文字列の正規表現
var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		//validPathとURLがマッチした箇所を変数に代入
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		//URL末尾にIDがついてきていると仮定して、ついている場合、数値型に変換して代入している。
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

//サーバ起動
func StartMainServer() error {
	//handerを返す
	// << Dir >> string型のディレクトリ名 "html" を http.Dir型 にキャストする
	// http.Dir型 は FileSystemインターフェースを満たすので
	// これを http.FileServer() に渡して 静的ファイルを返すハンドラを得る
	//<< StripPrefix >> filesを元に、先頭の "/static/" を除外するハンドラを作成する
	// DefaultServeMux に URLパス "/file/" と split_prefix_hdlr のペアを追加する

	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	//URL登録(url,template)
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	//(ポート番号, デフォルトマルチプレクサ)
	//登録されていないURLにアクセスすると404を返す設定がnil
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
