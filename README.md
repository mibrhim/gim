# gim
Golang Internal Messaging system, which supports internal messaging between different modules where each module run on a different go routine and waits for any internal message

## Why gim
you can use gim if you have a different domain that you want to decouple, you can use the internal message to start a certain procedures in each domain. this library best fits if you are creating a monolith application or it is a micro-service which have different components that

## How to use
to use this library you can import it using 
```go
import (
  "fmt"

	"github.com/mhmod1990/gim"
)
```
and then you have to create the modules (domains) which you want to add to the platform, this module must implement the `IAgent` interface function `OnReceive(im gim.IMessage` such as:
```go
// Module one type
type ModuleOneAgent struct {
}

// implement the IAgent interface OnReceive func
func (a ModuleOneAgent) OnReceive(im gim.IMessage) {

	switch im.MessageType {
	case NOTIFICATION_MESSAGE:
	  ...
  }
  ...
}
```
Now, you can add this module to the platform and send an internal message to it using:
```go
const (
	MODULE_ONE int = 1
  MODULE_TWO     = 2
  ...

	NOTIFICATION_MESSAGE = 1
)

func main() {
  // create an object of module one
  var moduleOneAgent ModuleOneAgent

  // register the module
  gim.Register(MODULE_ONE, moduleOneAgent)
  
  // create the internal message
  im := gim.IMessage{
		Receiver:    MODULE_ONE,
		Sender:      MODULE_TWO,
		MessageType: NOTIFICATION_MESSAGE,
	}
  
  // send internal message
  gim.Send(im)
}

```
### Custom Internal Messages
you can also add any object which is shared between the different module to be sent inside the internal message as a payload, you can use any type or struct as a Payload, However you will need to cast it back to its original type which you receive it in the next module
**NOTE** take care of the type conversion when you receive this message
```go
// new custom payload
type NoficationMessage struct {
	data string
}


func main(){
  ...
  
  // create object from the custom payload and set its internal values
  notification := NoficationMessage{
		data: "hello world",
	}
  
  // create new internal message with the custom payload
  im := gim.IMessage{
		Receiver:    MODULE_ONE,
		Sender:      MODULE_TWO,
		MessageType: NOTIFICATION_MESSAGE,
		Payload:     notification,
	}
  
  // send internal message
  gim.Send(im)
  
  ...
}
```
In the `OnReceive(im gim.IMessage)` function of the receiver module, you will get this internal message then you can check the type and restore 
the original payload type like this:
```go
func (a ModuleOneAgent) OnReceive(im gim.IMessage) {

	switch im.MessageType {
	case NOTIFICATION_MESSAGE:
    // restore the original type of the payload and check assert that it was the right type
		NotificationMessage, ok := im.Payload.(NoficationMessage)
    
    // check the conversion of the payload to its original type
    if !ok {
			fmt.Printf("[gimSample] Wrong type for message %d \n", im.MessageType)
			return
		}
		fmt.Printf("[gimTest] module one has OnReceive method with %s \n", NotificationMessage.data)
	}
  
  ...
}

```

## Finally
Kindly, contribute and report issues to this library. I still working on it to improve the exception handling and readability.
Thanks :).

