package main

import (
	"fmt"
	"gim"
)

type ModuleTwoAgent struct {
}

func (a ModuleTwoAgent) OnReceive(im gim.IMessage) {
	switch im.MessageType {
	case NOTIFICATION_MESSAGE:
		NotificationMessage, _ := im.Payload.(NoficationMessage)
		fmt.Printf("[gimTest] module Two has OnReceive method with %s \n", NotificationMessage.data)
	}
}
