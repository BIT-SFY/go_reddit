package kafka

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

var HOST = "localhost:9092"
var TOPIC = "my_topic"
var KEY = "whoiam"

type msgStruct struct {
	Id      int
	Name    string
	Content string
}

// 本实例展示的同步生产者的使用
func Producer(limit int) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{HOST}, config)
	if err != nil {
		log.Fatal("NewSyncProducer err:", err)
	}
	defer producer.Close()
	for i := 0; i < limit; i++ {
		name := "第" + strconv.Itoa(i) + "个人"
		content := "发送的第" + strconv.Itoa(i) + "个" + KEY + "消息"
		msgInfo := msgStruct{
			Id:      i,
			Name:    name,
			Content: content,
		}
		// 对要发送的数据进行编码
		msgStructJson, _ := json.Marshal(msgInfo)
		msgStructBase64 := base64.StdEncoding.EncodeToString(msgStructJson)

		msg := &sarama.ProducerMessage{
			Topic: TOPIC,                               // 消息的主题
			Key:   sarama.StringEncoder(KEY),           // 根据key的值hash出一个partition
			Value: sarama.ByteEncoder(msgStructBase64), // 消息内容
		}
		// 发送消息
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalf("send message(%s) fail, err=%s \n", string(msgStructJson), err)
		} else {
			log.Printf("send message(%s) success, part=%d, offset=%d \n", string(msgStructJson), partition, offset)
		}
	}
	log.Printf("发送完毕 总发送条数:%d \n", limit)
}

/*

那在kafka中，如果某个topic有多个partition，producer又怎么知道该将数据发往哪个partition呢?kafka中有几个原则:
	1.partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition。
	2.如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition。
	3.如果既没指定partition，又没有设置key，则会采用轮询方式，即每次取一小段时间的数据写入某个partition，下一小段的时间写入下一个partition。

*/
