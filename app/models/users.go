package models

import (
	"log"
	"time"
)

type User struct {
	ID         int
	UUID       string
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}

func (u *User) CreateUser() (err error) {
	//SQL生成
	cmd := `INSERT INTO users (
		uuid, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)`

	//SQL実行
	_, err = Db.Exec(
		cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ユーザ取得
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`

	//Scanで、検索結果を変数に代入している。
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Created_at,
	)
	return user, err
}

//ユーザ更新
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

//ユーザ削除
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
