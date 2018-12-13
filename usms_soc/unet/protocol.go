// protocol
package unet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"usms_soc/ulog"
	"usms_soc/umisc"
	//	"usms_soc/uptl"
	"strings"
	"usms_soc/usql"
)

const (
	SERVER_TYPE  = 0x20
	DEV_CMD_TYPE = 0x10
	PC_CMD_TYPE  = 0x30
	APP_CMD_TYPE = 0x40

	CMD_10 = 0x10

	CMD_SET_FAIL    = 0x02
	CMD_SET_SUCCESS = 0x01
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

type Cmd10_UP_INFO_Packet struct {
	imei         string
	imsi         string
	ptype        string
	wtype        string
	hw_info      string
	soft_version string
}

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//解包
func Unpack(buffer []byte, conn net.Conn) []byte {

	fmt.Printf("Unpack buf=%x\n", buffer)
	fmt.Printf("Unpack buf str=%s\n", string(buffer))
	ulog.Ul.Debugf("Unpack buf=%x", buffer)

	sms_content := string(buffer)
	d_len := len(sms_content)
	fmt.Printf("d_len=%d\n", d_len)

	if d_len > len("(UNITONE) ") && strings.Contains(sms_content, "(UNITONE)") == true {

		s_new := strings.Trim(sms_content, "(UNITONE)")
		fmt.Println(s_new, len(s_new))

		s := strings.Split(s_new, ";")

		fmt.Println(d_len, s, len(s))

		a := new(Cmd10_UP_INFO_Packet)

		for _, str := range s {

			b_str := []byte(str)
			if b_str[0] == 'i' {
				a.imei = string(b_str[1:])
				fmt.Println(a.imei)
			} else if b_str[0] == 't' {
				a.hw_info = string(b_str[1:])
				fmt.Println(a.hw_info)
			} else if b_str[0] == 'p' {
				a.ptype = string(b_str[1:])
				fmt.Println(a.ptype)
			} else if b_str[0] == 's' {
				a.imsi = string(b_str[1:])
				fmt.Println(a.imsi)
			} else if b_str[0] == 'v' {
				a.soft_version = string(b_str[1:])
				fmt.Println(a.soft_version)
			} else {
				fmt.Println("error sms content!")
				conn.Write([]byte("2"))
				return nil //false
			}

		}

		if IsValidImei(a.imei) {

			pdb := new(usql.Cmd10_UP_INFO)

			pdb.Imei = a.imei
			pdb.Imsi = a.imsi
			pdb.Datatime = umisc.CreateDateSting()
			pdb.Device_type = '0'
			pdb.Phone_type = a.ptype
			pdb.Phone_num = "0"
			pdb.Softwareinfo = a.soft_version
			pdb.Hardwartinfo = a.hw_info
			ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
			pdb.Ip_addr = ip
			pdb.Locationinfo = "no"

			pdb.InsertUssDeviceInfo()
			conn.Write([]byte("1"))
			return nil //true
		}
	}
	//	else {
	//		if IsValidPhoneNum(number) {
	//			pdb2 := new(usql.Invaild_sms_record)
	//			pdb2.Datatime = umisc.CreateDateSting()
	//			pdb2.Phone_num = number
	//			pdb2.SmsContent = sms_content
	//			pdb2.InsertUssInvaildSmsRecord()
	//		}
	//	}
	conn.Write([]byte("3"))
	return nil //false
	//	handle_num := 0
	//	//	return nil

	//	var rms_buf []byte

	//	for len(buffer[0:]) > 0 {

	//		switch {
	//		case buffer[0] == DEV_CMD_TYPE && buffer[3] == CMD_10:
	//			fmt.Println("Unpack dev CMD_10")
	//			//			ulog.Ul.Debugln("Unpack dev cmd_b1")
	//			//			p_10 := new(uptl.Cmd10_UP_INFO_Packet)
	//			//			rms_buf, _ = p_10.UnPack(buffer[0:], conn)
	//		default:
	//			fmt.Println("no 0x1BX")
	//			return buffer[len(buffer):]

	//		}

	//		buffer = nil

	//		if len(rms_buf) > 0 {
	//			fmt.Printf("rms_buf buf=%x\n", rms_buf)
	//			ulog.Ul.Debugf("rms_buf buf=%x", rms_buf)

	//			handle_num++
	//			buffer = rms_buf

	//			fmt.Printf("buffer buf=%x\n", buffer)
	//			ulog.Ul.Debugf("buffer buf=%x", buffer)

	//		}

	//		if handle_num > 5 {
	//			rms_buf = nil
	//			buffer = nil
	//			return nil
	//		}

	//	}
	//	//	length := len(buffer)
	//	//	var i int
	//	//	for i = 0; i < length; i = i + 1 {
	//	//		if length < i+ConstHeaderLength+ConstSaveDataLength {
	//	//			break
	//	//		}
	//	//		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
	//	//			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstSaveDataLength])
	//	//			if length < i+ConstHeaderLength+ConstSaveDataLength+messageLength {
	//	//				break
	//	//			}
	//	//			data := buffer[i+ConstHeaderLength+ConstSaveDataLength : i+ConstHeaderLength+ConstSaveDataLength+messageLength]
	//	//			readerChannel <- data

	//	//			i += ConstHeaderLength + ConstSaveDataLength + messageLength - 1
	//	//		}
	//	//	}

	//	//	if i == length {
	//	//	return make([]byte, 0)
	//	//	}
	//	//		return buffer
	//	//fmt.Println("return buffer byte1")
	//	return buffer[0:]
}

func IsValidImei(imei string) bool {
	if len(imei) != 15 {
		return false
	}

	for _, c := range []byte(imei) {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
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
