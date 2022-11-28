package koreaPost

import "fmt"

type PostSender struct {
	// ...
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("우체국 send %s parcel\n", parcel)

}
