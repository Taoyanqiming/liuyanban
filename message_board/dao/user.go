package dao

import (
	"message_board/model"
)

func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("Insert into user(name,password) values (?,?)", u.UserName, u.PassWord)
	return
}
func SearchUserByUserName(name string) (u model.User, err error) {
	row := DB.QueryRow("select id,name,password from user where name = ?", name)
	if err = row.Err(); row.Err() != nil {
		err = row.Scan(&u.Id, &u.UserName, &u.PassWord)
	}
	return
}
func UpdataPassword(username, newPassword string) error {
	_, err := DB.Exec("UPDATE user SET password=?WHERE username=?", newPassword, username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}
	row := DB.QueryRow("SELECT id,password FRON user WHERE username?", username)
	if row != nil {
		return user, row.Err()
	}
	err := row.Scan(&user.Id, &user.PassWord)
	if err != nil {
		return user, err
	}
	return user, nil
}
