package gim

type IPayload interface {
}

type IMessage struct {
	Receiver    int
	Sender      int
	MessageType int

	Payload IPayload
}
