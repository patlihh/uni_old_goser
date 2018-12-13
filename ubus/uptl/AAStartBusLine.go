// AAStartBusLine
package uptl

import (
	"fmt"
	"net"
	"strconv"
	"ubus/umisc"
	//	"ubus/unet/uconnect"
	"ubus/ulog"
	"ubus/usql"
)

type CmdAA struct {
	cmd_type         byte
	cmd              byte
	data_len         byte
	schedul_id       int64
	bce_id_len       byte
	bce_id           string
	direction        byte
	line_len         byte
	line_id          string
	schedul_time_len byte
}
type CmdAA_DEV_RET struct {
	cmd_type   byte
	cmd        byte
	schedul_id int64
	ret        byte
}
type CmdAA_MAN_RET struct {
	cmd_type   byte
	cmd        byte
	schedul_id int64
	ret        byte
}

//func (a *CmdAA) CmdPackToDev(buffer []byte, pc_conn net.Conn) (int, error) {
//	//	return append(append(IntToBytes(req.Index), []byte{0x01, req.CMD}...), []byte(req.Data[0:])...)
//	dev_conn_net := GetConnObjByBceId(a.bce_id)

//	_, port_str, _ := net.SplitHostPort(pc_conn.RemoteAddr().String())

//	port, _ := strconv.Atoi(port_str)

//	fmt.Printf("CmdAAPackToDev port==%d\n", port)

//	write_buf := append(append(append([]byte{SERVER_TYPE, a.cmd}, IntToBytes(port)...), []byte{a.direction, a.line_len}...), []byte(a.line_id[0:])...)

//	return dev_conn_net.Write(write_buf)
//}

//func (a8 *CmdA8) CmdPackToDev() []byte {
//	//	return append(append(IntToBytes(req.Index), []byte{0x01, req.CMD}...), []byte(req.Data[0:])...)
//	return append(append(append([]byte{0x01, a8.cmd}, IntToBytes(0)...), []byte{a8.direction, a8.line_len}...), []byte(a8.line_id[0:])...)
//}

//解包
func (a *CmdAA_MAN_RET) UnPcPack(buffer []byte, conn net.Conn) ([]byte, bool) {

	d_len := len(buffer)

	ulog.Ul.Debugf("CmdAA_MAN_RET UnPcPack buf len=%d", d_len)

	if d_len >= 11 {
		a.cmd_type = buffer[0]
		a.cmd = buffer[1]
		a.schedul_id = umisc.Bytes8ToIntOfBigEndian(buffer[2:10])
		a.ret = buffer[10]

		ulog.Ul.Debugf("CmdAA_DEV_RET Unpack a.schedul_id=%d;a.ret=%d\n", a.schedul_id, a.ret)

		/******write db******/
		if usql.QueryDispatchIDFromDispatchInfo(a.schedul_id) {

			pdb := new(usql.DispatchResult)

			pdb.DispatchId = a.schedul_id
			pdb.DispatchResultTime = umisc.CreateDateSting()
			pdb.DeviceType = fmt.Sprintf("%d", MAN_CMD_TYPE)
			pdb.DispatchResult = fmt.Sprintf("%d", a.ret)

			pdb.InsertDispatchResult()

		}

		return buffer[11:], true

	}

	return buffer[0:], false
}

//解包
func (a *CmdAA_DEV_RET) UnPcPack(buffer []byte, conn net.Conn) ([]byte, bool) {

	d_len := len(buffer)

	if d_len >= 11 {
		a.cmd_type = buffer[0]
		a.cmd = buffer[1]
		a.schedul_id = umisc.Bytes8ToIntOfBigEndian(buffer[2:10])
		a.ret = buffer[10]

		ulog.Ul.Debugf("CmdAA_DEV_RET Unpack a.schedul_id=%d\n", a.schedul_id)

		/******write db******/
		if usql.QueryDispatchIDFromDispatchInfo(a.schedul_id) {

			pdb := new(usql.DispatchResult)

			pdb.DispatchId = a.schedul_id
			pdb.DispatchResultTime = umisc.CreateDateSting()
			pdb.DeviceType = fmt.Sprintf("%d", DEV_CMD_TYPE)

			//			if a.ret == CMD_SET_SUCCESS {
			//				pdb.DispatchResult = "0"
			//			} else {
			//				pdb.DispatchResult = "1"
			//			}

			pdb.DispatchResult = fmt.Sprintf("%d", a.ret)

			pdb.InsertDispatchResult()
		}

		return buffer[11:], true

	}

	return buffer[0:], false
}

