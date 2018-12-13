// UDispatchResult

package usql

//
import (
	//	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"log"
	//	"time"
	//	"container/list"
	//	"strconv"
	"ubus/ulog"
)

type DispatchResult struct {
	DispatchId         int64
	DispatchResultTime string
	DeviceType         string
	DispatchResult     string
}

func (p *DispatchResult) InsertDispatchResult() {

	ulog.Ul.Debug("InsertDispatchResult")

	stmt, err := u_db.Prepare(`INSERT INTO ubus.dispatch_result (DispatchTimeStamp,DispatchResultTime,DeviceType,DispatchResult) VALUES(?,?,?,?)`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.DispatchId, p.DispatchResultTime, p.DeviceType, p.DispatchResult)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}

func QueryDispatchIDFromDispatchInfo(DispatchID int64) bool {

	query_str := fmt.Sprintf("select `TimeStamp` from ubus.dispatch_info where TimeStamp ='%d'", DispatchID)
	fmt.Println("query_str =", query_str)
	ulog.Ul.Debug("QueryDispatchIDFromDispatchInfo query_str =", query_str)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
	}
	defer rows.Close()

	var id int64
	id = -1
	for rows.Next() {
		rows.Scan(&id)
	}

	if id == DispatchID {
		ulog.Ul.Debugf("find DispatchID = %d\n", DispatchID)
		return true
	} else {
		ulog.Ul.Debugf("not find DispatchID= %d\n", DispatchID)
		return false
	}
}
