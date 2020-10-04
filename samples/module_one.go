package main

import (
	"fmt"
	"gim"
)

type ModuleOneAgent struct {
}

func (a ModuleOneAgent) OnReceive(im gim.IMessage) {

	switch im.MessageType {
	case NOTIFICATION_MESSAGE:
		NotificationMessage, _ := im.Payload.(NoficationMessage)
		fmt.Printf("[gimTest] module One has OnReceive method with %s \n", NotificationMessage.data)
	}

	x := gim.IMessage{
		Receiver:    MODULE_TWO,
		Sender:      MODULE_ONE,
		MessageType: NOTIFICATION_MESSAGE,
		Payload:     NoficationMessage{data: "Hello Back"},
	}

	gim.Send(x)
}
