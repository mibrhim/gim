package gim

import (
	"fmt"
)

type IAgent interface {
	OnReceive(im IMessage)
}

type imodule interface {
	run(messages chan IMessage)
}

type Module struct {
	module int
	agent  IAgent
}

var modules map[int]chan IMessage = make(map[int]chan IMessage)

func OnReceive(IMessage) {
	fmt.Printf("[gim] This module doesn't have OnReceive method\n")
}

func (m Module) run(messages chan IMessage) {
	for true {
		m.agent.OnReceive(<-messages)
	}
}

func Register(module int, agent IAgent) {
	modules[module] = make(chan IMessage)
	container := Module{
		module: module,
		agent:  agent,
	}
	go container.run(modules[module])

}

func Send(im IMessage) {
	if modules[im.Receiver] == nil {
		fmt.Printf("[gim] Module %d is not registered to the platform\n", im.Receiver)
	}

	modules[im.Receiver] <- im
}

func Unregister(module int) {
	delete(modules, module)
}
