/*
 * @Author: BIT-SFY
 * @Date: 2023-09-14 15:00:24
 * @LastEditTime: 2023-09-14 15:00:35
 * @Description:
 */
package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// 初始化mysql
func initMySQL() (err error) {
	//dsn，即data source name，指的是数据源名称。在golang中，其格式如下：
	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	dsn := "root:123456@tcp(127.0.0.1:3306)/reddit?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{ //并没有实际的和数据库建立连接
		Logger: logger.Default.LogMode(logger.Info), //配置日志级别，打印出所有的sql
	})
	if err != nil {
		panic(err)
	}
	// 新版本中没有Close方法关闭连接，有数据库连接池维护连接信息，也可以利用通用的数据库对象关掉连接
	DB, err := db.DB() //获取通用数据库对象
	if err != nil {
		panic(err)
	}
	//数据的大小设定需要根据具体情况而定
	DB.SetMaxOpenConns(200)          //用于设置连接池中空闲连接的最大数量。
	DB.SetMaxIdleConns(10)           //设置打开数据库连接的最大数量。
	DB.SetConnMaxIdleTime(time.Hour) //设置了连接可复用的最大时间。
	// defer DB.Close()                        //通过常规数据库接口关闭
	return
}

func main() {
	if err := initMySQL(); err != nil {
		fmt.Println("mysql start failed...", err)
	}
	fmt.Println("connect mysql successed!")
}
