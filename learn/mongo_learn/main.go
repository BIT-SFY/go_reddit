package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 定义学生结构体
type Student struct {
	Name string
	Age  int
}

var client *mongo.Client
var db *mongo.Database
var collection *mongo.Collection

// 初始化数据库
func initDB() {
	// 设置客户端选项
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// 连接 MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("连接失败")
		log.Panic(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("连接成功")
	// 链接数据库test,并将链接返回给db
	db = client.Database("test")
	// 链接集合student,并将链接返回给collection
	collection = db.Collection("student")
}

// insertData 插入单条数据
func insertData() {
	// 初始化一个结构体,作为即将插入的数据
	s := Student{
		Name: "sfy",
		Age:  23,
	}
	// 将数据插入
	ior, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
	}
}

// insertManyData 插入多条数据
func insertManyData() {
	//初始化多条数据
	s1 := Student{
		Name: "szy",
		Age:  22,
	}
	s2 := Student{
		Name: "anw",
		Age:  22,
	}
	s3 := Student{
		Name: "yjw",
		Age:  22,
	}
	s4 := Student{
		Name: "nsdd",
		Age:  25,
	}
	s5 := Student{
		Name: "aktr",
		Age:  21,
	}
	// 声明成空接口类型的切片
	stus := []interface{}{s1, s2, s3, s4, s5}
	// 插入多条数据
	ior, err := collection.InsertMany(context.TODO(), stus)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ior.InsertedIDs: %v\n", ior.InsertedIDs...)
	}
}

// updateData 修改单条数据
func updateData() {
	// filter[筛选文档]: 查询到name=sfy的文档
	filter := bson.D{{Key: "name", Value: "sfy"}}
	// update[更新文档]: 修改name为SFY
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: "SFY"},
		}},
	}
	ur, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount) //更新条数
}

// updateManyData 更新多条数据
func updateManyData() {
	// 查询到所有age=23的文档
	filter := bson.D{{Key: "age", Value: 23}}
	// 修改age加一岁 $inc增加 $set设置成
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "age", Value: 2},
		}},
	}
	ur, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ur.ModifiedCount: %v\n", ur.ModifiedCount) //更新条数
}

// deleteData 删除单个文档
func deleteData() {
	filter := bson.D{{Key: "name", Value: "sfy"}}
	dr, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount) //删除条数
}

// deleteManyData 删除多个文档
func deleteManyData() {
	filter := bson.D{{Key: "age", Value: 22}}
	dr, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dr.DeletedCount: %v\n", dr.DeletedCount) //删除条数
}

// findData 查询单个文档
func findData() {
	var s Student
	// 查询name=sfy
	filter := bson.D{{Key: "name", Value: "sfy"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&s)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(s)
	}
}

// findManyData 查询多个文档
func findManyData() {
	//查找age=22
	filter := bson.D{{Key: "age", Value: 22}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	// 关闭上下文
	defer cursor.Close(context.TODO())
	// 定义切片
	var stus []Student
	err = cursor.All(context.TODO(), &stus)
	if err != nil {
		log.Fatal(err)
	}
	for _, stu := range stus {
		fmt.Println(stu)
	}
}

// findGroupData 复合查询
func findGroupData() {
	//定义最大时间
	opts := options.Aggregate().SetMaxTime(2 * time.Second)

	// 查询语句 age的和
	// groupStageSum := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$major"}, {Key: "ageSum", Value: bson.D{{Key: "$sum", Value: "$age"}}}}}}

	// 查询语句 age平均值
	groupStageAvg := bson.D{{
		Key: "$group", Value: bson.D{{
			Key: "_id", Value: "$major",
		},
			{Key: "ageAvg", Value: bson.D{{
				Key: "$avg", Value: "$age",
			}}},
		},
	}}
	// 查询语句 age最小值
	// groupStageMin := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$major"}, {Key: "ageMin", Value: bson.D{{Key: "$min", Value: "$age"}}}}}}

	// 查询语句 age最大值
	// groupStageMax := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$major"}, {Key: "ageMax", Value: bson.D{{Key: "$max", Value: "$age"}}}}}}

	// 查询
	result, err := collection.Aggregate(context.TODO(), mongo.Pipeline{groupStageAvg}, opts)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭上下文
	defer result.Close(context.TODO())

	// bson.M
	var stus []bson.M
	err = result.All(context.TODO(), &stus)
	if err != nil {
		log.Fatal(err)
	}

	// // 自定义切片
	// var stus []Student
	// err = result.All(context.TODO(), &stus)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for _, stu := range stus {
		fmt.Println(stu)
	}
}

func main() {
	initDB()
	defer client.Disconnect(context.TODO())
	// insertData()
	// insertManyData()
	// updateData()
	// updateManyData()
	// deleteData()
	// deleteManyData()
	// findData()
	findManyData()
	// findGroupData()
}
