package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"reddit/models"
)

//把每一步数据库操作封装成函数 待Logic层根据业务需求调用

const secret = "3220231821-shenfuyuan"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CheckUserExist 根据用户名检查该用户是否存在
func CheckUserExist(username string) bool {
	user := models.User{}
	db.Where("username = ?", username).Find(&user)
	return user.ID != 0 //如果为0则不存在该用户,不为0则存在该用户
}

// InsertUser 向数据库中插入新用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句
	if err = db.Create(user).Error; err != nil {
		return err
	}
	return
}

// Login 用户登录
func Login(user *models.User) (err error) {
	//检查用户名是否存在
	if isExist := CheckUserExist(user.Username); !isExist {
		return ErrorUserNotExist
	}
	//校验密码
	user.Password = encryptPassword(user.Password)
	_user := models.User{}
	db.Where("username = ? and password = ?", user.Username, user.Password).Find(&_user)
	if _user.ID == 0 {
		return ErrorInvalidPassword
	}
	//获取userID,方便之后创建JWT
	user.UserID = _user.UserID
	return
}
