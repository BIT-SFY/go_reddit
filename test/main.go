package main

import (
	"reddit/logic"
	"reddit/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/reddit?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic("连接数据库失败!")
	}
	// InitUser()
	// InitCommunity()
	InitPost()
}

// 初始化用户表
func InitUser() {
	DB.AutoMigrate(&models.User{}) //创建表
}

// 初始化社区表
func InitCommunity() {
	DB.AutoMigrate(&models.Community{}) //创建表

	//插入数据
	c1 := models.Community{
		CommunityID:   1,
		CommunityName: "Beamng",
		Introduction:  "一个喜欢Beamng的玩家社区",
	}
	c2 := models.Community{
		CommunityID:   2,
		CommunityName: "Golang",
		Introduction:  "中文互联网最大的Go语言论坛!",
	}
	c3 := models.Community{
		CommunityID:   3,
		CommunityName: "Reddit",
		Introduction:  "论坛中的论坛,Reddit中的Reddit...",
	}
	c4 := models.Community{
		CommunityID:   4,
		CommunityName: "北京理工大学",
		Introduction:  "你说的对,但是北京理工大学是中国共产党创办的第一所理工科大学,隶属于中华人民共和国工业和信息化部,是全国重点大学。",
	}
	DB.Create(&c1)
	DB.Create(&c2)
	DB.Create(&c3)
	DB.Create(&c4)
}

// 初始化帖子表
func InitPost() {
	DB.AutoMigrate(&models.Post{}) //创建表
	//插入数据
	p1 := &models.Post{
		// PostID:      snowflake.GenID(),
		Title:       "Beamng最新MOD分享，奔驰S480迈巴赫精致模组",
		Content:     "本人转模后进行二次加工，下载链接：www.pan.baidu.co，车辆细节图以及车辆特点见楼下",
		AuthorID:    616223570137088,
		CommunityID: 1,
		Status:      1,
	}
	p6 := &models.Post{
		// PostID:      snowflake.GenID(),
		Title:       "【精品资源】外网大神最新发布兰博基尼-蝙蝠！速来手慢无！",
		Content:     "外网大神Trix今日收费模组惨遭泄露，我将在24h内删除模组链接，想下载的速来，这次蝙蝠的发布可谓是重磅炸弹，高达500M以及36种配置可选，以及上百种配件支持自定义，话不多说，快上车！！！<---↓↓↓链接在楼下↓↓↓--->",
		AuthorID:    616223570137088,
		CommunityID: 1,
		Status:      1,
	}
	p2 := &models.Post{
		// PostID:      snowflake.GenID(),
		Title:       "【精品资源】Golang仿reddit论坛",
		Content:     "里面包含go语言的系列教程以及golang后端的源码，想学习go语言的同学可以过来看看，学习完整个项目，代码能力绝对突飞猛进",
		AuthorID:    616359515918336,
		CommunityID: 2,
		Status:      1,
	}
	p7 := &models.Post{
		// PostID:      snowflake.GenID(),
		Title:       "求助！Golang学习遇到了一个问题",
		Content:     "今天在学习go语言的时候，出现了一个错误，我明明没有做错啊，请大佬看一下我的代码，帮我找找错误在哪",
		AuthorID:    616305577168896,
		CommunityID: 2,
		Status:      1,
	}
	p3 := &models.Post{
		// PostID:      snowflake.GenID(),
		Title:       "Reddit管理员须知",
		Content:     "1.不能骂人 2.不能涉密 3.不能涉恐 4.不能涉政",
		AuthorID:    616223570137088,
		CommunityID: 3,
		Status:      1,
	}
	p4 := &models.Post{
		// PostID:      snowflake.GenID(),
		Title:       "挂一个北理工的人",
		Content:     "今天在食堂碰到一个很恶心的人，他自己霸占的桌子，不给别人坐，还插队，我拍了他的照片，你们看看有认识他的没有",
		AuthorID:    616342638039040,
		CommunityID: 4,
		Status:      1,
	}
	p5 := &models.Post{
		// PostID:      snowflake.GenID(),
		Title:       "【二手交易】出售小牛电动车一辆!!!【捞一手】",
		Content:     "小牛电动车，2020年在小牛中关村南大街店买的，当时花了3800，今年要毕业了，车子8成新，1500出，可小刀，可面交，图片放在楼下了，我微信号:wechat996，非诚勿扰！！！",
		AuthorID:    616342638039040,
		CommunityID: 4,
		Status:      1,
	}
	logic.CreatePost(p1)
	logic.CreatePost(p2)
	logic.CreatePost(p3)
	logic.CreatePost(p4)
	logic.CreatePost(p5)
	logic.CreatePost(p6)
	logic.CreatePost(p7)
}
