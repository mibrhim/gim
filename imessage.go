/**
Author: Mahmoud Ibrahim
github: https://github.com/mhmod1990
upworks: https://www.upwork.com/o/profiles/users/~01c04883943f6add8f/
*/

package gim

// the payload interface which should act as a generic type to carry
// the different payload structure through the gim platform
type IPayload interface {
}

// the internal message structure which is used to deliver the message
// from module to the other through the gim platform
type IMessage struct {
	Receiver    int
	Sender      int
	MessageType int

	Payload IPayload // custom message
}
