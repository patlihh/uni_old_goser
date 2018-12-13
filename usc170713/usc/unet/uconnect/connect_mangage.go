package uconnect

import (
	"bytes"
	"container/list"
	"fmt"
	"net"
	"strings"
	"usc/ulog"
	//	"ubus/usql"
)

const (
	online   = 0
	offline  = 1
	abnormal = 2
	timeout  = 3

	device_id_len = 10
)

type Connect_Obj struct {
	cmd_type  byte
	Uconn     net.Conn
	device_id [device_id_len]byte
}

type Pc_Connect_Obj struct {
	cmd_type byte
	Uconn    net.Conn
	port     int
}

var dev_conn_lists = list.New()

var pc_conn_lists = list.New()

func (co *Connect_Obj) SetConnObj(c net.Conn, device_id [device_id_len]byte) {
	co.Uconn = c
	fmt.Println(c.RemoteAddr().String())
	copy(co.device_id[:], device_id[0:])
}

func GetDevConnObj(device_id []byte) net.Conn {

	var c Connect_Obj
	var n *list.Element

	for e := dev_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Connect_Obj)

		//		ulog.Ul.Debug("GET list c nw1=", c.Uconn.RemoteAddr().String())
		//		ulog.Ul.Debug("GET conn nw1=", conn.RemoteAddr().String())

		if bytes.Equal(device_id[0:], c.device_id[0:]) {
			ulog.Ul.Debug("get uconnobj success")
			return c.Uconn
		}

	}

	ulog.Ul.Debug("get dev connobj failed")

	return nil
}

func GetPcConnObj(port int) net.Conn {

	var c Pc_Connect_Obj
	var n *list.Element

	for e := pc_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Pc_Connect_Obj)

		//		fmt.Println("GET list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("GET conn nw1=", conn.RemoteAddr().String())

		if port == c.port {
			ulog.Ul.Debug("get uconnobj success")
			return c.Uconn
		}

	}

	ulog.Ul.Debug("get app connobj failed")

	return nil
}

//func SetConnObjMacAdr(conn net.Conn, mac_adr []byte) bool{
//	var c Connect_Obj

//	for e := Uconn_lists.Front(); e != nil; e.Next() {
//		c = e.Value.(Connect_Obj)
//		if strings.EqualFold(c.Uconn.RemoteAddr().String(), conn.Uconn.RemoteAddr().String()) {

//			fmt.Println("conn is exist")
//			return true
//		}

//	}
//}

func AddDevConnObjToList(conn Connect_Obj) bool {
	var c Connect_Obj
	var n *list.Element

	for e := dev_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Connect_Obj)

		//		fmt.Println("list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("conn nw1=", conn.Uconn.RemoteAddr().String())

		if strings.EqualFold(c.Uconn.RemoteAddr().String(), conn.Uconn.RemoteAddr().String()) {
			ulog.Ul.Debug("conn is exist in list")
			return false
		}

	}

	dev_conn_lists.PushBack(conn)

	ulog.Ul.Debug(ulog.CreateDateWithNanoSting(), "::", "dev conn added list ok. num=", dev_conn_lists.Len())

	return true
}

func AddPcConnObjToList(conn Pc_Connect_Obj) bool {
	var c Pc_Connect_Obj
	var n *list.Element

	for e := pc_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Pc_Connect_Obj)

		//		fmt.Println("list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("conn nw1=", conn.Uconn.RemoteAddr().String())

		if c.port == conn.port {
			ulog.Ul.Debug("pc conn is exist in list")
			return false
		}

	}

	pc_conn_lists.PushBack(conn)

	ulog.Ul.Debugf(ulog.CreateDateWithNanoSting(), "::", "pc conn added list ok. num=", pc_conn_lists.Len())

	return true
}

func AddDevConnAndMacAdrToList(cmd_type byte, conn net.Conn, device_id [device_id_len]byte) bool {

	c := Connect_Obj{cmd_type, conn, device_id}

	return AddDevConnObjToList(c)

	//	return true

}

