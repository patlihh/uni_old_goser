// UDriverWorkHis
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
	//	"ubus/ulog"
)

type DriverSighRet struct {
	WorkStateTime string
	DriverId      string
	SighRet       string
}

func (p *DriverSighRet) InsertDriverSighRet() {

	stmt, err := u_db.Prepare(`INSERT INTO ubus.driver_sign_his (WorkDatetime,DriverID,SignRet) VALUES(?,?,?)`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.WorkStateTime, p.DriverId, p.SighRet)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}
