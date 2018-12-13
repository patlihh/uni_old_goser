// Udce_device_status
package usql

import (
	"fmt"
	//	"strings"
	"usms/ulog"
	"usms/umisc"
)

func (p *Cmd10_UP_INFO) InsertUssDeviceUpInfoTab() bool {

	fmt.Printf("InsertUssDeviceUpInfoTab0\n")

	if len(QueryTabNameFromByPhoneType(p.Phone_type)) > 0 {

		fmt.Printf("phone_type is exit in device_up_info_tab\n")
		ulog.Ul.Debug("phone_type is exit in device_up_info_tab")
		return false
	}

	fmt.Printf("InsertUssDeviceUpInfoTab1\n")

	ulog.Ul.Debug("phone_type is no exit in device_up_info_tab, into phone_type=", p.Phone_type)

	insert_str := fmt.Sprintf("INSERT INTO `device_up_info_tab_name` (PhoneType,TabName,TabCreateDateTime) VALUES(?,?,?)")
	cur_table_name := fmt.Sprintf("%s_device_up_info", p.Phone_type)

	fmt.Printf("InsertUssDeviceUpInfoTab insert_str=%s\n", insert_str)

	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.Phone_type, cur_table_name, umisc.CreateDateSting())
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)

	p.CreateUssDeviceInfoTab()

	return true
}

func QueryTabNameFromByPhoneType(phone_type string) string {

	var tab_name string

	query_str0 := fmt.Sprintf("select `TabName` from usc.device_up_info_tab_name where PhoneType ='%s'", phone_type)
	ulog.Ul.Debug("query_str =", query_str0)
	rows, err := u_db.Query(query_str0)
	if err != nil {
		DbcheckErr(err)
		return ""
	}

	//		fmt.Println("rows =", rows)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tab_name)
	}

	ulog.Ul.Debug("QueryTabNameFromByPhoneType tab_name =", tab_name)
	return tab_name
}

func QueryRepairTimeFromByPhoneType(phone_type string) int32 {

	var repairtime int32

	query_str0 := fmt.Sprintf("select `GuaranteeTime` from usc.device_up_info_tab_name where PhoneType ='%s'", phone_type)
	ulog.Ul.Debug("query_str =", query_str0)
	rows, err := u_db.Query(query_str0)
	if err != nil {
		DbcheckErr(err)
		return -1
	}

	//		fmt.Println("rows =", rows)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&repairtime)
	}

	ulog.Ul.Debug("QueryRepairTimeFromByPhoneType repairtime =", repairtime)
	return repairtime
}

func QueryTabNameByImei(imei string) string {

	var tab_name string
	phone_type := QueryPhoneTypeByImei(imei)

	if len(phone_type) > 0 {
		query_str := fmt.Sprintf("select `TabName` from usc.device_up_info_tab_name where PhoneType ='%s'", phone_type)
		ulog.Ul.Debug("query_str =", query_str)
		rows, err := u_db.Query(query_str)
		if err != nil {
			DbcheckErr(err)
			return ""
		}

		//	fmt.Println("rows =", rows)

		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&tab_name)
		}

	}

	ulog.Ul.Debug("QueryTabNameFromByImei tab_name =", tab_name)
	return tab_name
}
