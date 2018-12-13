// Ubce_device_info
package usql

import (
	"fmt"
	//	"strings"
	"usms/ulog"
	//	"usms/umisc"
)

type Dserver_status struct {
	StartTime string
	EndTime   string
	Status    byte
}

func (p *Dserver_status) InsertUssServerStatus() {

	p.InsertUssServerStatusLog()

	ret, ds := QueryUssServerStatus()

	if ret == true {
		ulog.Ul.Debug("server status is exit and update")
		p.UpdateUssServerStatus(ds)
		return
	}

	ulog.Ul.Debug("server status no exit, into status")

	insert_str := fmt.Sprintf("INSERT INTO `server_status` (StartTime,Status) VALUES(?,?)")

	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.StartTime, string(p.Status))
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}

func (p *Dserver_status) InsertUssServerStatusLog() {

	insert_str := fmt.Sprintf("INSERT INTO `server_status_log` (StartTime,Status) VALUES(?,?)")

	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.StartTime, string(p.Status))
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}

func (p *Dserver_status) UpdateUssServerStatus(ds Dserver_status) {
	//	ulog.Ul.Debug("UpdateBceDeviceInfo 000")
	stmt, err := u_db.Prepare(`UPDATE usc.server_status SET StartTime=?, Status=? WHERE StartTime = ?`)

	defer stmt.Close()

	ulog.Ul.Debugf("UpdateUssServerStatus p.StartTime=%s, p.Status=%s, ds.StartTime=%s\n", p.StartTime, string(p.Status), ds.StartTime)
	DbcheckErr(err)
	res, err := stmt.Exec(p.StartTime, string(p.Status), ds.StartTime)
	DbcheckErr(err)
	id, err := res.RowsAffected()
	DbcheckErr(err)
	fmt.Println(id)
}

func QueryUssServerStatus() (bool, Dserver_status) {
	var ds Dserver_status

	query_str := fmt.Sprintf("select * from usc.`server_status`")
	ulog.Ul.Debug("query_str =", query_str)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)

		return false, ds
	}
	defer rows.Close()

	var start_time string
	var status byte

	for rows.Next() {
		rows.Scan(&start_time, &status)
	}

	if len(start_time) > 0 {
		ds.StartTime = start_time
		ds.Status = status
		return true, ds
	} else {
		return false, ds
	}

}
