/**
Author: Mahmoud Ibrahim
github: https://github.com/mhmod1990
upworks: https://www.upwork.com/o/profiles/users/~01c04883943f6add8f/
*/

/**
Example on how to use this library with the http server to run any tasks
on a background module
*/
package main

import (
	"fmt"
	"gim"
	"net/http"
)

// define the constants which represents the different modules
// and the different messages which will be used inside the system
const (
	MODULE_ONE  int = 1
	MODULE_MAIN     = 2

	NEW_REQUEST_MESSAGE = 1
)

// define the new custom internal message payload
// this types will be sent inside the internal message to carry data
// through the different modules of the system.
type NewRequestMessage struct {
	data string
}

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {

	// Functions serving as handlers take a
	// `http.ResponseWriter` and a `http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here our simple response is just
	// "hello\n".
	fmt.Fprintf(w, "hello\n")

	// create the data which required to be sent to module one
	notification := NewRequestMessage{
		data: "hello world",
	}

	// create the internal message which will carry the data
	im := gim.IMessage{
		Receiver:    MODULE_ONE,
		Sender:      MODULE_MAIN, // spoof module two identity
		MessageType: NEW_REQUEST_MESSAGE,
		Payload:     notification,
	}

	// send this message to the receiver module through the platform
	gim.Send(im)
}

func main() {

	// create the module variables
	var moduleOneAgent ModuleOneAgent

	// register the modules to the platform
	gim.Register(MODULE_ONE, moduleOneAgent)

	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.
	http.HandleFunc("/hello", hello)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.
	http.ListenAndServe(":8090", nil)
}
