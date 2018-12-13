// Uline
package usql

import (
	"fmt"
)

type LineInfo struct {
	LineID         string
	OrgStationID   string
	EndStationID   string
	FirmId         string
	LineType       byte
	LineBusType    byte
	UnmanTicket    byte
	UpMils         int16
	DownMils       int16
	UpSationNum    int16
	DownStationNum int16
	UpBeginTime    string
	UpEndTime      string
	DownBeginTime  string
	DownEndTime    string
	StartPrice     string
	FullPrice      string
	PlineSoundFile string
	BlineSoundFile string
	WlineSoundFile string
	Remarks        string
}

func QueryLineInfoByLineID(l_id string) bool {

	query_str := fmt.Sprintf("select `LineID` from ubus.line_info where LineID = '%s'", l_id)
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

func QueryBusLinesfoByLineIDAndBusId(l_id string, b_id string) bool {

	query_str := fmt.Sprintf("select `LineID` from ubus.bus_lines where LineID = '%s' and BusId='%s'", l_id, b_id)
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
