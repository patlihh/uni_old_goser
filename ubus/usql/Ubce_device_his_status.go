// Ubce_device_his_state.go
package usql

//
import (
	//	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"ubus/ulog"
	//	"log"
	//	"time"
	//	"strings"
)

func (p *BceDeviceStatus) InsertBceHisState() {

	//	fmt.Printf("BceDevicePosHis Longitude==%s\n", pHis.Longitude)
	//	fmt.Printf("BceDevicePosHis Latitude==%s\n", pHis.Latitude)
	//	fmt.Printf("BceDevicePosHis id_adr==%x\n", pHis.DeviceId)
	//	fmt.Printf("BceDevicePosHis id_adr2==%s\n", string(pHis.DeviceId[:]))
	//	fmt.Printf("BceDevicePosHis datetime==%s\n", pHis.DateTime)

	stmt, err := u_db.Prepare(`INSERT INTO ubus.bce_device_his_status (DeviceId, Time, IpAdrr,IpPort,DeviceNetStatus, Longitude, Latitude) VALUES(?,?,?,?,?,?,?)`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.DeviceId, p.Time, p.IpAdrr, p.IpPort, string(p.DeviceNetStatus), p.Longitude, p.Latitude)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)

	ulog.Ul.Debug("InsertBceHisState success")

}
