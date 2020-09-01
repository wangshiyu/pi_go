package tcp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
)

type DataPacket struct {
	Type string
	Body string
}

func ServerInit() {
	//绑定端口
	var tcpAddr, err = net.ResolveTCPAddr("tcp", ":19010")
	if err != nil {
		fmt.Println(err.Error())
	}
	//监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()
	//开始接收数据
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go Handler(conn)
	}

}

func Handler(conn net.Conn) {
	defer conn.Close()
	//每次读取数据长度
	buf := make([]byte, 256)
	_, err := conn.Read(buf)
	if err != nil {
		return
	}
	result, Body := check(buf)
	if result {
		fmt.Printf("接收到报文内容:{%s}\n", hex.EncodeToString(Body))
	}

}

func check(buf []byte) (bool, []byte) {
	Length := DataLength(buf)
	if Length < 3 || Length > 256 {
		return false, nil
	}
	Body := buf[:Length]
	return uint16(len(Body))-2 != Length, Body
}

func DataLength(buf []byte) uint16 {
	return binary.BigEndian.Uint16(inversion(buf[:2])) + 2
}

//反转字节
func inversion(buf []byte) []byte {
	for i := 0; i < len(buf)/2; i++ {
		temp := buf[i]
		buf[i] = buf[len(buf)-1-i]
		buf[len(buf)-1-i] = temp
	}
	return buf
}
