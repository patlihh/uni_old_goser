// A0SendTimeToInfoLed
package uptl

import (
	"fmt"
	//	"strconv"
)

func IsValidCmdA0Pack(pack []byte) bool {

	b_len := len(pack)
	//pack data len too small
	if b_len < cmd_index_len+4 {
		fmt.Printf("IsValidCmdA0Pack data len=%d too small\n", b_len)
		return false
	}

	//type error
	if pack[cmd_type_pos] != 0x01 {
		fmt.Printf("IsValidCmdA0Pack type error\n")
		return false
	}

	//cmd error
	if pack[cmd_pos] != 0xa0 {
		fmt.Printf("IsValidCmdA0Pack cmd error\n")
		return false
	}

	//	len, _ := strconv.Atoi(fmt.Sprintf("%d", pack[cmd_pos+cmd_index_len+2]))
	//	//protocol data len err
	if b_len != cmd_index_len+10 {
		fmt.Printf("IsValidCmdA0Pack data len error\n")
		return false
	}

	return true
}
