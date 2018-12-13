// protocol
package unet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"ubus/ulog"
	"ubus/unet/uconnect"
	"ubus/uptl"
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

	CMD_SET_FAIL    = 0x11
	CMD_SET_SUCCESS = 0x55
)

const (
	ConstHeader         = "www.01happy.com"
	ConstHeaderLength   = 15
	ConstSaveDataLength = 4
)

type baseCmd struct {
	Cmd_type byte
	Cmd_cmd  byte
	Cmd_data []byte
}

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//解包
func Unpack(buffer []byte, conn net.Conn) []byte {

	fmt.Printf("Unpack buf=%x\n", buffer)
	ulog.Ul.Debugf("Unpack buf=%x", buffer)

	handle_num := 0
	var rms_buf []byte
	var ret bool

	for len(buffer[0:]) > 0 {

		switch {
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_B3:
			fmt.Println("Unpack dev cmd_b3")
			ulog.Ul.Debugln("Unpack dev cmd_b3")
			p_B3 := new(uptl.CmdB3)
			rms_buf = p_B3.Unpack(buffer[0:], conn)
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_B2:
			fmt.Println("Unpack dev cmd_b2")
			ulog.Ul.Debugln("Unpack dev cmd_b2")
			p_B2 := new(uptl.CmdB2)
			rms_buf = p_B2.Unpack(buffer[0:], conn)
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_B1:
			fmt.Println("Unpack dev cmd_b1")
			ulog.Ul.Debugln("Unpack dev cmd_b1")
			p_B1 := new(uptl.CmdB1)
			rms_buf = p_B1.Unpack(buffer[0:], conn)
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_B0:
			fmt.Println("Unpack dev cmd_b0")
			ulog.Ul.Debugln("Unpack dev cmd_b0")
			p_B0 := new(uptl.CmdB0)
			rms_buf = p_B0.Unpack(buffer[0:], conn)
		case buffer[0] == PC_CMD_TYPE && buffer[1] == CMD_B0:
			fmt.Println("Unpack PC cmd_b0")
			ulog.Ul.Debugln("Unpack PC cmd_b0")
			p_B0 := new(uptl.CmdB0_PC)
			rms_buf = p_B0.Unpack(buffer[0:], conn)
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A0:
			fmt.Println("Unpack dev cmd_a0")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A1:
			fmt.Println("Unpack dev cmd_a1")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A2:
			fmt.Println("Unpack dev cmd_a2")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A3:
			fmt.Println("Unpack dev cmd_a3")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A4:
			fmt.Println("Unpack dev cmd_a4")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A5:
			fmt.Println("Unpack dev cmd_a5")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A6:
			fmt.Println("Unpack dev cmd_a6")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A7:
			fmt.Println("Unpack dev cmd_a7")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A8:
			fmt.Println("Unpack dev cmd_a8")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_A9:
			fmt.Println("Unpack dev cmd_a9")
			fallthrough

		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_90:
			fmt.Println("Unpack dev cmd_90")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_91:
			fmt.Println("Unpack dev cmd_91")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_92:
			fmt.Println("Unpack dev cmd_92")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_93:
			fmt.Println("Unpack dev cmd_93")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_94:
			fmt.Println("Unpack dev cmd_94")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_95:
			fmt.Println("Unpack dev cmd_95")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_96:
			fmt.Println("Unpack dev cmd_96")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_97:
			fmt.Println("Unpack dev cmd_97")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_98:
			fmt.Println("Unpack dev cmd_98")
			fallthrough
		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_99:
			fmt.Println("Unpack dev cmd_99")
			ulog.Ul.Debugln("Unpack dev cmd_99")

			p := new(uptl.CmdRet)
			rms_buf = p.UnDevRetpack(buffer[0:])

		case buffer[0] == PC_CMD_TYPE && buffer[1] == CMD_A8:
			ulog.Ul.Debugln("Unpack pc cmd_a8")
			fmt.Println("Unpack pc cmd_a8")
			p := new(uptl.CmdA8)
			rms_buf, ret = p.UnPcPack(buffer, conn)
			if ret {
				//			p.CmdPackToDev(conn)
			}
			//			return buf[0:]

		case buffer[0] == DEV_CMD_TYPE && buffer[1] == CMD_AA:
			ulog.Ul.Debugln("Unpack dev cmd_aa")
			fmt.Println("Unpack dev cmd_aa")
			p := new(uptl.CmdAA_DEV_RET)
			rms_buf, _ = p.UnPcPack(buffer, conn)
			//			return buf[0:]

		case buffer[0] == PC_CMD_TYPE && buffer[1] == CMD_AA:
			ulog.Ul.Debugln("Unpack pc cmd_aa")
			fmt.Println("Unpack pc cmd_aa")
			p := new(uptl.CmdAA)
			rms_buf, _ = p.UnPcPack(buffer, conn)
			//			return buf[0:]

		case buffer[0] == MAN_CMD_TYPE && buffer[1] == CMD_AA:
			ulog.Ul.Debugln("Unpack man cmd_aa")
			fmt.Println("Unpack man cmd_aa")
			p := new(uptl.CmdAA_MAN_RET)
			rms_buf, _ = p.UnPcPack(buffer, conn)
			//			return buf[0:]

		case buffer[0] == FRONT_LED_CMD_TYPE && buffer[1] == CMD_A4:
			ulog.Ul.Debugln("Unpack front led cmd_a4")
			fmt.Println("Unpack front led cmd_a4")
			p := new(uptl.CmdA4_FRONT_LED_RET)
			rms_buf, _ = p.UnPcPack(buffer, conn)
			//			return buf[0:]
		case buffer[0] == MID_LED_CMD_TYPE && buffer[1] == CMD_A6:
			ulog.Ul.Debugln("Unpack mid led cmd_a6")
			fmt.Println("Unpack mid led cmd_a6")
			p := new(uptl.CmdA6_MID_LED_RET)
			rms_buf, _ = p.UnPcPack(buffer, conn)
			//			return buf[0:]

		case buffer[0] == INFO_LED_CMD_TYPE && buffer[1] == CMD_A2:
			ulog.Ul.Debugln("Unpack info led cmd_a2")
			fmt.Println("Unpack info led cmd_a2")
			p := new(uptl.CmdA2_INFO_LED_RET)
			rms_buf, _ = p.UnPcPack(buffer, conn)
			//			return buf[0:]

		default:
			fmt.Println("no 0x1BX")
			return buffer[len(buffer):]

		}

		buffer = nil

		if len(rms_buf) > 0 {
			fmt.Printf("rms_buf buf=%x\n", rms_buf)
			ulog.Ul.Debugf("rms_buf buf=%x", rms_buf)

			handle_num++
			buffer = rms_buf

			fmt.Printf("buffer buf=%x\n", buffer)
			ulog.Ul.Debugf("buffer buf=%x", buffer)

		}

		if handle_num > 5 {
			rms_buf = nil
			buffer = nil
			return nil
		}

	}
	//	length := len(buffer)
	//	var i int
	//	for i = 0; i < length; i = i + 1 {
	//		if length < i+ConstHeaderLength+ConstSaveDataLength {
	//			break
	//		}
	//		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
	//			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
	//			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {
	//				break
	//			}
	//			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
	//			readerChannel <- data

	//			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
	//		}
	//	}

	//	if i == length {
	//	return make([]byte, 0)
	//	}
	//		return buffer
	//fmt.Println("return buffer byte1")
	return buffer[0:]
}

