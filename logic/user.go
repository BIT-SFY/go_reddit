package logic

import (
	"reddit/dao/mysql"
	"reddit/models"
	"reddit/pkg/snowflake"
)

//存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户存不存在
	if isExist := mysql.CheckUserExist(p.Username); isExist != false {
		return mysql.ErrorUserExist
	}
	//2.生成UID
	userID := snowflake.GenID()
	//构造一个User实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//3.保存到数据库
	return mysql.InsertUser(&user)
}

func Login(p *models.ParamLogin) error {
	//1.根据用户名和密码检验该用户是否存在或者密码是否输入正确
	user := models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(&user)
}