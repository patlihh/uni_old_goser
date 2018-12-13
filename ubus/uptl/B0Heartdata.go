// heartdata
package uptl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"strconv"
	"ubus/ulog"
	"ubus/unet/uconnect"
	"ubus/usql"
)

const (
	cmd_index_len = 4
	cmd_type_pos  = 0 //4
	cmd_pos       = 1 //5
	device_id_len = 10
)

type CmdB0 struct {
	Longitude [8]byte
	Latitude  [8]byte
	DeviceId  [device_id_len]byte
}

type CmdB0_PC struct {
	Mac_id []byte
}

//解包
func (b0 *CmdB0) Unpack(buffer []byte, conn net.Conn) []byte {

	buf_len := len(buffer)
	ulog.Ul.Debugf("CmdB0 Unpack buf len=%d", buf_len)

	if buffer[0] == DEV_CMD_TYPE {
		if buf_len > 2 {

			data_len := buffer[2]
			if buf_len >= (int)(data_len+3) {
				copy(b0.Longitude[:], buffer[3:11])
				copy(b0.Latitude[:], buffer[11:19])
				copy(b0.DeviceId[:], buffer[19:19+device_id_len])

				ulog.Ul.Debugf("Longitude=%s", strconv.FormatFloat(math.Float64frombits(binary.LittleEndian.Uint64(b0.Longitude[:])), 'f', 6, 64))
				ulog.Ul.Debugf("Latitude=%s", strconv.FormatFloat(math.Float64frombits(binary.LittleEndian.Uint64(b0.Latitude[:])), 'f', 6, 64))
				ulog.Ul.Debugf("id_adr=%s", b0.DeviceId)

				if uconnect.AddDevConnAndMacAdrToList(buffer[0], conn, b0.DeviceId) {

					ip, port_str, _ := net.SplitHostPort(conn.RemoteAddr().String())
					port, _ := strconv.Atoi(port_str)

					pbce := new(usql.BceDeviceInfo)
					pbce.DeviceId = fmt.Sprintf("%s", b0.DeviceId[0:])
					pbce.FirstDateTime = ulog.CreateDateSting()
					pbce.IpAdrr = ip
					pbce.IpPort = port
					pbce.DeviceState = '1'
					pbce.VersionName = fmt.Sprintf("%s", buffer[data_len:])

					pbce.InsertBceDeviceInfo()

					pbs := new(usql.BceDeviceStatus)
					pbs.DeviceId = fmt.Sprintf("%s", b0.DeviceId[0:])
					pbs.Time = ulog.CreateDateSting()
					pbs.Longitude = strconv.FormatFloat(math.Float64frombits(binary.LittleEndian.Uint64(b0.Longitude[:])), 'f', 6, 64)
					pbs.Latitude = strconv.FormatFloat(math.Float64frombits(binary.LittleEndian.Uint64(b0.Latitude[:])), 'f', 6, 64)
					pbs.IpAdrr = ip
					pbs.IpPort = port
					pbs.DeviceNetStatus = '1'
					pbs.InsertBceDeviceStatus()
				} else {
					pbs := new(usql.BceDeviceStatus)
					pbs.DeviceId = fmt.Sprintf("%s", b0.DeviceId[0:])
					if !usql.QueryNetStatusFromBceDeviceStatus(pbs.DeviceId) {
						pbs.Time = ulog.CreateDateSting()
						pbs.DeviceNetStatus = '1'
						pbs.UpdateBceDeviceStatus()
					}

				}

				//b0 Longitude==D7A3703D0A475940
				//b0 Latitude==EC51B81E85636040

				// A5 TEST TMP BEGIN
				//			fmt.Println("pack A8 and send")
				//			p_a9 := new(CmdA9)
				//			_, err := conn.Write(p_a9.Pack())
				//			if err != nil {
				//				fmt.Println("Error A5 write=", err)

				ppHis := new(usql.BceDevicePosHis)

				ppHis.DeviceId = fmt.Sprintf("%s", b0.DeviceId[0:])
				ppHis.DateTime = ulog.CreateDateSting()
				ppHis.Longitude = strconv.FormatFloat(math.Float64frombits(binary.LittleEndian.Uint64(b0.Longitude[:])), 'f', 6, 64)
				ppHis.Latitude = strconv.FormatFloat(math.Float64frombits(binary.LittleEndian.Uint64(b0.Latitude[:])), 'f', 6, 64)

				ppHis.InsertBceDevicePosHis()

				/*****write result to dev************/
				w_dev_buf := append([]byte{SERVER_TYPE, buffer[1]}, []byte{CMD_SET_SUCCESS}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				ulog.Ul.Debugf("Unpack w_dev_len=%d;w_dev_buf=%x", w_dev_len, w_dev_buf)

				//			}
				//A5 TEST TMP END

				//TEST MYSQL BEGIN
				//				pUser := new(usql.SysUser)
				//				pUser.QueryAllSysUser()

				//			pController := new(usql.Controller)
				//			pController.QueryAllController()
				//TEST MYSQL END

				return buffer[data_len+3:]
			}
		}
	}

	return buffer[0:]
}

//解包
func (b0 *CmdB0_PC) Unpack(buffer []byte, conn net.Conn) []byte {

	buf_len := len(buffer)
	ulog.Ul.Debugf("CmdB0_PC Unpack buf len=%d", buf_len)

	if buffer[0] == PC_CMD_TYPE {
		if buf_len > 2 {

			data_len := buffer[2]
			if buf_len >= (int)(data_len+3) {
				copy(b0.Mac_id[:], buffer[3:3+data_len])

				ulog.Ul.Debugf("Mac_id=%X", b0.Mac_id)

				/*****write result to PC************/
				w_dev_buf := append([]byte{SERVER_TYPE, buffer[1]}, []byte{CMD_SET_SUCCESS}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				ulog.Ul.Debugf("Unpack w_dev_len=%d;w_dev_buf=%x", w_dev_len, w_dev_buf)
				return buffer[data_len+3:]
			}
		}
	}

	return buffer[0:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//4字节转换成整形
func Bytes4ToInt(b []byte) int {

	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//2字节转换成整形
func Bytes2ToInt(b []byte) int {

	bytesBuffer := bytes.NewBuffer(b)

	var x int16
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//1字节转换成整形
func Bytes1ToInt(b []byte) int {

	bytesBuffer := bytes.NewBuffer(b)

	var x int8
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
