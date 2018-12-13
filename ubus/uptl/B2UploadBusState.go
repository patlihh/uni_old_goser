// B2UploadBusState
package uptl

import (
	//	"bytes"
	//	"encoding/binary"
	//	"fmt"
	//	"math"
	"net"
	//	"strings"
	//	"strconv"
	"ubus/ulog"
	//	"ubus/unet/uconnect"
	"ubus/umisc"
	"ubus/usql"
)

type CmdB2 struct {
	Cmd_type              byte
	Cmd_cmd               byte
	data_len              byte
	device_id             string
	driver_id_len         byte
	driver_id             string
	Line_id_len           byte
	Line_id               string
	Line_direction        byte //0 up; 1 down
	Line_station_name_len byte
	Line_station_name     string
	bus_state             byte //0 begin station; 1 end station; 2 mid station
}

//解包
func (b *CmdB2) Unpack(buffer []byte, conn net.Conn) []byte {

	buf_len := len(buffer)

	if buffer[0] == DEV_CMD_TYPE {
		b.Cmd_type = buffer[0]
		b.Cmd_cmd = buffer[1]
		b.data_len = buffer[2]

		ulog.Ul.Debugf("CmdB2_bus_state Unpack b.data_len=%d", b.data_len)

		pos := 3
		if buf_len >= int(b.data_len+3) {
			b.device_id = string(buffer[pos : pos+device_id_len])
			pos += device_id_len
			b.driver_id_len = buffer[pos]
			pos++
			b.driver_id = string(buffer[pos : pos+int(b.driver_id_len)])
			pos += int(b.driver_id_len)
			b.Line_id_len = buffer[pos]
			pos++
			b.Line_id = string(buffer[pos : pos+int(b.Line_id_len)])
			pos += int(b.Line_id_len)
			b.Line_direction = buffer[pos]
			pos++
			b.Line_station_name_len = buffer[pos]
			pos++
			b.Line_station_name = string(buffer[pos : pos+int(b.Line_station_name_len)])
			pos += int(b.Line_station_name_len)
			b.bus_state = buffer[pos]
			pos++

			ulog.Ul.Debugf("CmdB2_bus_state Line_station_name_len=%d", b.Line_station_name_len)

			//NO DRIVER ID
			if !usql.QueryDriverInfoByDriverID(b.driver_id) {
				w_dev_buf := append([]byte{SERVER_TYPE, b.Cmd_cmd}, []byte{NO_DRIVER_ID}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				ulog.Ul.Debugf("b2 Unpack no driver id w_dev_len=%d;w_dev_buf=%x\n", w_dev_len, w_dev_buf)

				return buffer[pos:]
			}
			//NO LINE_ID
			if !usql.QueryLineInfoByLineID(b.Line_id) {
				w_dev_buf := append([]byte{SERVER_TYPE, b.Cmd_cmd}, []byte{NO_LINE_ID}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				ulog.Ul.Debugf("b2 Unpack no line id w_dev_len=%d;w_dev_buf=%x\n", w_dev_len, w_dev_buf)

				return buffer[pos:]
			}

			bus_id := usql.QueryBusIdFromBusDeviceByDeviceId(b.device_id)

			//bus NO LINE_ID
			if !usql.QueryBusLinesfoByLineIDAndBusId(b.Line_id, bus_id) {
				w_dev_buf := append([]byte{SERVER_TYPE, b.Cmd_cmd}, []byte{NO_LINE_ID}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				ulog.Ul.Debugf("b2 Unpack bus no line id w_dev_len=%d;w_dev_buf=%x\n", w_dev_len, w_dev_buf)

				return buffer[pos:]
			}

			if len(bus_id) > 0 {
				pbs := new(usql.BusRealState)
				pbs.BusID = bus_id

				if b.bus_state == BUS_BEGIN_START || b.bus_state == BUS_END_START {
					pbs.NowStartTime = umisc.CreateDateSting()
				} else if b.bus_state == BUS_BEGIN_ARRIVE || b.bus_state == BUS_END_ARRIVE {
					pbs.EndingTime = umisc.CreateDateSting()
				} else {
					pbs.UpdateTime = umisc.CreateDateSting()
				}

				pbs.DriverID = b.driver_id
				pbs.LineID = b.Line_id
				pbs.CurStationName = b.Line_station_name
				pbs.Direction = b.Line_direction
				pbs.BusStatus = b.bus_state

				pbs.InsertBusStateMsg()

			} else {
				w_dev_buf := append([]byte{SERVER_TYPE, b.Cmd_cmd}, []byte{NO_BUS_ID}...)

				w_dev_len, _ := conn.Write(w_dev_buf)

				ulog.Ul.Debugf("b2 Unpack no busid w_dev_len=%d;w_dev_buf=%x\n", w_dev_len, w_dev_buf)

				return buffer[pos:]
			}

			/*****write result to dev************/
			w_dev_buf := append([]byte{SERVER_TYPE, b.Cmd_cmd}, []byte{CMD_SET_SUCCESS}...)

			w_dev_len, _ := conn.Write(w_dev_buf)

			ulog.Ul.Debugf("b2 Unpack w_dev_len=%d;w_dev_buf=%x\n", w_dev_len, w_dev_buf)

			return buffer[pos:]

		}

	}

	return buffer[0:]
}
