package handler

import (
	"github.com/gookit/event"
	log "github.com/sirupsen/logrus"
)

type ReceiveFrameHandler struct{}

func (rf *ReceiveFrameHandler) Handle(e event.Event) error {
	log.WithField("event", e.Name()).Debug("Received frame event")
	return nil
}