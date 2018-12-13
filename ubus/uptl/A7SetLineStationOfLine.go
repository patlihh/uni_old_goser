//A7SetLineStationOfLine
package uptl

import (
	"fmt"
)

func IsValidCmdA7Pack(pack []byte) bool {

	b_len := len(pack)
	//pack data len too small
	if b_len < cmd_index_len+4 {
		fmt.Printf("IsValidCmdA7Pack data len=%d too small\n", b_len)
		return false
	}

	//type error
	if pack[cmd_type_pos] != 0x01 {
		fmt.Printf("IsValidCmdA7Pack type error\n")
		return false
	}

	//cmd error
	if pack[cmd_pos] != 0xa7 {
		fmt.Printf("IsValidCmdA7Pack cmd error\n")
		return false
	}

	//get data_len
	data_len := Bytes2ToInt(pack[cmd_pos+cmd_index_len+1 : cmd_pos+cmd_index_len+3])

	//get data_len
	line_num := Bytes1ToInt(pack[cmd_pos+cmd_index_len+3 : cmd_pos+cmd_index_len+4])

	fmt.Printf("IsValidCmdA7Pack data_len=%d;line_num=%d\n", data_len, line_num)

	//protocol data len err
	if b_len != cmd_index_len+4+data_len {
		fmt.Printf("IsValidCmdA7Pack data len error\n")
		return false
	}

	return true
}
