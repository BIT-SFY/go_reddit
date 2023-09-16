package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"reddit/models"

	"gorm.io/gorm"
)

//把每一步数据库操作封装成函数 待Logic层根据业务需求调用

// CheckUserExist 根据用户名检查该用户是否存在
func CheckUserExist(username string) (err error) {
	user := models.User{}
	if err = db.Where("username = ?", username).Find(&user).Error; err != gorm.ErrRecordNotFound && err != nil {
		return err
	}
	if user.ID != 0 {
		return errors.New("该用户已存在")
	}
	return nil
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

const secret = "3220232821"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
