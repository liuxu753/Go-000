package main

import (
	"fmt"
	"net"
	"sync"
)

type Client struct {
	conn  net.Conn
	read  chan string
	write chan string
}

func (client *Client) close() {
	once := sync.Once{}
	once.Do(func() {
		client.conn.Close()
	})
}

func (client *Client) connWrite() {
	defer client.close()
	for true {
		select {
		case w := <-client.write:
			if w == "" {
				return
			}
			_, err := client.conn.Write([]byte(w))
			if err != nil {
				return
			}
		}
	}
}

func (client *Client) connRead() {
	defer client.close()
	defer close(client.read)
	for true {
		msg := make([]byte, 32)
		n, err := client.conn.Read(msg)
		if err != nil {

			return
		}
		fmt.Println(n)
		fmt.Println(string(msg))
		client.read <- string(msg)
	}
}

func (client *Client) connLogic() {
	defer client.close()
	defer close(client.write)
	//做一些鉴权操作
	//服务端是否使用某些校验逻辑，校验无效的tcp连接并主动关闭？
	for true {
		select {
		case msg := <-client.read:
			if msg == "" {
				return
			}
			client.write <- msg
		}
	}
}

func main() {
	server := net.TCPAddr{
		IP:   nil,
		Port: 12345,
		Zone: "",
	}
	listen, err := net.ListenTCP("tcp", &server)
	if err != nil {
		return
	}
	for true {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		client := Client{
			conn:  conn,
			read:  make(chan string, 10),
			write: make(chan string, 10),
		}
		go client.connWrite()
		go client.connRead()
		go client.connLogic()
	}
}
