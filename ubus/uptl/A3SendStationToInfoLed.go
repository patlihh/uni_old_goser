// A3SendLineToBTLed
package uptl

import (
	"fmt"

	"strconv"
)

func IsValidCmdA3Pack(pack []byte) bool {

	b_len := len(pack)
	//pack data len too small
	if b_len < cmd_index_len+4 {
		fmt.Printf("IsValidCmdA3Pack data len=%d too small\n", b_len)
		return false
	}

	//type error
	if pack[cmd_type_pos] != 0x01 {
		fmt.Printf("IsValidCmdA3Pack type error\n")
		return false
	}

	//cmd error
	if pack[cmd_pos] != 0xa3 {
		fmt.Printf("IsValidCmdA3Pack cmd error\n")
		return false
	}

	//direction error

	if pack[cmd_pos+cmd_index_len+1] != '0' && pack[cmd_pos+cmd_index_len+1] != '1' {
		fmt.Printf("IsValidCmdA3Pack direction error\n")
		return false
	}

	line_id_len, _ := strconv.Atoi(fmt.Sprintf("%d", pack[cmd_pos+cmd_index_len+2]))
	//protocol data len err
	if b_len != cmd_index_len+4+line_id_len {
		fmt.Printf("IsValidCmdA3Pack data len error\n")
		return false
	}

	return true
}
