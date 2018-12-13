// A8SwitchLine
package uptl

import (
	"fmt"
	"net"
	"strconv"
	//	"ubus/unet/uconnect"
)

type CmdA8 struct {
	cmd_type    byte
	cmd         byte
	data_len    byte
	bus_num_len byte
	bus_num     string
	direction   byte
	line_len    byte
	line_id     string
}

func (a8 *CmdA8) CmdPackToDev(pc_conn net.Conn) (int, error) {
	//	return append(append(IntToBytes(req.Index), []byte{0x01, req.CMD}...), []byte(req.Data[0:])...)
	dev_conn_net := GetConnObjByBusNum(a8.bus_num)

	_, port_str, _ := net.SplitHostPort(pc_conn.RemoteAddr().String())

	port, _ := strconv.Atoi(port_str)

	fmt.Printf("CmdPackToDev port==%d\n", port)

	write_buf := append(append(append([]byte{0x01, a8.cmd}, IntToBytes(port)...), []byte{a8.direction, a8.line_len}...), []byte(a8.line_id[0:])...)

	return dev_conn_net.Write(write_buf)
}

//func (a8 *CmdA8) CmdPackToDev() []byte {
//	//	return append(append(IntToBytes(req.Index), []byte{0x01, req.CMD}...), []byte(req.Data[0:])...)
//	return append(append(append([]byte{0x01, a8.cmd}, IntToBytes(0)...), []byte{a8.direction, a8.line_len}...), []byte(a8.line_id[0:])...)
//}

//解包
func (a8 *CmdA8) UnPcPack(buffer []byte, conn net.Conn) ([]byte, bool) {

	d_len := len(buffer)

	if d_len > 8 {
		a8.cmd_type = buffer[0]
		a8.cmd = buffer[1]
		a8.data_len = buffer[2]

		fmt.Printf("a8 Unpack a8.data_len=%d\n", a8.data_len)

		if d_len >= int(a8.data_len+3) {
			a8.bus_num_len = buffer[3]
			fmt.Printf("a8 Unpack a8.bus_num_len=%d\n", a8.bus_num_len)

			//			if (int)(4+a8.bus_num_len) <= d_len {
			//				a8.bus_num = string(buffer[4 : 4+a8.bus_num_len])
			//			}
			//			a8.direction = buffer[4+a8.bus_num_len]
			//			a8.line_len = buffer[5+a8.bus_num_len]
			//			if (int)(6+a8.bus_num_len+a8.line_len) <= d_len {
			//				a8.line_id = string(buffer[6+a8.bus_num_len : 6+a8.bus_num_len+a8.line_len])
			//			}

			//			//			fmt.Printf("a8 Unpack data len=%d;line_len=%d\n", d_len, a8.line_len)
			//			fmt.Printf("a8 Unpack data len=%d;bus_num=%s;line_id=%s\n", d_len, a8.bus_num, a8.line_id)

			//			_, port_str, _ := net.SplitHostPort(conn.RemoteAddr().String())

			//			port, _ := strconv.Atoi(port_str)
			//			uconnect.AddPcConnToList(buffer[0], conn, port)

			return buffer[a8.data_len+3:], true
		}

		return buffer[0:], false

	}

	return buffer[0:], false
}

func IsValidCmdA8Pack(pack []byte) bool {

	b_len := len(pack)
	//pack data len too small
	if b_len < cmd_index_len+4 {
		fmt.Printf("IsValidCmdA8Pack data len=%d too small\n", b_len)
		return false
	}

	//type error
	if pack[cmd_type_pos] != 0x01 {
		fmt.Printf("IsValidCmdA8Pack type error\n")
		return false
	}

	//cmd error
	if pack[cmd_pos] != 0xa8 {
		fmt.Printf("IsValidCmdA8Pack cmd error\n")
		return false
	}

	//direction error

	if pack[cmd_pos+cmd_index_len+1] != '0' && pack[cmd_pos+cmd_index_len+1] != '1' {
		fmt.Printf("IsValidCmdA8Pack direction error\n")
		return false
	}

	line_id_len, _ := strconv.Atoi(fmt.Sprintf("%d", pack[cmd_pos+cmd_index_len+2]))
	//protocol data len err
	if b_len != cmd_index_len+4+line_id_len {
		fmt.Printf("IsValidCmdA8Pack data len error\n")
		return false
	}

	return true
}