func CodeClientReq() {

	//	fmt.Println("CodeClientReq GetClientReqNum=", usql.QueryClientReqProtocolCode())

	if usql.QueryClientReqProtocolCode() > 0 {

		fmt.Printf("%s::req protocol num=%d;\n", ulog.CreateDateSting(), usql.GetClientReqNum())
		//get first client req
		req := usql.GetClientReqFirst()
		fmt.Printf("%s::first req index=%d cmd=%x;\n", ulog.CreateDateSting(), req.PIndex, req.CMD)

		//get device_id from busnum
		device_id := usql.QueryDeviceIdFromBceDeviceInfoByBusNum(req.BusNum)

		//get socket conn from device_id
		conn_net := uconnect.GetDevConnObj([]byte(device_id)[0:])

		fmt.Println(" req device_id=", device_id)
		fmt.Println(" req conn_net=", conn_net)

		var write_buf []byte

		write_buf = uptl.CmdPack(req)

		uptl.SendCmdPack(write_buf, conn_net)

		//	fmt.Println("send req cmd pack=", write_buf)
		//	if conn_net != nil {
		//		w_n, _ := conn_net.Write(write_buf)
		//		fmt.Println("write buf w_n=%d", w_n)
		//	}
	} else {
		fmt.Printf("%s::req protocol num=%d;\n", ulog.CreateDateSting(), usql.GetClientReqNum())

	}

}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
