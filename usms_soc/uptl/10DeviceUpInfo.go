// 10DeviceUpInfo
package uptl

import (
	"fmt"
	"net"
	//	"strconv"
	"usms_soc/umisc"
	//	"usms_soc/unet/uconnect"
	"strings"
	//	"usms_soc/ulog"
	"usms_soc/usql"
)

type Cmd10_UP_INFO_Packet struct {
	imei         string
	imsi         string
	ptype        string
	wtype        string
	hw_info      string
	soft_version string
}

//解包
func (a *Cmd10_UP_INFO_Packet) UnPack(sms_content string, number string) bool {

	fmt.Println(sms_content)
	d_len := len(sms_content)
	if d_len > len("(UNITONE) ") && strings.Contains(sms_content, "(UNITONE)") == true {

		s_new := strings.Trim(sms_content, "(UNITONE)")
		fmt.Println(s_new, len(s_new))

		s := strings.Split(s_new, ";")

		fmt.Println(d_len, s, len(s))

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
				return false
			}

		}

		if IsValidImei(a.imei) {

			pdb := new(usql.Cmd10_UP_INFO)

			pdb.Imei = a.imei
			pdb.Imsi = a.imsi
			pdb.Datatime = umisc.CreateDateSting()
			pdb.Device_type = '0'
			pdb.Phone_type = a.ptype
			pdb.Phone_num = number
			pdb.Softwareinfo = a.soft_version
			pdb.Hardwartinfo = a.hw_info
			//	ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
			pdb.Ip_addr = "0.0.0.0"
			pdb.Locationinfo = "no"

			pdb.InsertUssDeviceInfo()

			return true
		}
	} else {
		if IsValidPhoneNum(number) {
			pdb2 := new(usql.Invaild_sms_record)
			pdb2.Datatime = umisc.CreateDateSting()
			pdb2.Phone_num = number
			pdb2.SmsContent = sms_content
			pdb2.InsertUssInvaildSmsRecord()
		}
	}

	return false

}

func (a *Cmd10_UP_INFO_Packet) SendFailPack(ret byte, conn net.Conn) {
	//	ret_buf := append([]byte{0x20}, umisc.Int16ToBytesLitEndian(a.seq_num)...)
	//	ret_buf = append(ret_buf[0:], []byte{CMD_10}...)
	//	ret_buf = append(ret_buf[0:], []byte{0x01}...)
	//	ret_buf = append(ret_buf[0:], []byte{ret}...)
	//	conn.Write(ret_buf)
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

func IsValidPhoneNum(num string) bool {

	for _, c := range []byte(num) {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}

func IsValidBrand(brand string) bool {
	if strings.EqualFold(brand, "unitone") ||
		strings.EqualFold(brand, "UNITONE") ||
		strings.EqualFold(brand, "Xiaomi") ||
		strings.EqualFold(brand, "ZUUM") {

		return true
	} else {
		return false
	}
}

func IsValidCmdA2Pack(pack []byte) bool {

	//	b_len := len(pack)
	//	//pack data len too small
	//	if b_len < cmd_index_len+4 {
	//		fmt.Printf("IsValidCmdA2Pack data len=%d too small\n", b_len)
	//		return false
	//	}

	//	//type error
	//	if pack[cmd_type_pos] != 0x01 {
	//		fmt.Printf("IsValidCmdA2Pack type error =%d\n", pack[cmd_type_pos])
	//		return false
	//	}

	//	//cmd error
	//	if pack[cmd_pos] != 0xa2 {
	//		fmt.Printf("IsValidCmdA2Pack cmd error = %d\n", pack[cmd_pos])
	//		return false
	//	}

	//	//get info_len
	//	info_len := Bytes2ToInt(pack[cmd_pos+cmd_index_len+1 : cmd_pos+cmd_index_len+3])
	//	//protocol data len err
	//	if b_len != cmd_index_len+4+info_len {
	//		fmt.Printf("IsValidCmdA2Pack data len error b_len=%d; len2=%d\n", b_len, cmd_index_len+4+info_len)
	//		return false
	//	}

	return true
}
