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
	Todos      []Todo
}

type Session struct {
	ID         int
	UUID       string
	Email      string
	UserID     int
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

//セレクト(ユーザが入力したEmialで探す)
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.Created_at)

	return user, err
}

//セッション生成,取得
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (uuid,email, user_id,created_at)
	values(?,?,?,?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}
	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.Created_at)

	return session, err
}

//ユーザがセッションを持っているかチェックしている。
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.Created_at)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

//セッション削除する処理 DeleteSessionByUUID
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id,uuid, name, email, created_at from users where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Created_at)
	return user, err
}
