// Udce_device_status
package usql

import (
	"fmt"
	"strings"
	"ubus/ulog"
	"ubus/umisc"
)

type BceDeviceStatus struct {
	DeviceId        string
	Time            string
	IpAdrr          string
	IpPort          int
	DeviceNetStatus byte
	Longitude       string
	Latitude        string
}

func QueryDeviceIdFromBceDeviceStatus(device_id string) bool {

	query_str := fmt.Sprintf("select `DeviceId` from ubus.bce_device_status where DeviceId ='%s'", device_id)
	ulog.Ul.Debug("QueryDeviceIdFromBceDeviceStatus query_str =", query_str)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
		return false
	}

	defer rows.Close()

	var id string
	for rows.Next() {
		rows.Scan(&id)
	}

	if strings.EqualFold(id, device_id) {
		ulog.Ul.Debug("bce_device_status find device_id")
		return true
	} else {
		ulog.Ul.Debug("bce_device_status not find device_id")
		return false
	}
}

func QueryNetStatusFromBceDeviceStatus(device_id string) bool {

	query_str := fmt.Sprintf("select `DeviceNetStatus` from ubus.bce_device_status where DeviceId ='%s'", device_id)
	ulog.Ul.Debug("QueryNetStatusFromBceDeviceStatus query_str =", query_str)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
		return false
	}

	defer rows.Close()

	var state string
	for rows.Next() {
		rows.Scan(&state)
	}

	if strings.EqualFold(state, "1") {
		ulog.Ul.Debug("bce_device_status is connected status!")
		return true
	} else {
		ulog.Ul.Debug("bce_device_status is not connected status!")
		return false
	}
}

func (p *BceDeviceStatus) InsertBceDeviceStatus() {

	//	fmt.Printf("BceDevicePosHis Longitude==%s\n", pHis.Longitude)
	//	fmt.Printf("BceDevicePosHis Latitude==%s\n", pHis.Latitude)
	//	fmt.Printf("BceDevicePosHis id_adr==%x\n", pHis.DeviceId)
	//	fmt.Printf("BceDevicePosHis id_adr2==%s\n", string(pHis.DeviceId[:]))
	//	fmt.Printf("BceDevicePosHis datetime==%s\n", pHis.DateTime)
	ulog.Ul.Debug("InsertBceDeviceStatus p.DeviceId=", p.DeviceId)
	//	p.InsertBceHisState()

	if QueryDeviceIdFromBceDeviceStatus(p.DeviceId) {
		p.UpdateBceDeviceStatus()
		return
	}

	stmt, err := u_db.Prepare(`INSERT INTO ubus.bce_device_status (DeviceId, Time, IpAdrr,IpPort,DeviceNetStatus, Longitude, Latitude) VALUES(?,?,?,?,?,?,?)`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.DeviceId, p.Time, p.IpAdrr, p.IpPort, string(p.DeviceNetStatus), p.Longitude, p.Latitude)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)

	ulog.Ul.Debug("InsertBceDeviceStatus success")
}

func (p *BceDeviceStatus) UpdateBceDeviceStatus() {
	//	fmt.Println("UpdateBceDeviceStatus 000")
	stmt, err := u_db.Prepare(`UPDATE ubus.bce_device_status SET Time=?, DeviceNetStatus=? WHERE DeviceId = ?`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	ulog.Ul.Debugf("UpdateBceDeviceStatus p.DeviceNetStatus=%x\n", p.DeviceNetStatus)
	DbcheckErr(err)
	res, err := stmt.Exec(p.Time, string(p.DeviceNetStatus), p.DeviceId)
	DbcheckErr(err)
	id, err := res.RowsAffected()
	DbcheckErr(err)
	fmt.Println(id)
}

func InitBceDeviceStatus() {
	//	fmt.Println("InitBceDeviceStatus 000")
	stmt, err := u_db.Prepare(`UPDATE ubus.bce_device_status SET Time=?, DeviceNetStatus=?`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	ulog.Ul.Debugf("InitBceDeviceStatus p.DeviceNetStatus=%x\n", 0x30)
	DbcheckErr(err)
	res, err := stmt.Exec(umisc.CreateDateSting(), string("0"))
	DbcheckErr(err)
	id, err := res.RowsAffected()
	DbcheckErr(err)
	fmt.Println(id)
}
