// B1GetStationInfo
package uptl

import (
	//	"fmt"
	"net"
)

type CmdB1 struct {
	Longitude [8]byte
	Latitude  [8]byte
	DeviceId  [device_id_len]byte
}

//解包
func (b *CmdB1) Unpack(buffer []byte, conn net.Conn) []byte {
	return buffer[0:]
}
