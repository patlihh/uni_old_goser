// 10DeviceUpInfo
package uptl

import (
	"fmt"
	"net"
	//	"strconv"
	"usc/umisc"
	//	"usc/unet/uconnect"
	"strings"
	"usc/ulog"
	"usc/usql"
)

type Cmd10_UP_INFO_Packet struct {
	cmd_type         byte
	seq_num          int16
	cmd              byte
	date_time        string
	data_len         int16
	imei_len         byte
	imei             string
	imsi_len         byte
	imsi             string
	device_type      byte
	brand_len        byte
	brand            string
	phone_type_len   byte
	phone_type       string
	phone_num_len    byte
	phone_num        string
	softwareinfo_len byte
	softwareinfo     string
	hardwareinfo_len byte
	hardwartinfo     string
	locationinfo_len byte
	locationinfo     string
	ip_addr          string
}

//解包
func (a *Cmd10_UP_INFO_Packet) UnPack(buffer []byte, conn net.Conn) ([]byte, bool) {

	var package_flag = true
	d_len := len(buffer)
	//	ulog.Ul.Debugf("Cmd0010_UP_INFO UnPcPack buf len=%d", d_len)
	a.seq_num = umisc.Bytes2ToIntLitEndian(buffer[1:3])
	ulog.Ul.Debugf("Cmd0010_UP_INFO UnPcPack buf len=%d;seq_num=%d", d_len, a.seq_num)
	fmt.Printf("Cmd0010_UP_INFO buf len=%d; seq_num=%d!\n", d_len, a.seq_num)
	pos := 3
	if d_len >= 20 {
		a.cmd = buffer[pos]
		pos++
		a.data_len = umisc.Bytes2ToIntLitEndian(buffer[pos : pos+2])
		fmt.Printf("Cmd0010_UP_INFO data_len=%d!\n", a.data_len)
		ulog.Ul.Debugf("Cmd0010_UP_INFO data_len=%d", a.data_len)

		if d_len >= int(a.data_len)+6 {

			pos += 2
			a.imei_len = buffer[pos]
			ulog.Ul.Debugf("a.imei_len=%d", a.imei_len)
			pos++
			if a.imei_len > 0 && a.imei_len < MAX_IMEI_LEN {
				if pos+int(a.imei_len) < d_len {
					a.imei = string(buffer[pos : pos+int(a.imei_len)])
					ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d; imei=%s", a.imei_len, a.imei)
					pos += int(a.imei_len)
					fmt.Printf("a.imei=%s\n", a.imei)
				} else {
					ulog.Ul.Debug("a.imei_len invaild!!!")
					return buffer[d_len:], false
				}

			} else {
				ulog.Ul.Debugf("Cmd0010_UP_INFO not or invaild imei!")
				fmt.Printf("Cmd0010_UP_INFO not or invaild imei!\n")
				return buffer[d_len:], false
			}

			a.imsi_len = buffer[pos]
			ulog.Ul.Debugf("a.imsi_len=%d", a.imsi_len)
			pos++
			if a.imsi_len > 0 {
				if pos+int(a.imsi_len) < d_len {
					a.imsi = string(buffer[pos : pos+int(a.imsi_len)])
					ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d; imsi=%s", a.imsi_len, a.imsi)
					pos += int(a.imsi_len)
					fmt.Printf("a.imsi=%s\n", a.imsi)
				} else {
					ulog.Ul.Debug("a.imsi_len invaild!!!")
					return buffer[d_len:], false
				}

			}

			a.device_type = buffer[pos]
			fmt.Printf("a.device_type=%d\n", a.device_type)
			ulog.Ul.Debugf("a.device_type=%d", a.device_type)

			pos++
			a.brand_len = buffer[pos]
			fmt.Printf("a.brand_len=%d\n", a.brand_len)
			ulog.Ul.Debugf("a.brand_len=%d", a.brand_len)
			pos++
			if a.brand_len > 0 {
				if pos+int(a.brand_len) < d_len {
					a.brand = string(buffer[pos : pos+int(a.brand_len)])
					pos += int(a.brand_len)
					ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d; brand=%s", a.brand_len, a.brand)
					fmt.Printf("a.brand=%s\n", a.brand)
				} else {
					ulog.Ul.Debug("a.brand_len invaild!!!")
					return buffer[d_len:], false
				}

			}

			a.phone_type_len = buffer[pos]
			fmt.Printf("a.phone_type_len=%d\n", a.phone_type_len)
			ulog.Ul.Debugf("a.phone_type_len=%d", a.phone_type_len)

			pos++
			if a.phone_type_len > 0 {
				if pos+int(a.phone_type_len) < d_len {
					a.phone_type = string(buffer[pos : pos+int(a.phone_type_len)])
					pos += int(a.phone_type_len)
					ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d; phone_type=%s", a.phone_type_len, a.phone_type)
					fmt.Printf("a.phone_type=%s\n", a.phone_type)
				} else {
					ulog.Ul.Debug("a.phone_type_len invaild!!!")
					return buffer[d_len:], false
				}

			}

			a.phone_num_len = buffer[pos]
			fmt.Printf("a.phone_num_len=%d\n", a.phone_num_len)
			ulog.Ul.Debugf("a.phone_num_len=%d", a.phone_num_len)
			pos++
			if a.phone_num_len > 0 {
				if pos+int(a.phone_num_len) < d_len {
					a.phone_num = string(buffer[pos : pos+int(a.phone_num_len)])
					pos += int(a.phone_num_len)
					ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d;phone_num=%s", a.phone_num_len, a.phone_num)
					fmt.Printf("a.phone_num=%s\n", a.phone_num)
				} else {
					ulog.Ul.Debug("a.phone_num_len invaild!!!")
					return buffer[d_len:], false
				}
			}

			a.softwareinfo_len = buffer[pos]
			fmt.Printf("a.softwareinfo_len=%d\n", a.softwareinfo_len)
			ulog.Ul.Debugf("a.softwareinfo_len=%d", a.softwareinfo_len)
			pos++
			if a.softwareinfo_len > 0 {
				if pos+int(a.softwareinfo_len) <= d_len {
					a.softwareinfo = string(buffer[pos : pos+int(a.softwareinfo_len)])
					pos += int(a.softwareinfo_len)
					ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d;softwareinfo=%s", a.softwareinfo_len, a.softwareinfo)
					fmt.Printf("a.softwareinfo=%s\n", a.softwareinfo)
				} else {
					ulog.Ul.Debug("a.softwareinfo_len invaild!!!")
					a.softwareinfo = "0000000000"
					a.hardwartinfo = "0000000000"
					a.locationinfo = "00000000"
					package_flag = false
					//	return buffer[d_len:], false
				}

			}

			if package_flag {
				a.hardwareinfo_len = buffer[pos]
				fmt.Printf("a.hardwareinfo_len=%d\n", a.hardwareinfo_len)
				ulog.Ul.Debugf("a.hardwareinfo_len=%d", a.hardwareinfo_len)
				pos++
				if package_flag || a.hardwareinfo_len > 0 {
					if pos+int(a.hardwareinfo_len) <= d_len {
						a.hardwartinfo = string(buffer[pos : pos+int(a.hardwareinfo_len)])
						pos += int(a.hardwareinfo_len)
						ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d pos = %d hardwartinfo=%s ", a.hardwareinfo_len, pos, a.hardwartinfo)
						fmt.Printf("a.hardwartinfo=%s\n", a.hardwartinfo)
					} else {
						ulog.Ul.Debug("a.hardwareinfo_len invaild!!!")
						return buffer[d_len:], false
					}

				}

				if pos >= d_len {
					//		return buffer[d_len:], false
				} else {

					a.locationinfo_len = buffer[pos]
					fmt.Printf("a.locationinfo_len=%d\n", a.locationinfo_len)
					ulog.Ul.Debugf("a.locationinfo_len=%d", a.locationinfo_len)
					pos++
					if package_flag || a.locationinfo_len > 0 {
						if pos+int(a.locationinfo_len) <= d_len {
							a.locationinfo = string(buffer[pos : pos+int(a.locationinfo_len)])
							pos += int(a.locationinfo_len)
							ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d;locationinfo=%s", a.locationinfo_len, a.locationinfo)
							fmt.Printf("a.locationinfo=%s\n", a.locationinfo)
						} else {
							ulog.Ul.Debug("a.locationinfo_len invaild!!!")
							return buffer[d_len:], false
						}

					}
				}
			}

			if IsValidImei(a.imei) == false {
				ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d; imei=%s is not valid!", a.imei_len, a.imei)

				a.SendFailPack(CMD_INVAILD_IMEI, conn)
				return buffer[d_len:], false
			}

			if IsValidBrand(a.brand) == false {
				ulog.Ul.Debugf("Cmd0010_UP_INFO len=%d; brand=%s is not valid!", a.brand_len, a.brand)

				a.SendFailPack(CMD_INVAILD_BRAND, conn)

				return buffer[d_len:], false
			}

			pdb := new(usql.Cmd10_UP_INFO)

			pdb.Imei = a.imei
			pdb.Imsi = a.imsi
			pdb.Datatime = umisc.CreateDateSting()
			pdb.Device_type = a.device_type
			pdb.Phone_type = a.phone_type
			pdb.Phone_num = a.phone_num
			pdb.Softwareinfo = a.softwareinfo
			pdb.Hardwartinfo = a.hardwartinfo
			ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
			pdb.Ip_addr = ip
			pdb.Locationinfo = a.locationinfo

			pdb.InsertUssDeviceInfo()

			if pdb.QueryDeviceInfoUssDeviceInfo(pdb.Imei) {
				ret_buf := append([]byte{0x20}, umisc.Int16ToBytesLitEndian(a.seq_num)...)
				ret_buf = append(ret_buf[0:], []byte{CMD_10}...)
				ret_date := []byte(pdb.Datatime)
				ret_date_len := byte(len(ret_date))
				ret_data := append([]byte{CMD_SET_SUCCESS}, []byte{ret_date_len}...)
				if ret_date_len > 0 {
					ret_data = append(ret_data[0:], ret_date...)
				}
				ret_ip := []byte(pdb.Ip_addr)
				ret_ip_len := byte(len(ret_ip))
				ret_data = append(ret_data[0:], []byte{ret_ip_len}...)
				if ret_ip_len > 0 {
					ret_data = append(ret_data[0:], ret_ip...)
				}
				ret_lac := []byte(pdb.Locationinfo)
				ret_lac_len := byte(len(ret_lac))
				ret_data = append(ret_data[0:], []byte{ret_lac_len}...)
				if ret_lac_len > 0 {
					ret_data = append(ret_data[0:], ret_lac...)
				}

				ret_repair_time := usql.QueryRepairTimeFromByPhoneType(pdb.Phone_type)
				if ret_repair_time >= 0 {
					ret_data = append(ret_data[0:], umisc.IntToBytesLittleEndian(int(ret_repair_time))...)
				}

				ret_data_len := byte(len(ret_data))
				ret_buf = append(ret_buf[0:], []byte{ret_data_len}...)

				conn.Write(append(ret_buf[0:], ret_data[0:]...))
			}

			if !package_flag {
				return buffer[d_len:], false
			}
			return buffer[pos:], true

		}
		return buffer[0:], false
	}

	return buffer[0:], false
}

func (a *Cmd10_UP_INFO_Packet) SendFailPack(ret byte, conn net.Conn) {
	ret_buf := append([]byte{0x20}, umisc.Int16ToBytesLitEndian(a.seq_num)...)
	ret_buf = append(ret_buf[0:], []byte{CMD_10}...)
	ret_buf = append(ret_buf[0:], []byte{0x01}...)
	ret_buf = append(ret_buf[0:], []byte{ret}...)
	conn.Write(ret_buf)
}

func IsValidImei(imei string) bool {
	if len(imei) < 14 || len(imei) > 15 {
		return false
	}

	for _, c := range []byte(imei) {
		if c < '0' || c > 'z' || (c > '9' && c < 'A') || (c > 'Z' && c < 'a') {
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

func IsValidPhoneType(phone_type string) bool {
	if strings.EqualFold(phone_type, "unitone") ||
		strings.EqualFold(phone_type, "UNITONE") ||
		strings.EqualFold(phone_type, "Xiaomi") ||
		strings.EqualFold(phone_type, "ZUUM") {

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
