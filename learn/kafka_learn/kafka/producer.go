package kafka

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

// 本实例展示的同步生产者的使用
func Pstart() {
	//生产10个消息
	Producer(10)
}

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
		str := "生产第" + strconv.Itoa(i) + "条消息"
		msg := &sarama.ProducerMessage{
			Topic: TOPIC,
			Key:   nil,
			Value: sarama.StringEncoder(str)}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("SendMessage:%d err:%v\n ", i, err)
			errors++
			continue
		}
		success++
		log.Printf("[Producer] partitionid: %d; offset:%d, value: %s\n", partition, offset, str)
	}
	log.Printf("发送完毕 总发送条数:%d successes: %d errors: %d\n", limit, success, errors)
}
