package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

var num *int = flag.Int("n", 1024, "please input a num like 1024")

func main() {
	flag.Parse()
	for i := 0; i < *num; i++ {
		go connect()
	}

	dur, err := time.ParseDuration("10s")
	if err != nil {
		fmt.Printf("[-] couldn't parse %s: %s\n", "10s", err.Error())
		return
	}
	//每隔10秒钟执行一次。
	for {
		<-time.After(dur)
	}
}

func connect() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("connected!")
	go onMessageRecived(conn)

	var fDumpCache string = "1s"
	if !(fDumpCache == "") {
		dur, err := time.ParseDuration(fDumpCache)
		if err != nil {
			fmt.Printf("[-] couldn't parse %s: %s\n", fDumpCache, err.Error())
			return
		}
		//每隔五秒钟发送一次消息。
		for {
			msg := time.Now().String()
			b := []byte(msg + "aaaaaaaaaaa\n")
			fmt.Println(string(b))
			conn.Write(b)
			//time.Sleep()表示休眠多少时间，休眠时处于阻塞状态，后续程序无法执行．
			//time.After()表示多少时间之后，但是在取出channel内容之前不阻塞，后续程序可以继续执行。
			<-time.After(dur)
		}
	}
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			break
		}
	}
}
