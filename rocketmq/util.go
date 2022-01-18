package rocketmq

import "fmt"

func GetMqAddr() string {
	return fmt.Sprintf("%s:%d", IP, Port)
}
