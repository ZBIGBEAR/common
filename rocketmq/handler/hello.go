package handler

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func HandleHello(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	fmt.Println("HandleHello recv msgs")
	for i := range msgs {
		fmt.Printf("subscribe callback: %v \n", msgs[i])
	}

	return consumer.ConsumeSuccess, nil
}
