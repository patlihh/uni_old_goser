// Ubus_real_schedul_msg
package usql

import (
	"fmt"
	"strings"
	"ubus/ulog"
)

type BusRealState struct {
	UpdateTime     string
	DriverID       string
	LineID         string
	BusID          string
	CurStationName string
	NowStartTime   string
	EndingTime     string
	BusStatus      byte
	Direction      byte
}

func (p *BusRealState) InsertBusStateMsg() {
	ulog.Ul.Debug("InsertBusStateMsg")

	if QuerBusIdFromBusRealSchedulMsg(p.BusID) == true {
		ulog.Ul.Debug("BusID is exit and update")
		p.UpdateBusStateMsg()
		return
	}

	ulog.Ul.Debugf("BusID no exit, insert id=%s", p.BusID)
	ulog.Ul.Debugf("CurStationName = %s\n", p.CurStationName)

	if len(p.NowStartTime) > 0 {

		ulog.Ul.Debug("IS BEGIN STATION")

		stmt, err := u_db.Prepare(`INSERT INTO ubus.bus_real_schedul_msg (BusID, NowStartTime,DriverID,LineID,CurStationName,BusStatus,Direction) VALUES(?,?,?,?,?,?,?)`)
		defer stmt.Close()

		DbcheckErr(err)
		res, err := stmt.Exec(p.BusID, p.NowStartTime, p.DriverID, p.LineID, p.CurStationName, string(p.BusStatus), string(p.Direction))
		DbcheckErr(err)
		id, err := res.LastInsertId()
		DbcheckErr(err)
		ulog.Ul.Debug(id)
	} else if len(p.EndingTime) > 0 {

		ulog.Ul.Debug("IS END STATION")

		stmt, err := u_db.Prepare(`INSERT INTO ubus.bus_real_schedul_msg (BusID, EndingTime,DriverID,LineID,CurStationName,BusStatus,Direction) VALUES(?,?,?,?,?,?,?)`)
		defer stmt.Close()

		DbcheckErr(err)
		res, err := stmt.Exec(p.BusID, p.EndingTime, p.DriverID, p.LineID, p.CurStationName, string(p.BusStatus), string(p.Direction))
		DbcheckErr(err)
		id, err := res.LastInsertId()
		DbcheckErr(err)
		ulog.Ul.Debug(id)
	} else {
		ulog.Ul.Debug("IS MID STATION")
		stmt, err := u_db.Prepare(`INSERT INTO ubus.bus_real_schedul_msg (BusID, DriverID,LineID,CurStationName,BusStatus,Direction,UpdateTime) VALUES(?,?,?,?,?,?,?)`)
		defer stmt.Close()

		ulog.Ul.Debugf("InsertBusStateMsg p.Busid=%s, p.DriverID=%s, p.LineID=%s, p.CurStationName=%s, BusStatus=%c, Direction=%c, UpdateTime=%s\n", p.BusID, p.DriverID, p.LineID, p.CurStationName, p.BusStatus, p.Direction, p.UpdateTime)

		DbcheckErr(err)
		res, err := stmt.Exec(p.BusID, p.DriverID, p.LineID, p.CurStationName, string(p.BusStatus), string(p.Direction), p.UpdateTime)
		DbcheckErr(err)
		id, err := res.LastInsertId()
		DbcheckErr(err)
		ulog.Ul.Debug(id)
	}
}

func (p *BusRealState) UpdateBusStateMsg() {
	ulog.Ul.Debug("UpdateBusStateMsg")

	if len(p.NowStartTime) > 0 {
		stmt, err := u_db.Prepare(`UPDATE ubus.bus_real_schedul_msg SET NowStartTime=?, DriverID=?, LineID=? ,CurStationName=?,BusStatus=?,Direction=? WHERE BusID = ?`)
		defer stmt.Close()

		ulog.Ul.Debugf("BEGIN UpdateBusStateMsg p.Busid=%s, p.NowStartTime = %s, p.DriverID=%s, p.LineID=%s, p.CurStationName=%s, BusStatus=%c, Direction=%c\n", p.BusID, p.NowStartTime, p.DriverID, p.LineID, p.CurStationName, p.BusStatus, p.Direction)
		DbcheckErr(err)
		res, err := stmt.Exec(p.NowStartTime, p.DriverID, p.LineID, p.CurStationName, string(p.BusStatus), string(p.Direction), p.BusID)
		DbcheckErr(err)
		id, err := res.RowsAffected()
		DbcheckErr(err)
		ulog.Ul.Debug(id)

	} else if len(p.EndingTime) > 0 {
		stmt, err := u_db.Prepare(`UPDATE ubus.bus_real_schedul_msg SET EndingTime=?, DriverID=?, LineID=? ,CurStationName=?,BusStatus=?,Direction=? WHERE BusID = ?`)
		defer stmt.Close()

		ulog.Ul.Debugf("END UpdateBusStateMsg p.BusId=%s, p.DriverID=%s, p.LineID=%s, p.CurStationName=%s, BusStatus=%c, Direction=%c\n", p.BusID, p.DriverID, p.LineID, p.CurStationName, p.BusStatus, p.Direction)
		DbcheckErr(err)
		res, err := stmt.Exec(p.EndingTime, p.DriverID, p.LineID, p.CurStationName, string(p.BusStatus), string(p.Direction), p.BusID)
		DbcheckErr(err)
		id, err := res.RowsAffected()
		DbcheckErr(err)
		ulog.Ul.Debug(id)

	} else {
		stmt, err := u_db.Prepare(`UPDATE ubus.bus_real_schedul_msg SET DriverID=?, LineID=? ,CurStationName=?,BusStatus=?,Direction=?,UpdateTime=? WHERE BusID = ?`)
		defer stmt.Close()

		ulog.Ul.Debugf("MID UpdateBusStateMsg p.BusId=%s, p.DriverID=%s, p.LineID=%s, p.CurStationName=%s, BusStatus=%c, Direction=%c, UpdateTime=%s\n", p.BusID, p.DriverID, p.LineID, p.CurStationName, p.BusStatus, p.Direction, p.UpdateTime)
		DbcheckErr(err)
		res, err := stmt.Exec(p.DriverID, p.LineID, p.CurStationName, string(p.BusStatus), string(p.Direction), p.UpdateTime, p.BusID)
		DbcheckErr(err)
		id, err := res.RowsAffected()
		DbcheckErr(err)
		ulog.Ul.Debug(id)
	}

}

func QuerBusIdFromBusRealSchedulMsg(bus_id string) bool {
	ulog.Ul.Debug("QuerBusIdFromBusRealSchedulMsg")

	query_str := fmt.Sprintf("select `BusID` from ubus.bus_real_schedul_msg where BusID ='%s'", bus_id)
	ulog.Ul.Debug("QuerBusIdFromBusRealSchedulMsg query_str =", query_str)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
	}
	defer rows.Close()

	var id string
	for rows.Next() {
		rows.Scan(&id)
	}

	if strings.EqualFold(id, bus_id) {
		ulog.Ul.Debugf("find BusID=%s", bus_id)
		return true
	} else {
		ulog.Ul.Debugf("not find BusID=%s", bus_id)
		return false
	}
}
