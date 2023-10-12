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
	var success, errors int
	for i := 0; i < limit; i++ {
		name := "第" + strconv.Itoa(i) + "个人"
		content := "发送的第" + strconv.Itoa(i) + "个消息"
		msgInfo := msgStruct{
			Id:      i,
			Name:    name,
			Content: content,
		}
		// 对要发送的数据进行编码
		msgStructJson, _ := json.Marshal(msgInfo)
		msgStructBase64 := base64.StdEncoding.EncodeToString(msgStructJson)

		msg := &sarama.ProducerMessage{
			Topic: TOPIC,
			Key:   nil,
			Value: sarama.ByteEncoder(msgStructBase64),
		}
		// 发送消息
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalf("send message(%s) fail, err=%s \n", string(msgStructJson), err)
		} else {
			log.Printf("send message(%s) success, part=%d, offset=%d \n", string(msgStructJson), partition, offset)
		}
	}
	log.Printf("发送完毕 总发送条数:%d successes: %d errors: %d\n", limit, success, errors)
}
