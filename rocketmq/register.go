package rocketmq

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type HandlerFunc func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)

var (
	handlerMap map[string]HandlerFunc
)

func init() {
	handlerMap = make(map[string]HandlerFunc)
}

func Register(topic string, handlerFunc HandlerFunc) {
	if _, ok := handlerMap[topic]; ok {
		panic(errors.New(fmt.Sprintf("topic:%s exists", topic)))
	}

	handlerMap[topic] = handlerFunc
}
