// Ubus_device

package usql

import (
	"fmt"
	"strings"
	"ubus/ulog"
)

type BusDevice struct {
	BusId    string
	DeviceId string
	LineId   string
}

func (pb *BusDevice) InsertBusDevice() {

	if QueryDeviceIdFromBusDevice(pb.DeviceId) == true {

		ulog.Ul.Debug("device_id is exit and update")
		pb.UpdateBusDeviceLineId()
		return
	}

	ulog.Ul.Debug("device_id no exit, into id=", pb.DeviceId)

	stmt, err := u_db.Prepare(`INSERT INTO ubus.bus_device (BusId,DeviceId,LineID) VALUES(?,?,?)`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(pb.BusId, pb.DeviceId, pb.LineId)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}

func (p *BusDevice) UpdateBusDeviceLineId() {
	//	ulog.Ul.Debug("UpdateBceDeviceInfo 000")
	stmt, err := u_db.Prepare(`UPDATE ubus.bus_device SET LineID=? WHERE DeviceId = ?`)

	defer stmt.Close()

	ulog.Ul.Debugf("UpdateBusDeviceLineId p.LineID=%s, p.device_id=%s\n", p.LineId, p.DeviceId)
	DbcheckErr(err)
	res, err := stmt.Exec(p.LineId, p.DeviceId)
	DbcheckErr(err)
	id, err := res.RowsAffected()
	DbcheckErr(err)
	fmt.Println(id)
}

func QueryDeviceIdFromBusDevice(device_id string) bool {

	query_str := fmt.Sprintf("select `DeviceId` from ubus.bus_device where DeviceId ='%s'", device_id)
	ulog.Ul.Debug("query_str =", query_str)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
	}
	defer rows.Close()

	var id string
	for rows.Next() {
		rows.Scan(&id)
	}

	if strings.EqualFold(id, device_id) {
		ulog.Ul.Debug("find device_id")
		return true
	} else {
		ulog.Ul.Debug("not find device_id")
		return false
	}
}
func QueryBusIdFromBusDeviceByDeviceId(device_id string) string {
	ulog.Ul.Debug("QueryBusIdFromBceDeviceInfoByDeviceId id =", device_id)

	query_str := fmt.Sprintf("select `BusId` from ubus.bus_device where DeviceId ='%s'", device_id)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
		return ""
	}

	defer rows.Close()

	var bus_id string
	for rows.Next() {
		err = rows.Scan(&bus_id)
	}

	return bus_id
}
