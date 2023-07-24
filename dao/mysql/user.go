package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"tieba/models"
)

const secret = "261077486@qq.com"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)

// CheckUserExit 将数据库操作封装到Dao层，让logic层根据需求来调用
func CheckUserExit(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int64
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	//加密密码
	user.Password = encryptPassword(user.Password)
	//执行sql入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登录时输入的密码
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username) //找到数据库中对当前用户名对应的密码，赋值给user.Password
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}
