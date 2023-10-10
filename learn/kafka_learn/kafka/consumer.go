package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

var HOST = "127.0.0.1:9092"
var TOPIC = "learn_kafka"

func Cstart() {
	//单分区消费
	SinglePartition()
}

func SinglePartition() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{HOST}, config)
	if err != nil {
		log.Fatal("NewConsumer err:", err)
	}
	defer consumer.Close()
	// 参数1 指定消费那个topic
	// 参数2 分区，这里默信0号分区
	// 参数3 offset 从哪儿开始消费起走，正常情况下每次消费完都会将这次的offset提交到kafka
	// 如果改为 sarama.OffsetOldest 则会从最旧的消息开始消费，即每次重启 consumer 都会把该 topic 下的所有消息消费一次
	partitionConsumer, err := consumer.ConsumePartition(TOPIC, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("ConsumerParition err:", err)
	}
	defer partitionConsumer.Close()
	// 会一直阻塞在这里
	for message := range partitionConsumer.Messages() {
		log.Printf("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
	}
}