func AddPcConnToList(cmd_type byte, conn net.Conn, port int) bool {

	ulog.Ul.Debugf("AddPcConnToList port==%d", port)

	c := Pc_Connect_Obj{cmd_type, conn, port}

	return AddPcConnObjToList(c)

	//	return true

}

func GetDevConnObjMarAdrFromList(conn net.Conn) []byte {

	var c Connect_Obj
	var n *list.Element

	for e := dev_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Connect_Obj)

		//		fmt.Println("GET list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("GET conn nw1=", conn.RemoteAddr().String())

		if strings.EqualFold(c.Uconn.RemoteAddr().String(), conn.RemoteAddr().String()) {
			ulog.Ul.Debug("get uconnobj success")
			return c.device_id[0:device_id_len]
		}

	}

	ulog.Ul.Debug("get dev connobj failed")
	return make([]byte, 0)

}

func GetPcConnPortFromList(conn net.Conn) int {

	var c Pc_Connect_Obj
	var n *list.Element

	for e := pc_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Pc_Connect_Obj)

		//		fmt.Println("GET list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("GET conn nw1=", conn.RemoteAddr().String())

		if strings.EqualFold(c.Uconn.RemoteAddr().String(), conn.RemoteAddr().String()) {
			ulog.Ul.Debug("get uconnobj success")
			return c.port
		}

	}

	ulog.Ul.Debugf("get pc connobj failed")
	return 0

}

func CountDevConnList() int {
	ulog.Ul.Debug("conn dev_conn_lists count=", dev_conn_lists.Len())
	return dev_conn_lists.Len()
}

func CountPcConnList() int {
	ulog.Ul.Debug("conn pc_conn_lists count=", pc_conn_lists.Len())
	return pc_conn_lists.Len()
}

func RemoveDevConnFromList(conn net.Conn) bool {

	var c Connect_Obj
	var n *list.Element

	for e := dev_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Connect_Obj)

		//		fmt.Println("remove list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("remove conn nw1=", conn.RemoteAddr().String())

		if strings.EqualFold(c.Uconn.RemoteAddr().String(), conn.RemoteAddr().String()) {
			dev_conn_lists.Remove(e)

			//dev net disconnect state write database
			//			pns := new(usql.BceDeviceStatus)
			//			pns.DeviceId = fmt.Sprintf("%s", c.device_id[0:])
			//			pns.Time = ulog.CreateDateSting()
			//			pns.DeviceNetStatus = '0'
			//			pns.InsertBceDeviceStatus()

			ulog.Ul.Debug(ulog.CreateDateWithNanoSting(), "::", "dev conn device_id=", string(c.device_id[:]), "removed ok. num=", dev_conn_lists.Len())
			return true
		}

	}
	ulog.Ul.Debug(ulog.CreateDateWithNanoSting(), "::", "dev conn not exist and removed fail")
	return false
}

func RemovePcConnFromList(conn net.Conn) bool {

	var c Pc_Connect_Obj
	var n *list.Element

	for e := pc_conn_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(Pc_Connect_Obj)

		//		fmt.Println("remove list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("remove conn nw1=", conn.RemoteAddr().String())

		if strings.EqualFold(c.Uconn.RemoteAddr().String(), conn.RemoteAddr().String()) {
			pc_conn_lists.Remove(e)
			ulog.Ul.Debug(ulog.CreateDateWithNanoSting(), "::", "pc conn port=", c.port, "removed ok. num=", pc_conn_lists.Len())
			return true
		}

	}
	ulog.Ul.Debug(ulog.CreateDateWithNanoSting(), "::", "pc conn not exist and removed fail")
	return false
}

func RemoveUconnFromList(conn net.Conn) {
	RemoveDevConnFromList(conn)
	RemovePcConnFromList(conn)
}
