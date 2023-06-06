package dingtalk

import (
	"github.com/axelwong/CqepcAuto/global"
	"github.com/blinkbean/dingtalk"
	"go.uber.org/zap"
)

type Message struct {
	Token  string
	Secret string
	Phone  []string
}

func (m *Message) New(message Message) *Message {
	return &message
}

// SendText 发送文本消息
func (m *Message) SendText(msg string) error {
	client := dingtalk.InitDingTalkWithSecret(m.Token, m.Secret)
	err := client.SendTextMessage(msg, dingtalk.WithAtMobiles(m.Phone))
	if err != nil {
		global.APP_LOG.Error("钉钉发送文本消息失败", zap.Error(err))
		return err
	}
	return nil
}

// SendMarkdown 发送markdown消息
func (m *Message) SendMarkdown(title string, msg []string) error {
	client := dingtalk.InitDingTalkWithSecret(m.Token, m.Secret)
	err := client.SendMarkDownMessageBySlice(title, msg, dingtalk.WithAtMobiles(m.Phone))
	if err != nil {
		global.APP_LOG.Error("钉钉发送markdown消息失败", zap.Error(err))
		return err
	}
	return nil
}
