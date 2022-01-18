package rocketmq

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var (
	mqProducer rocketmq.Producer
)

type localProducer struct {
	rocketmq.Producer
}

func GetProducer() rocketmq.Producer {
	return mqProducer
}

func InitProducer() error {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{GetMqAddr()})),
		// producer.WithNameServer(endPoint),
		//producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(Retry),
		//producer.WithGroupName("GID_xxxxxx"),
	)

	if err != nil {
		return err
	}

	mqProducer = &localProducer{p}

	return mqProducer.Start()
}

func StopProducer() error {
	if mqProducer == nil {
		return errors.New("producer not init")
	}

	return mqProducer.Shutdown()
}

func SendMessage(ctx context.Context, topic, msg string) error {
	if mqProducer == nil {
		return errors.New("producer not init")
	}

	mqMsg := &primitive.Message{
		Topic: topic,
		Body:  []byte(msg),
	}
	res, err := mqProducer.SendSync(context.Background(), mqMsg)
	if err != nil {
		return err
	} else {
		fmt.Println(res.String())
		return nil
	}
}
