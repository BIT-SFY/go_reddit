package kafka

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

var wg sync.WaitGroup

// 消费
func Consumer() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{HOST}, config)
	if err != nil {
		log.Fatal("NewConsumer err:", err)
	}
	defer consumer.Close()
	// 参数1 指定消费那个topic
	// 参数2 分区，这里默认信0号分区
	// 参数3 offset 从哪儿开始消费起走
	// 如果改为 sarama.OffsetOldest 则会从最旧的消息开始消费，即每次重启 consumer 都会把该 topic 下的所有消息消费一次
	partitions, err := consumer.Partitions(TOPIC)
	if err != nil {
		log.Printf("return topic partitions error %s\n", err.Error())
		return
	}
	for _, partitionId := range partitions {
		PartitionConsumer, err := consumer.ConsumePartition(TOPIC, partitionId, sarama.OffsetOldest)
		if err != nil {
			log.Printf("try create partition_consumer error %s\n", err.Error())
			return
		}
		wg.Add(1)

		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for message := range pc.Messages() {
				msg := handlerReaderMsg(message)
				log.Printf("[Consumer] partitionid: %d; offset:%d, value: %v\n", message.Partition, message.Offset, *msg)
			}
		}(PartitionConsumer)
	}
	wg.Wait()
}

// 处理kafka传来的base64数据
func handlerReaderMsg(message *sarama.ConsumerMessage) *msgStruct {
	msgInfo, err := base64.StdEncoding.DecodeString(string(message.Value))
	if err != nil {
		log.Print(err)
	}

	// 将解码后的数据反序列化为 msgStruct 结构体
	msg := &msgStruct{}
	err = json.Unmarshal(msgInfo, &msg)
	if err != nil {
		log.Print(err)
	}
	return msg
}
