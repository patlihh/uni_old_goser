// CmdBase
package uptl

import (
	"fmt"
	"net"
	"ubus/ulog"
	"ubus/unet/uconnect"
	"ubus/usql"
)

const (
	SERVER_TYPE        = 0x01
	DEV_CMD_TYPE       = 0x10
	PC_CMD_TYPE        = 0x20
	APP_CMD_TYPE       = 0x30
	MAN_CMD_TYPE       = 0X40
	FRONT_LED_CMD_TYPE = 0X50
	TAIL_LED_CMD_TYPE  = 0X51
	INFO_LED_CMD_TYPE  = 0X52
	MID_LED_CMD_TYPE   = 0X53

	CMD_A0 = 0xA0
	CMD_A1 = 0xA1
	CMD_A2 = 0xA2
	CMD_A3 = 0xA3
	CMD_A4 = 0xA4
	CMD_A5 = 0xA5
	CMD_A6 = 0xA6
	CMD_A7 = 0xA7
	CMD_A8 = 0xA8
	CMD_A9 = 0xA9
	CMD_AA = 0xAA
	CMD_90 = 0x90
	CMD_91 = 0x91
	CMD_92 = 0x92
	CMD_93 = 0x93
	CMD_94 = 0x94
	CMD_95 = 0x95
	CMD_96 = 0x96
	CMD_97 = 0x97
	CMD_98 = 0x98
	CMD_99 = 0x99
	CMD_C4 = 0xC4
	CMD_C5 = 0xC5
	CMD_C6 = 0xC6

	CMD_B0 = 0xB0
	CMD_B1 = 0xB1
	CMD_B2 = 0xB2
	CMD_B3 = 0xB3
	CMD_B4 = 0xB4
	CMD_B5 = 0xB5
	CMD_B6 = 0xB6

	CMD_SET_FAIL    = 0x10
	CMD_SET_SUCCESS = 0x50

	NO_FIRM_ID    = 0XE0
	NO_LINE_ID    = 0XE1
	NO_STATION_ID = 0XE2
	NO_BCE_DEVICE = 0XE3
	NO_BUS_ID     = 0XE4
	NO_DRIVER_ID  = 0XE5

	BUS_BEGIN_START  = 0X30
	BUS_END_START    = 0X31
	BUS_BEGIN_ARRIVE = 0X32
	BUS_END_ARRIVE   = 0X33
	BUS_FAULT        = 0X34
	BUS_RUNING       = 0X35
	BUS_SCHEDULING   = 0X36
)

type CmdRet struct {
	cmd_type   byte
	cmd_cmd    byte
	pindex     [4]byte
	ret_status byte
}

func CmdPack(req usql.BceProtocolCode) []byte {
	//	return append(append(IntToBytes(req.Index), []byte{0x01, req.CMD}...), []byte(req.Data[0:])...)
	return append(append([]byte{0x01, req.CMD}, IntToBytes(req.PIndex)...), []byte(req.Data[0:])...)
}

func GetConnObjByBusNum(bus_num string) net.Conn {
	//get device_id from busnum
	device_id := usql.QueryDeviceIdFromBceDeviceInfoByBusNum(string(bus_num[0:]))

	//get socket conn from device_id
	return uconnect.GetDevConnObj([]byte(device_id)[0:])

}

func GetConnObjByBceId(bce_id string) net.Conn {
	//get device_id from busnum
	//	device_id := usql.QueryDeviceIdFromBceDeviceInfoByBusNum(string(bus_num[0:]))

	//get socket conn from device_id
	return uconnect.GetDevConnObj([]byte(bce_id)[0:])

}

func SendCmdPack(buffer []byte, conn net.Conn) {

	fmt.Printf("SendCmdAxPack cmd =%x\n", buffer[1])
	fmt.Println("SendCmdAxPack buffer =", buffer)
	//confirm pack valid
	var is_valid bool
	is_valid = false
	switch buffer[1] {
	case 0xA0:
		is_valid = IsValidCmdA0Pack(buffer)
	case 0xA1:
		is_valid = IsValidCmdA1Pack(buffer)
	case 0xA2:
		is_valid = IsValidCmdA2Pack(buffer)
	case 0xA3:
		is_valid = IsValidCmdA3Pack(buffer)
	case 0xA4:
		is_valid = IsValidCmdA4Pack(buffer)
	case 0xA5:
		is_valid = IsValidCmdA5Pack(buffer)
	case 0xA6:
		is_valid = IsValidCmdA6Pack(buffer)
	case 0xA7:
		is_valid = IsValidCmdA7Pack(buffer)
	case 0xA8:
		//		buffer = append(buffer, []byte{0x03, '3', '5', '2'}...)
		is_valid = IsValidCmdA8Pack(buffer)
	case 0xA9:
		is_valid = IsValidCmdA9Pack(buffer)
	case 0x90:
		is_valid = IsValidCmdU90Pack(buffer)
	case 0x91:
		is_valid = IsValidCmdU91Pack(buffer)
	case 0x92:
		is_valid = IsValidCmdU92Pack(buffer)
	case 0x93:
		is_valid = IsValidCmdU93Pack(buffer)
	case 0x94:
		is_valid = IsValidCmdU94Pack(buffer)
	case 0x95:
		is_valid = IsValidCmdU95Pack(buffer)
	case 0x96:
		is_valid = IsValidCmdU96Pack(buffer)
	case 0x97:
		is_valid = IsValidCmdU97Pack(buffer)
	case 0x98:
		is_valid = IsValidCmdU98Pack(buffer)
	case 0x99:
		is_valid = IsValidCmdU99Pack(buffer)

	default:
		fmt.Printf("SendCmdAxPack invalid cmd =%x\n", buffer[1])

	}

	//Remove invalid req from list and set fail in mysql database
	if is_valid == false {
		ppc := new(usql.BceProtocolCode)

		ppc.PIndex = Bytes4ToInt(buffer[2:6])
		ppc.RspDateTime = ulog.CreateDateSting()
		ppc.WorkStatus = '1'
		ppc.RspCode = '1'
		//		ppc.Remarks = ""

		ppc.UpdateBceProtocolCodeRsp()
		usql.RemoveClientReqByPIndex(Bytes4ToInt(buffer[2:6]))
	}

	//pack valid and conn valid then send pack
	if is_valid == true && conn != nil {

		//		fmt.Printf("CMDA8 SendPack buffer %v \n", buffer)
		w_n, err := conn.Write(buffer)
		//		fmt.Printf("CMDA8 SendPack len = %d; err = %s \n", w_n, err)

		if err != nil {
			fmt.Printf("CMDAx SendPack len = %d; err = %s \n", w_n, err)
		}

	}

}

