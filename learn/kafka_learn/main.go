package main

import "reddit/learn/kafka_learn/kafka"

/*
生产者 producer -> 推送消息的工具
消费者 consumer -> 消费消息的地方
消费组 consumer-group -> 绑定多个topic 消费者绑定一个topic
*/

func main() {
	kafka.Producer(10)
	kafka.Consumer()
	// MyConsumerPartition(0)
}

/*

kafka学习:
需要安装Zookeeper
https://www.bilibili.com/video/BV1xD4y1173J/
再安装kafka
https://www.bilibili.com/video/BV1QW4y1v7kn/

根据配置文件server.properties来启动kafka
.\bin\windows\kafka-server-start.bat .\config\server.properties
创建一个名为 "test" 的主题，其中包含 1 个分区和 1 个副本[但是因为本地只有一个kafka服务器，所以副本只能为1]
.\bin\windows\kafka-topics.bat --create --topic test --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092
创建一个名为 "my_topic" 的主题，其中包含 3 个分区和 1 个副本[但是因为本地只有一个kafka服务器，所以副本只能为1]
.\bin\windows\kafka-topics.bat --create --topic my_topic --partitions 3 --replication-factor 1 --bootstrap-server localhost:9092
显示有哪些topic
.\bin\windows\kafka-topics.bat -list --bootstrap-server localhost:9092
显示某个topic具体内容
.\bin\windows\kafka-topics.bat  --describe --topic learn_kafka --bootstrap-server localhost:9092
删除一个topic,以删除my_topic为例: 详情见：https://blog.csdn.net/lzl_anonym/article/details/118577002
1.首先删除log文件中的my_topic相关文件夹
2.在zkCli.cmd中进行如下操作
[zk: localhost:2181] ls /brokers/topics
-> [__consumer_offsets, my_topic]
[zk: localhost:2181] deleteall /brokers/topics/my_topic
[zk: localhost:2181] ls /brokers
-> [__consumer_offsets]

开启一个生产者
.\bin\windows\kafka-console-producer.bat --broker-list localhost:9092 --topic test
开启一个消费者
.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic test --from-beginning

*/
