package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"message_board/dao"
	"message_board/model"
)

func SearchUserByUserName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUserName(name)
	return
}
func CreateUser(u model.User) error {
	//对用户插入操作
	err := dao.InsertUser(u)
	return err
}

func ChangePassword(username, newPassword string) error {
	err := dao.UpdataPassword(username, newPassword)
	return err

}

func IsRepeatUserName(username string) (bool, error) {
	_, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err

	}
	return true, nil

}

func IsPasswordCorrect(username, password string) (bool, error) {
	user, err := dao.SelectUserByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	if user.PassWord != password {
		return false, err
	}
	return true, nil

}
