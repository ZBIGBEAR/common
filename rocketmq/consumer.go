package rocketmq

import (
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/pkg/errors"
)

var (
	mqConsumer rocketmq.PushConsumer
)

type localConsumer struct {
	rocketmq.PushConsumer
}

func GetConsumer() rocketmq.PushConsumer {
	return mqConsumer
}

func InitConsumer() error {
	// init consumer
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{GetMqAddr()})),
	)

	if err != nil {
		return err
	}

	mqConsumer = &localConsumer{c}

	// subscribe
	return subscribe()
}

func subscribe() error {
	for topic, handlerFunc := range handlerMap {
		if err := GetConsumer().Subscribe(topic, consumer.MessageSelector{}, handlerFunc); err != nil {
			return err
		} else {
			fmt.Println(fmt.Sprintf("topic:%s register success", topic))
		}
	}

	// 先订阅，后启动
	return GetConsumer().Start()
}

func Stop() error {
	if mqConsumer == nil {
		return errors.New("producer not init")
	}

	return mqConsumer.Shutdown()
}
