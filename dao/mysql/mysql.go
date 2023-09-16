package mysql

import (
	"fmt"
	"reddit/settings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// 初始化mysql
func Init(cfg *settings.MySQLConfig) (err error) {
	//dsn，即data source name，指的是数据源名称。在golang中，其格式如下：
	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)

	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	DB, err := db.DB() //获取通用数据库对象
	if err != nil {
		panic(err)
	}

	//数据的大小设定需要根据具体情况而定
	DB.SetMaxOpenConns(cfg.MaxOpenConns) //用于设置连接池中空闲连接的最大数量。
	DB.SetMaxIdleConns(cfg.MaxIdleConns) //设置打开数据库连接的最大数量。
	return
}

func Close() {
	// 新版本中没有Close方法关闭连接，但可以利用通用的数据库对象关掉连接
	DB, err := db.DB() //获取通用数据库对象
	if err != nil {
		panic(err)
	}
	DB.Close()
}
