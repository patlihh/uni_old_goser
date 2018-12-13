// A9SetStationInfo
package uptl

import (
	"fmt"
)

func IsValidCmdA9Pack(pack []byte) bool {

	b_len := len(pack)
	//pack data len too small
	if b_len < cmd_index_len+4 {
		fmt.Printf("IsValidCmdA9Pack data len=%d too small\n", b_len)
		return false
	}

	//type error
	if pack[cmd_type_pos] != 0x01 {
		fmt.Printf("IsValidCmdA9Pack type error\n")
		return false
	}

	//cmd error
	if pack[cmd_pos] != 0xa9 {
		fmt.Printf("IsValidCmdA9Pack cmd error\n")
		return false
	}

	//get info_len
	data_len := Bytes2ToInt(pack[cmd_pos+cmd_index_len+1 : cmd_pos+cmd_index_len+3])
	//protocol data len err
	if b_len != cmd_index_len+4+data_len {
		fmt.Printf("IsValidCmdA9Pack data len error b_len=%d; len2=%d\n", b_len, cmd_index_len+4+data_len)
		return false
	}

	return true
}
