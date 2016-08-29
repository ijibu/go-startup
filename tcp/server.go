package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

//用来记录所有的客户端连接
type ConnInfos struct {
	Conn map[string]*net.TCPConn
	lock sync.RWMutex
}

var ConnMap ConnInfos

func main() {
	var tcpAddr *net.TCPAddr
	ConnMap := &ConnInfos{Conn: make(map[string]*net.TCPConn)}

	tcpAddr, _ = net.ResolveTCPAddr("tcp", ":9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer func() {
		tcpListener.Close()
	}()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		// 新连接加入map
		ConnMap.lock.Lock()
		ConnMap.Conn[tcpConn.RemoteAddr().String()] = tcpConn
		ConnMap.lock.Unlock()
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for {
		message, _, err := reader.ReadRune()

		if err != nil {
			return
		}
		fmt.Println(message)
		// 这里返回消息改为了广播
		boradcastMessage(conn.RemoteAddr().String() + ":" + string(message))
	}
}

func boradcastMessage(message string) {
	b := []byte(message)
	// 遍历所有客户端并发送消息
	// 此处存在bug，会抛出panic("concurrent map read and map write")
	// http://studygolang.com/articles/2775 (优化 Go 中的 map 并发存取)
	ConnMap.lock.RLock()
	for _, conn := range ConnMap.Conn {
		conn.Write(b)
	}
	ConnMap.lock.RUnlock()
}
