package main

import "reddit/learn/kafka_learn/kafka"

func main() {
	kafka.Pstart()
	kafka.Cstart()
}

/*

kafka学习:
需要安装Zookeeper
https://www.bilibili.com/video/BV1xD4y1173J/
再安装kafka
https://www.bilibili.com/video/BV1QW4y1v7kn/

根据配置文件对kafka进行初始化
.\bin\windows\kafka-server-start.bat .\config\server.properties
创建一个topic: test
.\bin\windows\kafka-topics.bat --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test
显示有哪些topic
.\bin\windows\kafka-topics.bat -list --bootstrap-server localhost:9092
开启一个生产者
.\bin\windows\kafka-console-producer.bat --broker-list localhost:9092 --topic test
开启一个消费者
.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic test --from-beginning
*/
