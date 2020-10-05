/**
Author: Mahmoud Ibrahim
github: https://github.com/mhmod1990
upworks: https://www.upwork.com/o/profiles/users/~01c04883943f6add8f/
*/

/**
first module that uses the platform to send and receive messages
*/
package main

import (
	"fmt"
	"gim"
)

// create the struct type that represents the ModuleOneAgent
type ModuleOneAgent struct {
}

// implement the OnReceive funtion which is required by the gim platform
// interface IAgent
func (a ModuleOneAgent) OnReceive(im gim.IMessage) {

	//switch on the internal message type
	switch im.MessageType {
	// if the message is a Notification message
	case NEW_REQUEST_MESSAGE:
		// convert the payload to the structure of the notification message
		NotificationMessage, _ := im.Payload.(NewRequestMessage)
		fmt.Printf("[gimSample] module One has OnReceive method with %s \n", NotificationMessage.data)
	}
}
