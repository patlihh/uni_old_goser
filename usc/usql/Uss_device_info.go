// Ubce_device_info
package usql

import (
	"fmt"
	"strings"
	"usc/ulog"
	"usc/umisc"
)

var Uss_coded_num int

type Cmd10_UP_INFO struct {
	Imei         string
	Imsi         string
	Datatime     string
	Device_type  byte
	Phone_type   string
	Phone_num    string
	Softwareinfo string
	Hardwartinfo string
	Locationinfo string
	Ip_addr      string
}

func (p *Cmd10_UP_INFO) InsertUssDeviceInfo() bool {

	p.InsertUssDeviceUpInfoTab()

	p.InsertUssPhoneType()

	if QueryIMEIFromUssDeviceInfo(p.Imei) == true {

		ulog.Ul.Debug("imei is exit and update")

		return false
	}

	ulog.Ul.Debug("imei no exit, into imei=", p.Imei)

	if !ExistTabByPhoneType(p.Phone_type) {
		p.CreateUssDeviceInfoTab()
	}

	cur_table_name := fmt.Sprintf("%s_device_up_info", p.Phone_type)
	insert_str := fmt.Sprintf("INSERT INTO `%s` (IMEI,IMSI,DateTime,DeviceType,PhoneType,PhoneNum,SoftWareInfo,HardWareInfo,Ipaddr,LocationInfo) VALUES(?,?,?,?,?,?,?,?,?,?)", cur_table_name)

	//	stmt, err := u_db.Prepare(`INSERT INTO %s (DeviceId,FirstDateTime,DeviceState,IpAdrr,IpPort,VersionName) VALUES(?,?,?,?,?,?)`)
	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.Imei, p.Imsi, umisc.CreateDateSting(), string(p.Device_type), p.Phone_type, p.Phone_num, p.Softwareinfo, p.Hardwartinfo, p.Ip_addr, p.Locationinfo)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)

	Uss_coded_num++

	return true
}

func (p *Cmd10_UP_INFO) CreateUssDeviceInfoTab() {
	//	ulog.Ul.Debug("UpdateBceDeviceInfo 000")
	cur_table_name := fmt.Sprintf("%s_device_up_info", p.Phone_type)

	cur_day_table := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s`(`Index` SERIAL PRIMARY KEY,`IMEI` VARCHAR(20) NOT NULL,`IMSI` VARCHAR(20),`DateTime` VARCHAR(30),`DeviceType` CHAR,`PhoneType` VARCHAR(20),`PhoneNum` VARCHAR(20),`SoftWareInfo` VARCHAR(100),`HardWareInfo` VARCHAR(100),`Ipaddr` VARCHAR(30),`LocationInfo` VARCHAR(150)) ENGINE=`innodb`,CHARACTER SET=utf8;", cur_table_name)

	ulog.Ul.Debug("CreateUssDeviceInfoTab cur_day_table=" + cur_day_table)

	res, err := u_db.Exec(cur_day_table)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	ulog.Ul.Debug(id)
	//	fmt.Println(id)
}

//func (p *Cmd10_UP_INFO) UpdateUssDeviceInfo() {
//	//	ulog.Ul.Debug("UpdateBceDeviceInfo 000")
//	stmt, err := u_db.Prepare(`UPDATE usc.bce_device_info SET IpAdrr=?, IpPort=? ,VersionName=? WHERE DeviceId = ?`)

//	defer stmt.Close()

//	ulog.Ul.Debugf("UpdateBceDeviceInfo p.IP=%s, p.port=%d, p.Version=%s\n", p.IpAdrr, p.IpPort, p.VersionName)
//	DbcheckErr(err)
//	res, err := stmt.Exec(p.IpAdrr, p.IpPort, p.VersionName, p.DeviceId)
//	DbcheckErr(err)
//	id, err := res.RowsAffected()
//	DbcheckErr(err)
//	fmt.Println(id)
//}

func QueryIMEIFromUssDeviceInfo(imei string) bool {

	tab_name := QueryTabNameByImei(imei)

	if len(tab_name) > 0 {
		query_str := fmt.Sprintf("select `IMEI` from usc.`%s` where IMEI ='%s'", tab_name, imei)
		ulog.Ul.Debug("query_str =", query_str)
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

		if strings.EqualFold(id, imei) {
			ulog.Ul.Debug("find imei")
			return true
		} else {
			ulog.Ul.Debug("not find imei")
			return false
		}
	}

	ulog.Ul.Debug("not find imei")
	return false
}

func (p *Cmd10_UP_INFO) QueryDeviceInfoUssDeviceInfo(imei string) bool {

	tab_name := QueryTabNameByImei(imei)

	if len(tab_name) > 0 {
		query_str := fmt.Sprintf("select `IMEI`,`DateTime`, `Ipaddr`, `LocationInfo` from usc.`%s` where IMEI ='%s'", tab_name, imei)
		ulog.Ul.Debug("query_str =", query_str)
		rows, err := u_db.Query(query_str)
		if err != nil {
			DbcheckErr(err)
			return false
		}
		defer rows.Close()

		var id string
		var datetime string
		var ipaddr string
		var lac string
		for rows.Next() {
			rows.Scan(&id, &datetime, &ipaddr, &lac)
		}

		if strings.EqualFold(id, imei) {
			ulog.Ul.Debug("find imei")
			p.Datatime = datetime
			p.Ip_addr = ipaddr
			p.Locationinfo = lac
			return true
		} else {
			ulog.Ul.Debug("not find imei")
			return false
		}
	}

	ulog.Ul.Debug("not find imei")
	return false
}

func ExistTabByPhoneType(phone_type string) bool {

	tab_name := QueryTabNameFromByPhoneType(phone_type)

	if len(tab_name) > 0 {
		query_str := fmt.Sprintf("select * from usc.`%s` where PhoneType ='%s'", tab_name, phone_type)
		ulog.Ul.Debug("query_str =", query_str)
		rows, err := u_db.Query(query_str)
		if err != nil {
			DbcheckErr(err)
			return false
		}
		defer rows.Close()

		if rows.Next() {
			ulog.Ul.Debug("ExistTab OK,phone_type=", phone_type)
			return true
		}
	}

	ulog.Ul.Debug("ExistTab FAIL,phone_type=", phone_type)

	return false
}
