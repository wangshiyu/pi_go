package main

import (
	"pi_common/communication/tcp"
	"time"
)

func main() {
	tcp.InitTcpServerCore()
	for true {
		time.Sleep(10000)
	}
}
