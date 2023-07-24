package logic

import (
	"tieba/dao/mysql"
	"tieba/models"
	"tieba/pkg/jwt"
	"tieba/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户是否存在
	if err = mysql.CheckUserExit(p.Username); err != nil {
		return err
	}
	//2.生产UID
	userID := snowflake.GenID()
	//构造一个User实例接收p中的注册信息
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//判断密码是否正确
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	//生成jwt
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
