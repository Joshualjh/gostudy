package main

import (
	"interface/fedex"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &fedex.FedexSender{}
	SendBook("어린왕자", sender)
	SendBook("그리스인 조르바", sender)
}
