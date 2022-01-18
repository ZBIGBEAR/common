package rocketmq

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var adm admin.Admin

type localAdmin struct {
	admin.Admin
}

func GetMqAdmin() admin.Admin {
	return adm
}

func init() {
	tmpAdm, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{GetMqAddr()})))
	if err != nil {
		panic(err)
	}

	adm = tmpAdm
}

func CreateTopic(ctx context.Context, topic string) error {
	brokerAddr := "127.0.0.1:10911" // or "127.0.0.1:10909"
	return GetMqAdmin().CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)
}

func DeleteTopic(ctx context.Context, topic string) string {
	return ""
}