//解包
func (a *CmdAA) UnPcPack(buffer []byte, conn net.Conn) ([]byte, bool) {

	d_len := len(buffer)

	ret := CMD_SET_SUCCESS

	if d_len > 8 {
		a.cmd_type = buffer[0]
		a.cmd = buffer[1]
		a.data_len = buffer[2]

		ulog.Ul.Debugf("CmdAA Unpack a.data_len=%d", a.data_len)
		pos := 3

		if d_len >= int(a.data_len)+pos {

			a.schedul_id = umisc.Bytes8ToIntOfLitEndian(buffer[pos : pos+8])
			pos += 8
			a.bce_id_len = buffer[pos]
			ulog.Ul.Debugf("aa Unpack a.bce_id_len=%d", a.bce_id_len)

			pos++
			a.bce_id = string(buffer[pos : pos+int(a.bce_id_len)])

			pos += int(a.bce_id_len)
			a.direction = buffer[pos]
			pos++
			a.line_len = buffer[pos]
			pos++
			if pos+int(a.line_len) <= d_len {
				a.line_id = string(buffer[pos : pos+int(a.line_len)])
			}
			pos += int(a.line_len)
			a.schedul_time_len = buffer[pos]

			//			fmt.Printf("a8 Unpack data len=%d;line_len=%d\n", d_len, a8.line_len)
			ulog.Ul.Debugf("Unpack data len=%d;schid=%d;bce_id=%s;line_id=%s;time_len=%d", d_len, a.schedul_id, a.bce_id, a.line_id, a.schedul_time_len)

			//			_, port_str, _ := net.SplitHostPort(conn.RemoteAddr().String())

			//			port, _ := strconv.Atoi(port_str)

			//			uconnect.AddPcConnToList(buffer[0], conn, port)

			/*****write protocol data to dev*********/
			dev_conn_net := GetConnObjByBceId(a.bce_id)
			if dev_conn_net != nil {
				w_dev_len, _ := dev_conn_net.Write(buffer[0 : a.data_len+3])
				ulog.Ul.Debugf("Unpack w_dev_len=%d;w_dev_buf=%x", w_dev_len, buffer[0:a.data_len+3])

				/*****write success result to pc************/
				if conn != nil {
					w_pc_buf := append([]byte{SERVER_TYPE, a.cmd}, []byte{CMD_SET_SUCCESS}...)

					w_pc_len, d_err := conn.Write(w_pc_buf)

					//update bus_device line_id
					if d_err != nil {
						pb := new(usql.BusDevice)
						pb.BusId = usql.QueryBusIdFromBusDeviceByDeviceId(a.bce_id)
						if len(pb.BusId) > 0 {
							pb.LineId = a.line_id
							pb.UpdateBusDeviceLineId()
						}
					}

					ulog.Ul.Debugf("Unpack w_pc_len=%d;w_pc_buf=%x", w_pc_len, w_pc_buf)
				}

			} else {
				/*****write no dev result to pc************/
				if conn != nil {
					w_pc_buf := append([]byte{SERVER_TYPE, a.cmd}, []byte{NO_BCE_DEVICE}...)
					ret = NO_BCE_DEVICE
					w_pc_len, _ := conn.Write(w_pc_buf)

					ulog.Ul.Debugf("Unpack w_pc_len=%d;w_pc_buf=%x", w_pc_len, w_pc_buf)
				}
			}
			/******write db******/

			if usql.QueryDispatchIDFromDispatchInfo(a.schedul_id) {
				pdb := new(usql.DispatchResult)

				pdb.DispatchId = a.schedul_id
				pdb.DispatchResultTime = umisc.CreateDateSting()
				pdb.DeviceType = fmt.Sprintf("%d", SERVER_TYPE)
				pdb.DispatchResult = fmt.Sprintf("%d", ret)

				pdb.InsertDispatchResult()
			}

			return buffer[a.data_len+3:], true
		}

		return buffer[0:], false

	}

	return buffer[0:], false
}

func IsValidCmdAAPack(pack []byte) bool {

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
