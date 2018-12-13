// utcp_manage
package unet

import (
	"fmt"
	"net"
	"os"
	//	"strconv"
	"time"
	"usms/ulog"
	"usms/unet/uconnect"
)

const (
	MAX_CONN_NUM = 2000
)

//var cur_conn_num int = 0
//var conn_chan = make(chan net.Conn)
var ch_conn_change = make(chan int)

//var Conns_chan = make(chan net.Conn, MAX_CONN_NUM)

func TcpManage() {

	ulog.Ul.Debug("tcp_manage start....listen 9966")

	listener, err := net.Listen("tcp", ":9966")
	//listener, err := net.Listen("tcp", ":5234")

	if err != nil {
		ulog.Ul.Debugln("error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	var cur_conn_num int = 0
	conn_chan := make(chan net.Conn)
	//	ch_conn_change := make(chan int)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
		}
	}()

	go func() {
		for _ = range time.Tick(6 * 1e10) {
			ulog.Ul.Debugf("%s:cur con num=%d;dev con in lis=%d;pc con in lis=%d;\n", ulog.CreateDateSting(), cur_conn_num, uconnect.CountDevConnList(), uconnect.CountPcConnList())
			fmt.Printf("%s:cur con num=%d;dev con in lis=%d;pc con in lis=%d;\n", ulog.CreateDateSting(), cur_conn_num, uconnect.CountDevConnList(), uconnect.CountPcConnList())
		}
	}()

	//	go func() {
	//		for _ = range time.Tick(1e9 + 10000) {
	//			CodeClientReq()
	//			//			usql.QueryClientReqProtocolCode()
	//			//			fmt.Printf("%s::req protocol num=%d;\n", ulog.CreateDateSting(), usql.GetClientReqNum())
	//		}
	//	}()

	for i := 0; i < MAX_CONN_NUM; i++ {
		go func() {
			for conn := range conn_chan {
				go Conn_Func(conn)
				//				ch_conn_change <- -1
			}
		}()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			ulog.Ul.Debug("Error accept:", err.Error())
			return
		}
		conn_chan <- conn
	}
}

func Conn_Func(conn net.Conn) {
	ch_conn_change <- 1

	remote_adr := conn.RemoteAddr()

	//	w_f := CreateDataFlie()
	//	defer CloseFile(w_f)

	println(ulog.CreateDateSting(), "::", "Conn_Func Start", remote_adr.Network(), remote_adr.String())

	//	_, port_str, _ := net.SplitHostPort(remote_adr.String())

	//	port, _ := strconv.Atoi(port_str)

	//	fmt.Printf("conn.LocalAddr port=%d\n", port)

	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//	//	//声明一个管道用于接收解包的数据
	//	readerChannel := make(chan []byte, 16)

	//	go porotol_data_proc(conn, readerChannel)

	buffer := make([]byte, 5*1024)

	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(1 * time.Minute))

		len, err := conn.Read(buffer)
		if err != nil {
			ulog.Ul.Debugln(ulog.CreateDateSting(), "::", "Error reading:", err.Error())
			uconnect.RemoveUconnFromList(conn)
			ch_conn_change <- -1
			return
		}

		//		fmt.Printf("sockt readed nw=%s;len=%d; connlistnum=%d\n", remote_adr.String(), len, uconnect.CountUconnList())

		//		if buffer[0] == 0x01 && (buffer[1] == 0xB1 || buffer[1] == 0xB2 || buffer[1] == 0xB3 || buffer[1] == 0xB5) {
		//			//			WriteStringToFile(fmt.Sprintf("len=%d;% X\n", len, buffer[0:len]), w_f)
		//			fmt.Printf("sockt readed nw=%s;len=%d;buf=%X\n", remote_adr.String(), len, buffer[0:len])
		//		} else {
		//			//			WriteStringToFile(fmt.Sprintf("len=%d\n", len), w_f)
		//			fmt.Printf("soctk readed len = %v;\n", len)
		//		}

		tmpBuffer = Unpack(append(tmpBuffer, buffer[:len]...), conn)
	}

	uconnect.RemoveUconnFromList(conn)
	ch_conn_change <- -1

	ulog.Ul.Debug("Conn_Func end")
}

//func porotol_data_proc(conn net.Conn, readerChannel chan []byte) {

//	for {

//		select {

//		case data := <-readerChannel:

//			fmt.Printf("data len = %d;\n", len(data))

//			//			Log(string(data))
//		}
//	}

//}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
