// bce_device_pos_his
package usql

//
import (
	//	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"log"
	"time"
	"ubus/ulog"
)

type BceDevicePosHis struct {
	Index     int
	DeviceId  string
	DateTime  string
	Longitude string
	Latitude  string
}

var cur_day int
var cur_month time.Month
var cur_year int

var insert_str string
var insert_dce_info bool
var insert_dce_status bool

func CreateBceDevicePosHisTable() {

	if cur_day != time.Now().Day() {

		cur_day = time.Now().Day()

		cur_day_table_name := fmt.Sprintf("bce_device_pos_hisY%dM%02dD%02d",
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day())

		cur_day_table := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`(`Index` SERIAL PRIMARY KEY,`DeviceId` VARCHAR(40) NOT NULL,`Time` VARCHAR(30),`Longitude` VARCHAR(20),`Latitude` VARCHAR(20)) ENGINE=`innodb`,CHARACTER SET=utf8;", cur_day_table_name)

		ulog.Ul.Debug("CreateBceDevicePosHisTable cur_day_table=" + cur_day_table)

		res, err := u_db.Exec(cur_day_table)
		DbcheckErr(err)
		id, err := res.LastInsertId()
		DbcheckErr(err)
		ulog.Ul.Debug(id)

		insert_str = fmt.Sprintf("INSERT INTO `%s` (DeviceId,Time,Longitude,Latitude) VALUES(?,?,?,?)", cur_day_table_name)

	}

	if cur_month != time.Now().Month() {

		cur_month = time.Now().Month()

		cur_month_table_name := fmt.Sprintf("bce_device_pos_hisY%dM%02d",
			time.Now().Year(),
			time.Now().Month())

		cur_month_table := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`(`Index` SERIAL PRIMARY KEY,`DeviceId` VARCHAR(40) NOT NULL,`Time` VARCHAR(30),`Longitude` VARCHAR(20),`Latitude` VARCHAR(20)) ENGINE=`innodb`,CHARACTER SET=utf8;", cur_month_table_name)

		ulog.Ul.Debug("CreateBceDevicePosHisTable cur_month_table=" + cur_month_table)

		res, err := u_db.Exec(cur_month_table)
		DbcheckErr(err)
		id, err := res.LastInsertId()
		DbcheckErr(err)
		ulog.Ul.Debug(id)
	}

	if cur_year != time.Now().Year() {

		cur_year = time.Now().Year()

		cur_year_table_name := fmt.Sprintf("bce_device_pos_hisY%d",
			time.Now().Year())

		cur_year_table := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`(`Index` SERIAL PRIMARY KEY,`DeviceId` VARCHAR(40) NOT NULL,`Time` VARCHAR(30),`Longitude` VARCHAR(20),`Latitude` VARCHAR(20)) ENGINE=`innodb`,CHARACTER SET=utf8;", cur_year_table_name)

		ulog.Ul.Debug("CreateBceDevicePosHisTable cur_year_table=" + cur_year_table)

		res, err := u_db.Exec(cur_year_table)
		DbcheckErr(err)
		id, err := res.LastInsertId()
		DbcheckErr(err)
		ulog.Ul.Debug(id)
	}
}

func (pHis *BceDevicePosHis) InsertBceDevicePosHis() {

	ulog.Ul.Debug("InsertBceDevicePosHis")
	//	fmt.Printf("BceDevicePosHis Longitude==%s\n", pHis.Longitude)
	//	fmt.Printf("BceDevicePosHis Latitude==%s\n", pHis.Latitude)
	//	fmt.Printf("BceDevicePosHis id_adr==%x\n", pHis.DeviceId)
	//	fmt.Printf("BceDevicePosHis id_adr2==%s\n", string(pHis.DeviceId[:]))
	//	fmt.Printf("BceDevicePosHis datetime==%s\n", pHis.DateTime)
	CreateBceDevicePosHisTable()

	//	stmt, err := u_db.Prepare(`INSERT INTO bce_device_pos_his (DeviceId,Time,Longitude,Latitude) VALUES(?,?,?,?)`)
	//	fmt.Println("insert_str = ", insert_str)
	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(pHis.DeviceId, pHis.DateTime, pHis.Longitude, pHis.Latitude)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	ulog.Ul.Debug(id)

	//bce_device_info table not device_id, insert the device_id to bce_device_info table
	//	if insert_dce_info == false {
	//	InsertBceDeviceInfo(string(pHis.DeviceId[0:]))
	//		insert_dce_info = true
	//	}

	//	//update the device_id's status in bce_device_status table
	//	pbs := new(BceDeviceStatus)
	//	if insert_dce_status == false {
	//		//		fmt.Println("pbs=", pbs)
	//		pbs.DeviceId = string(pHis.DeviceId[0:])
	//		pbs.Time = pHis.DateTime
	//		pbs.Longitude = pHis.Longitude
	//		pbs.Latitude = pHis.Latitude
	//		pbs.DeviceNetStatus = '1'
	//		pbs.InsertBceDeviceStatus()
	//		insert_dce_status = true
	//	} else {
	//		pbs.DeviceId = string(pHis.DeviceId[0:])
	//		pbs.Time = pHis.DateTime
	//		pbs.DeviceNetStatus = '1'

	//		//		fmt.Printf("InsertBceDevicePosHis pbs.DeviceNetStatus=%x\n", pbs.DeviceNetStatus)

	//		//		pbs.UpdateBceDeviceStatus()
	//		pbs.InsertBceDeviceStatus()
	//	}

}
