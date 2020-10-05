/**
Author: Mahmoud Ibrahim
github: https://github.com/mhmod1990
upworks: https://www.upwork.com/o/profiles/users/~01c04883943f6add8f/
*/

/**
second module that uses the platform to send and receive messages
*/
package main

import (
	"fmt"
	"gim"
)

// the structure type thaty represents the second module
type ModuleTwoAgent struct {
}

// OnReceive function whuch is required by the IAgent interface that process
// the received messages from the gim platform
func (a ModuleTwoAgent) OnReceive(im gim.IMessage) {
	switch im.MessageType {
	case NOTIFICATION_MESSAGE:
		NotificationMessage, _ := im.Payload.(NoficationMessage)
		fmt.Printf("[gimSample] module Two has OnReceive method with %s \n", NotificationMessage.data)
	}
}
