package models

import (
	"Go_CRUD/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	//table作成 %s=文字列(string)
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	  uuid STRING NOT NULL UNIQUE,
	  name STRING,
	  email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)

	//table作成
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY AUTOINCREMENT, content TEXT, user_id INTEGER, created_at DATETIME)`, tableNameTodo)
	Db.Exec(cmdT)
}

//userのUUID生成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

//パスワードをハッシュ値にする処理
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
