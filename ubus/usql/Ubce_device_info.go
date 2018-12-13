// Ubce_device_info
package usql

import (
	"fmt"
	"strings"
	"ubus/ulog"
)

type BceDeviceInfo struct {
	DeviceId      string
	FirstDateTime string
	FirmId        string
	DeviceState   byte
	DeviceType    byte
	SimNum        string
	IpAdrr        string
	IpPort        int
	Manufacturer  string
	VersionName   string
	PurchaseDate  string
	Remarks       string
}

func (pbce *BceDeviceInfo) InsertBceDeviceInfo() {

	if QueryDeviceIdFromBceDeviceInfo(pbce.DeviceId) == true {

		ulog.Ul.Debug("device_id is exit and update")
		pbce.UpdateBceDeviceInfo()
		return
	}

	ulog.Ul.Debug("device_id no exit, into id=", pbce.DeviceId)

	stmt, err := u_db.Prepare(`INSERT INTO ubus.bce_device_info (DeviceId,FirstDateTime,DeviceState,IpAdrr,IpPort,VersionName) VALUES(?,?,?,?,?,?)`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(pbce.DeviceId, pbce.FirstDateTime, string(pbce.DeviceState), pbce.IpAdrr, pbce.IpPort, pbce.VersionName)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}

func (p *BceDeviceInfo) UpdateBceDeviceInfo() {
	//	ulog.Ul.Debug("UpdateBceDeviceInfo 000")
	stmt, err := u_db.Prepare(`UPDATE ubus.bce_device_info SET IpAdrr=?, IpPort=? ,VersionName=? WHERE DeviceId = ?`)

	defer stmt.Close()

	ulog.Ul.Debugf("UpdateBceDeviceInfo p.IP=%s, p.port=%d, p.Version=%s\n", p.IpAdrr, p.IpPort, p.VersionName)
	DbcheckErr(err)
	res, err := stmt.Exec(p.IpAdrr, p.IpPort, p.VersionName, p.DeviceId)
	DbcheckErr(err)
	id, err := res.RowsAffected()
	DbcheckErr(err)
	fmt.Println(id)
}

func QueryDeviceIdFromBceDeviceInfo(device_id string) bool {

	query_str := fmt.Sprintf("select `DeviceId` from ubus.bce_device_info where DeviceId ='%s'", device_id)
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

func QueryDeviceIdFromBceDeviceInfoByBusNum(bus_num string) string {

	var bus_id string
	var device_id string

	query_str0 := fmt.Sprintf("select `BusId` from ubus.bus_info where BusNum ='%s'", bus_num)
	ulog.Ul.Debug("query_str =", query_str0)
	rows, err := u_db.Query(query_str0)
	if err != nil {
		DbcheckErr(err)
		return ""
	}

	//	fmt.Println("rows =", rows)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&bus_id)
	}

	if len(bus_id) > 0 {
		query_str := fmt.Sprintf("select `DeviceId` from ubus.bus_device where BusId ='%s'", bus_id)
		ulog.Ul.Debug("query_str =", query_str)
		rows, err := u_db.Query(query_str)
		if err != nil {
			DbcheckErr(err)
			return ""
		}

		//	fmt.Println("rows =", rows)

		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&device_id)
		}

	}

	ulog.Ul.Debug("QueryDeviceIdFromBceDeviceInfoByBusNum id =", device_id)
	return device_id
}

func QueryBusNumFromBceDeviceInfoByDeviceId(device_id string) string {
	var bus_id string
	var bus_num string

	ulog.Ul.Debug("QueryBusNumFromBceDeviceInfoByDeviceId id =", device_id)

	query_str0 := fmt.Sprintf("select `BusId` from ubus.bus_device where DeviceId ='%s'", device_id)
	rows, err := u_db.Query(query_str0)
	if err != nil {
		DbcheckErr(err)
		return ""
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&bus_id)
	}

	if len(bus_id) > 0 {
		query_str := fmt.Sprintf("select `BusNum` from ubus.bus_info where BusId ='%s'", bus_id)
		rows, err := u_db.Query(query_str)
		if err != nil {
			DbcheckErr(err)
			return ""
		}

		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&bus_num)
		}
	}

	return bus_num
}
