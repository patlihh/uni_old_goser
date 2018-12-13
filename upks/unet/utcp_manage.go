// utcp_manage
package unet

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	MAX_CONN_NUM = 10000
)

//var cur_conn_num int = 0
//var conn_chan = make(chan net.Conn)
//var ch_conn_change := make(chan int)
var conns_chan = make(chan net.Conn, MAX_CONN_NUM)

func TcpManage() {

	fmt.Println("tcp_manage start....")

	listener, err := net.Listen("tcp", "0.0.0.0:8088")
	if err != nil {
		fmt.Println("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	var cur_conn_num int = 0
	conn_chan := make(chan net.Conn)
	ch_conn_change := make(chan int)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
		}
	}()

	go func() {
		for _ = range time.Tick(1e10) {
			fmt.Printf("cur conn num: %f\n", cur_conn_num)
		}
	}()

	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range conn_chan {
				ch_conn_change <- 1
				Conn_Func(conn)
				ch_conn_change <- -1
			}
		}()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			println("Error accept:", err.Error())
			return
		}
		conn_chan <- conn
	}
}

func Conn_Func(conn net.Conn) {

	remote_adr := conn.RemoteAddr()

	println("Conn_Func Start", remote_adr.Network(), remote_adr.String())

	defer conn.Close()
	buf := make([]byte, 1024)

	for {
		_, err := conn.Read(buf)

		println("Echo begin readed")
		fmt.Println(string(buf))

		if err != nil {
			println("Error reading:", err.Error())
			return
		}
		//send reply
		println("Echo begin write0")
		_, err = conn.Write(buf)
		println("Echo begin writeed")

		if err != nil {
			println("Error send reply:", err.Error())
			return
		}
	}
}
