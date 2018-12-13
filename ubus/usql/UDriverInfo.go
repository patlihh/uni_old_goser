// UDriverInfo
package usql

import (
	"fmt"
)

type DriverInfo struct {
	DriverID string
}

func QueryDriverInfoByDriverID(d_id string) bool {

	query_str := fmt.Sprintf("select `DriverID` from ubus.driver_info where DriverID = '%s'", d_id)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
		return false
	}

	defer rows.Close()

	var id string
	for rows.Next() {
		err = rows.Scan(&id)
		return true
	}

	return false

}
