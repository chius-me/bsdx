package main

import (
	"bsdx/config"
	"bsdx/handler"

	"github.com/kaiheila/golang-bot/api/base"
	"github.com/kaiheila/golang-bot/api/helper/compress"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	session := base.NewWebSocketSession(
		config.Token,
		config.BaseUrl,
		"./session.pid",
		"",
		0,
		compress.CompressTypeNone,
		"",
		1,
	)

	session.On(base.EventReceiveFrame, &handler.ReceiveFrameHandler{})
	session.On("GROUP_9", &handler.EchoHandler{
		Token:   config.Token,
		BaseUrl: config.BaseUrl,
	})

	session.Start()
}