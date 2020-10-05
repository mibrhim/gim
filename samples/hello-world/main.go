/**
Author: Mahmoud Ibrahim
github: https://github.com/mhmod1990
upworks: https://www.upwork.com/o/profiles/users/~01c04883943f6add8f/
*/

//package main for the first sample that is used to try the gim package
//and to understand how it is works
package main

import (
	"fmt"
	"gim"
	"time"
)

// define the constants which represents the different modules
// and the different messages which will be used inside the system
const (
	MODULE_ONE int = 1
	MODULE_TWO     = 2

	NOTIFICATION_MESSAGE = 1
)

// define the new custom internal message payload
// this types will be sent inside the internal message to carry data
// through the different modules of the system.
type NoficationMessage struct {
	data string
}

// The main function which starts the sample application
// it creates the different modules and register them to the system
// and then it initiates the system by sending the first message
func main() {
	fmt.Printf("[gimSample] Starting Test Program\n")

	// create the module variables
	var moduleOneAgent ModuleOneAgent
	var moduleTwoAgent ModuleTwoAgent

	// register the modules to the platform
	gim.Register(MODULE_ONE, moduleOneAgent)
	gim.Register(MODULE_TWO, moduleTwoAgent)

	// create the data which required to be sent to module one
	notif := NoficationMessage{
		data: "hello world",
	}

	// create the internal message which will carry the data
	im := gim.IMessage{
		Receiver:    MODULE_ONE,
		Sender:      MODULE_TWO, // spoof module two identity
		MessageType: NOTIFICATION_MESSAGE,
		Payload:     notif,
	}

	// send this message to the receiver module through the platform
	gim.Send(im)

	// sleep the main process for 2 seconds.
	time.Sleep(2 * time.Second)
}
