package main

import (
	"interface/fedex"
	"interface/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &fedex.FedexSender{}
	var sender1 Sender = &koreaPost.PostSender{}
	SendBook("어린왕자", sender)
	SendBook("그리스인 조르바", sender)
	SendBook("어린왕자", sender1)
	SendBook("그리스인 조르바", sender1)

}
