package main

import (
	"fmt"
	"gim"
	"time"
)

const (
	MODULE_ONE int = 1
	MODULE_TWO     = 2

	NOTIFICATION_MESSAGE = 1
)

type NoficationMessage struct {
	data string
}

func main() {
	fmt.Printf("[gimTest] Starting Test Program\n")

	var moduleOneAgent ModuleOneAgent
	var moduleTwoAgent ModuleTwoAgent
	gim.Register(MODULE_ONE, moduleOneAgent)
	gim.Register(MODULE_TWO, moduleTwoAgent)

	notif := NoficationMessage{
		data: "hello world",
	}

	im := gim.IMessage{
		Receiver:    MODULE_ONE,
		Sender:      MODULE_TWO,
		MessageType: NOTIFICATION_MESSAGE,
		Payload:     notif,
	}

	gim.Send(im)

	time.Sleep(2 * time.Second)
}
