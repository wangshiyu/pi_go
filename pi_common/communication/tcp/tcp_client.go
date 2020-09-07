package tcp

import (
	"fmt"
	"pi_common/communication/tcp/stpro"
)

type Client2 struct {
	Phost  string
	Pmap   map[byte]string
	Client *stpro.Client
}

func (c Client2) Ptype(in []byte) {
	fmt.Printf("收到了type包的回复:%s\n", in)
}

func (c Client2) Pname(in []byte) {
	fmt.Printf("收到了name包的回复:%s\n", in)
}

func ClientInit() {
	var a1 = Client2{
		Phost: "localhost:9091",
		Pmap: map[byte]string{
			0x01: "type",
			0x02: "name",
		},
	}
	client, err := stpro.NewClient(a1)
	a1.Client = &client
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Send(0x02, []byte("jintianzhenhao"))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Send(0x01, []byte("jintianzhenhao3333"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
