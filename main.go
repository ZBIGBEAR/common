package main

import (
	"common/rocketmq"
	"common/rocketmq/handler"
	"context"
	"fmt"
)

func main() {
	// 1.注册topic
	topic := "testtopic"
	if err := rocketmq.CreateTopic(context.Background(), topic); err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("CreateTopic success. topic:%s", topic))
	}

	/* 输出
	INFO[0000] create topic success                          broker="127.0.0.1:10911" topic=testtopic
	CreateTopic success. topic:testtopic
	*/

	// 2.生产消息
	// 初始化producer
	if err := rocketmq.InitProducer(); err != nil {
		panic(err)
	} else {
		fmt.Println("InitProducer success")
	}

	msg := "hello, testtopic"
	// 生产消息
	if err := rocketmq.SendMessage(context.Background(), topic, msg); err != nil {
		panic(err)
	} else {
		fmt.Println(fmt.Sprintf("SendMessage success. topic:%s, msg:%s", topic, msg))
	}

	// 3.消费消息
	// 注册消息处理函数
	rocketmq.Register(topic, handler.HandleHello)
	// 初始化consumer
	if err := rocketmq.InitConsumer(); err != nil {
		panic(err)
	} else {
		fmt.Println("InitConsumer success.")
	}
}
