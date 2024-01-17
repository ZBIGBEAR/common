package logger

import (
	"common/http"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

const (
	MaxMsgLength = 10000
	MsgTypeText  = "text"
)

type feiShuNotify struct {
	webhook string
	http    *http.HTTP
}

type MsgContent struct {
	Text string `json:"text"`
}

type LarkMessage struct {
	MsgType string      `json:"msg_type"`
	Content *MsgContent `json:"content"`
}

type Response struct {
	Extra         string
	StatusCode    int
	StatusMessage string
}

func newFeiShuNotify(webhook string) Notify {
	return &feiShuNotify{
		webhook: webhook,
		http:    http.Default(),
	}
}

func (f *feiShuNotify) Notify(ctx context.Context, msg string) error {
	if len(msg) >= MaxMsgLength {
		msg = msg[:MaxMsgLength]
	}
	req := &LarkMessage{
		MsgType: MsgTypeText,
		Content: &MsgContent{
			Text: msg,
		},
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return errors.Wrapf(err, "[SendMessage] Marshal")
	}
	resp, err := f.http.Do(ctx, "POST", f.webhook, reqBytes)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("closeErr:%v", closeErr)
		}
	}()

	result := &Response{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}
	if result.StatusCode != 0 {
		return errors.New(fmt.Sprintf("%v", result.StatusMessage))
	}

	return nil
}
