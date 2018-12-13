// B3UploadDriverSignRet
package uptl

import (
	//	"bytes"
	//	"encoding/binary"
	"fmt"
	//	"math"
	"net"
	//	"strconv"
	"ubus/ulog"
	//	"ubus/unet/uconnect"
	"ubus/umisc"
	"ubus/usql"
)

type CmdB3 struct {
	Cmd_type      byte
	Cmd_cmd       byte
	Driver_id_len byte
	Driver_id     string
	Sign_ret      byte //
}

//解包
func (b *CmdB3) Unpack(buffer []byte, conn net.Conn) []byte {

	buf_len := len(buffer)

	if buffer[0] == DEV_CMD_TYPE {
		b.Cmd_cmd = buffer[1]
		b.Driver_id_len = buffer[2]
		pos := 3

		ulog.Ul.Debugf("CmdB3 Unpack a.bug=%d", buf_len)

		if buf_len >= int(b.Driver_id_len+4) {
			b.Driver_id = string(buffer[pos : pos+int(b.Driver_id_len)])
			pos += int(b.Driver_id_len)
			b.Sign_ret = buffer[pos]
			pos++

			if usql.QueryDriverInfoByDriverID(b.Driver_id) {
				ds := new(usql.DriverSighRet)
				ds.DriverId = b.Driver_id
				ds.WorkStateTime = umisc.CreateDateSting()
				ds.SighRet = string(b.Sign_ret)
				ds.InsertDriverSighRet()

				/*****write success result to dev************/
				w_dev_buf := append([]byte{SERVER_TYPE, b.Cmd_cmd}, []byte{CMD_SET_SUCCESS}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				fmt.Printf("aa Unpack w_dev_len=%d;w_dev_buf=%x\n", w_dev_len, w_dev_buf)
			} else {
				/*****write fail result to dev************/
				w_dev_buf := append([]byte{SERVER_TYPE, b.Cmd_cmd}, []byte{CMD_SET_FAIL}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				fmt.Printf("aa Unpack w_dev_len=%d;w_dev_buf=%x\n", w_dev_len, w_dev_buf)

			}

			return buffer[pos:]
		}
	}

	return buffer[0:]
}
