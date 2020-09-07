package main

import "pi_common/communication/tcp"

func main() {
	//client := tcp.Client1{Address: "localhost:19010"}
	//client.InitClient()
	//client.Send("2300a78c070000000000000a00000000000000030000000000010618120250ba2700004eb0")
	tcp.ClientInit()
}
