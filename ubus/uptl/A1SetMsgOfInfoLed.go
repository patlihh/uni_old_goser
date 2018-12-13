// A1ReadAllPark
package uptl

import (
	"fmt"
)

func IsValidCmdA1Pack(pack []byte) bool {

	b_len := len(pack)
	//pack data len too small
	if b_len < cmd_index_len+4 {
		fmt.Printf("IsValidCmdA1Pack data len=%d too small\n", b_len)
		return false
	}

	//type error
	if pack[cmd_type_pos] != 0x01 {
		fmt.Printf("IsValidCmdA1Pack type error\n")
		return false
	}

	//cmd error
	if pack[cmd_pos] != 0xa1 {
		fmt.Printf("IsValidCmdA1Pack cmd error=%d\n", pack[cmd_pos])
		return false
	}

	//info_len
	info_len := Bytes2ToInt(pack[cmd_pos+cmd_index_len+1 : cmd_pos+cmd_index_len+3])
	//protocol data len err
	if b_len != cmd_index_len+5+info_len {
		fmt.Printf("IsValidCmdA1Pack data len error b_len=%d;len=%\n", b_len, cmd_index_len+4+info_len)
		return false
	}

	return true
}
