// A2SendMsgToInfoLed
package uptl

import (
	"fmt"
	"net"
	//	"strconv"
	"ubus/umisc"
	//	"ubus/unet/uconnect"
	"ubus/ulog"
	"ubus/usql"
)

type CmdA2_INFO_LED_RET struct {
	cmd_type   byte
	cmd        byte
	schedul_id int64
	ret        byte
}

//解包
func (a *CmdA2_INFO_LED_RET) UnPcPack(buffer []byte, conn net.Conn) ([]byte, bool) {

	d_len := len(buffer)
	ulog.Ul.Debugf("CmdA2_INFO_LED_RET UnPcPack buf len=%d", d_len)

	if d_len >= 11 {
		a.cmd_type = buffer[0]
		a.cmd = buffer[1]
		a.schedul_id = umisc.Bytes8ToIntOfBigEndian(buffer[2:10])
		a.ret = buffer[10]

		ulog.Ul.Debugf("CmdA2_INFO_LED_RET Unpack a.schedul_id=%d;a.ret=%d\n", a.schedul_id, a.ret)

		/******write db******/
		if usql.QueryDispatchIDFromDispatchInfo(a.schedul_id) {

			pdb := new(usql.DispatchResult)

			pdb.DispatchId = a.schedul_id
			pdb.DispatchResultTime = umisc.CreateDateSting()
			pdb.DeviceType = fmt.Sprintf("%d", INFO_LED_CMD_TYPE)
			pdb.DispatchResult = fmt.Sprintf("%d", a.ret)

			pdb.InsertDispatchResult()
		} else {

			ulog.Ul.Debugf("CmdA2_INFO_LED_RET not web schedul ret = %d\n", a.ret)

		}

		return buffer[11:], true

	}

	return buffer[0:], false
}

func IsValidCmdA2Pack(pack []byte) bool {

	b_len := len(pack)
	//pack data len too small
	if b_len < cmd_index_len+4 {
		fmt.Printf("IsValidCmdA2Pack data len=%d too small\n", b_len)
		return false
	}

	//type error
	if pack[cmd_type_pos] != 0x01 {
		fmt.Printf("IsValidCmdA2Pack type error =%d\n", pack[cmd_type_pos])
		return false
	}

	//cmd error
	if pack[cmd_pos] != 0xa2 {
		fmt.Printf("IsValidCmdA2Pack cmd error = %d\n", pack[cmd_pos])
		return false
	}

	//get info_len
	info_len := Bytes2ToInt(pack[cmd_pos+cmd_index_len+1 : cmd_pos+cmd_index_len+3])
	//protocol data len err
	if b_len != cmd_index_len+4+info_len {
		fmt.Printf("IsValidCmdA2Pack data len error b_len=%d; len2=%d\n", b_len, cmd_index_len+4+info_len)
		return false
	}

	return true
}
