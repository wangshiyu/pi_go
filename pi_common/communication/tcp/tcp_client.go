package tcp

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
)

/**
客户端对象
*/
type Client struct {
	Address    string
	connection *net.TCPConn
	server     *net.TCPAddr
	stopChan   chan struct{}
}

func (client *Client) InitClient() {
	if client.Address == "" {
		errors.New("io: address is empty")
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", client.Address)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	connection, err := net.DialTCP("tcp", nil, tcpAddr)
	client.connection = connection
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

/**
接收数据包
*/
func (client *Client) receivePackets() {
	reader := bufio.NewReader(client.connection)
	for {
		//承接上面说的服务器端的偷懒，我这里读也只是以\n为界限来读区分包
		msg, err := reader.ReadString('\n')
		if err != nil {
			//在这里也请处理如果服务器关闭时的异常
			close(client.stopChan)
			break
		}
		fmt.Print(msg)
	}
}

func (client *Client) Send(msg string) {
	if client.connection == nil {
		fmt.Errorf("client connection is nil")
		return
	}
	defer client.connection.Close()
	decodeString, _ := hex.DecodeString(msg)
	_, err := client.connection.Write(decodeString)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
