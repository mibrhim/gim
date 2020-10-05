/**
Author: Mahmoud Ibrahim
github: https://github.com/mhmod1990
upworks: https://www.upwork.com/o/profiles/users/~01c04883943f6add8f/
*/

// gim (go internal Messaging) package which gives anyone
// the ability to create a module and send different messages
// between each other.
// .. this module act as a platform for message delivery between modules
package gim

import (
	"fmt"
)

// the IAgent interface which must be implemented by the modules
// it is considered as the entry point to the module
type IAgent interface {
	OnReceive(im IMessage)
}

// it is internal interface which used to create the contianer which
// the module will run inside
type imodule interface {
	run(messages chan IMessage)
}

// it is the container of the modules agent
// it uses the module agent to send the message to the right module
type Module struct {
	module int
	agent  IAgent
}

// map to store the chan and connect it with the modules which
// use it.
var modules map[int]chan IMessage = make(map[int]chan IMessage)

// implementation of the run function of the imodule interface
// that could give us the ability to create a container to each agent
func (m Module) run(messages chan IMessage) {
	for true {
		// send the message to the agent which lives inside this container
		m.agent.OnReceive(<-messages)
	}
}

// add new module to the platform using the module unique id
// and the agent which is the entry point to the module
func Register(module int, agent IAgent) {
	// prepare the container for the agent of the module
	modules[module] = make(chan IMessage)
	container := Module{
		module: module,
		agent:  agent,
	}
	// start the new go routine for the container where this agent
	// will live inside
	go container.run(modules[module])

}

// send internal message using the gim platform between different module
func Send(im IMessage) {
	if modules[im.Receiver] == nil {
		fmt.Printf("[gim] Module %d is not registered to the platform\n", im.Receiver)
	}

	// send the message to the correspont agent
	modules[im.Receiver] <- im
}

// unregister and remove a module from the gim platform
// so that this module can't receive any messages
func Unregister(module int) {
	delete(modules, module)
}
