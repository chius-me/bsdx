package handler

import (
	"errors"

	"github.com/bytedance/sonic"
	"github.com/gookit/event"
	"github.com/kaiheila/golang-bot/api/base"
	event2 "github.com/kaiheila/golang-bot/api/base/event"
	"github.com/kaiheila/golang-bot/api/helper"
	log "github.com/sirupsen/logrus"
)

type EchoHandler struct {
	Token   string
	BaseUrl string
}

func (h *EchoHandler) Handle(e event.Event) error {
	err := func() error {
		frameVal, ok := e.Data()[base.EventDataFrameKey]
		if !ok {
			return errors.New("event data has no frame field")
		}
		frame, ok := frameVal.(*event2.FrameMap)
		if !ok {
			return errors.New("frame is not FrameMap type")
		}
		data, err := sonic.Marshal(frame.Data)
		if err != nil {
			return err
		}
		msgEvent := &event2.MessageKMarkdownEvent{}
		if err := sonic.Unmarshal(data, msgEvent); err != nil {
			return err
		}

		if msgEvent.Author.Bot {
			log.Debug("Ignoring bot message")
			return nil
		}

		log.Infof("Received message: channel=%s, author=%s, content=%s",
			msgEvent.TargetId, msgEvent.Author.Username, msgEvent.KMarkdown.RawContent)

		client := helper.NewApiHelper("/v3/message/create", h.Token, h.BaseUrl, "", "")
		echoData := map[string]string{
			"type":      "9",
			"target_id": msgEvent.TargetId,
			"content":   "echo: " + msgEvent.KMarkdown.RawContent,
		}
		echoBytes, err := sonic.Marshal(echoData)
		if err != nil {
			return err
		}
		resp, err := client.SetBody(echoBytes).Post()
		if err != nil {
			log.WithError(err).Error("Failed to send echo message")
			return err
		}
		log.Infof("Echo sent, response: %s", string(resp))
		return nil
	}()
	if err != nil {
		log.WithError(err).Error("EchoHandler error")
	}
	return nil
}