func (a_r *CmdRet) UnDevRetpack(buffer []byte) []byte {

	buf_length := len(buffer)

	fmt.Println("buf_len=", buf_length, "Unpack buf=", buffer)

	if buf_length >= 7 {
		a_r.cmd_type = buffer[0]
		a_r.cmd_cmd = buffer[1]
		a_r.pindex[0] = buffer[2]
		a_r.pindex[1] = buffer[3]
		a_r.pindex[2] = buffer[4]
		a_r.pindex[3] = buffer[5]
		a_r.ret_status = buffer[6]

		fmt.Printf("cmd ret_status=%x\n", a_r.ret_status)

		port := Bytes4ToInt(a_r.pindex[0:])

		fmt.Printf("UnDevRetpack port==%d\n", port)

		if a_r.ret_status == 0x55 {

			conn := uconnect.GetPcConnObj(port)

			if conn != nil {

				w_len, _ := conn.Write(buffer[0:7])
				fmt.Printf("pc conn write 0x55 len=%d!\n", w_len)
			}

			//			ppc := new(usql.BceProtocolCode)

			//			ppc.PIndex = Bytes4ToInt(a_r.pindex[0:])
			//			ppc.RspDateTime = ulog.CreateDateSting()
			//			ppc.WorkStatus = '1'
			//			ppc.RspCode = '0'
			//			//			ppc.Remarks = ""

			//			ppc.UpdateBceProtocolCodeRsp()

			//			usql.RemoveClientReqByPIndex(ppc.PIndex)

		}
		if a_r.ret_status == 0x11 {

			//			ppc := new(usql.BceProtocolCode)

			//			ppc.PIndex = Bytes4ToInt(a_r.pindex[0:])
			//			ppc.RspDateTime = ulog.CreateDateSting()
			//			ppc.WorkStatus = '1'
			//			ppc.RspCode = '1'
			//			//			ppc.Remarks = ""

			//			ppc.UpdateBceProtocolCodeRsp()
			//			usql.RemoveClientReqByPIndex(ppc.PIndex)
			conn := uconnect.GetPcConnObj(port)

			if conn != nil {

				w_len, _ := conn.Write(buffer[0:7])
				fmt.Printf("pc conn write 0x11 len=%d!\n", w_len)
			}

		}

		return buffer[7:]
	}

	return buffer[0:]
}

func (a_r *CmdRet) Unpack(buffer []byte, conn net.Conn) []byte {

	buf_length := len(buffer)

	fmt.Println("buf_len=", buf_length, "Unpack buf=", buffer)

	if buf_length >= 7 {
		a_r.cmd_type = buffer[0]
		a_r.cmd_cmd = buffer[1]
		a_r.pindex[0] = buffer[2]
		a_r.pindex[1] = buffer[3]
		a_r.pindex[2] = buffer[4]
		a_r.pindex[3] = buffer[5]
		a_r.ret_status = buffer[6]

		fmt.Printf("cmd ret_status=%x\n", a_r.ret_status)
		if a_r.ret_status == 0x55 {

			ppc := new(usql.BceProtocolCode)

			ppc.PIndex = Bytes4ToInt(a_r.pindex[0:])
			ppc.RspDateTime = ulog.CreateDateSting()
			ppc.WorkStatus = '1'
			ppc.RspCode = '0'
			//			ppc.Remarks = ""

			ppc.UpdateBceProtocolCodeRsp()

			usql.RemoveClientReqByPIndex(ppc.PIndex)

		}
		if a_r.ret_status == 0x11 {

			ppc := new(usql.BceProtocolCode)

			ppc.PIndex = Bytes4ToInt(a_r.pindex[0:])
			ppc.RspDateTime = ulog.CreateDateSting()
			ppc.WorkStatus = '1'
			ppc.RspCode = '1'
			//			ppc.Remarks = ""

			ppc.UpdateBceProtocolCodeRsp()
			usql.RemoveClientReqByPIndex(ppc.PIndex)

		}

		return buffer[7:]
	}

	return buffer[0:]
}
