// bce_device_pos_his
package usql

//
import (
	//	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"log"
	"strings"
	"usms/ulog"
)

func (p *Cmd10_UP_INFO) InsertUssPhoneType() {

	if QueryByImei(p.Imei) == true {

		ulog.Ul.Debug("imei is exit in phone_type tab")
		return
	}

	ulog.Ul.Debug("imei no exit in phone_type, into imei=", p.Imei)

	insert_str := fmt.Sprintf("INSERT INTO `device_phone_type` (IMEI,PhoneType) VALUES(?,?)")
	ulog.Ul.Debug("insert_str=", insert_str)

	//	stmt, err := u_db.Prepare(`INSERT INTO usc.%s (DeviceId,FirstDateTime,DeviceState,IpAdrr,IpPort,VersionName) VALUES(?,?,?,?,?,?)`)
	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.Imei, p.Phone_type)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}

func QueryPhoneTypeByImei(imei string) string {

	var phone_type string

	query_str0 := fmt.Sprintf("select `PhoneType` from usc.device_phone_type where IMEI ='%s'", imei)
	ulog.Ul.Debug("query_str =", query_str0)
	rows, err := u_db.Query(query_str0)
	if err != nil {
		DbcheckErr(err)
		return ""
	}

	//		fmt.Println("rows =", rows)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&phone_type)
	}

	ulog.Ul.Debug("QueryPhoneTypeFromByImei phone_type =", phone_type)
	return phone_type
}

func QueryByImei(imei string) bool {

	var id string

	query_str0 := fmt.Sprintf("select `IMEI` from usc.device_phone_type where IMEI ='%s'", imei)
	ulog.Ul.Debug("query_str =", query_str0)
	rows, err := u_db.Query(query_str0)
	if err != nil {
		DbcheckErr(err)
		return false
	}

	//		fmt.Println("rows =", rows)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
	}
	ulog.Ul.Debug("id=", id)
	ulog.Ul.Debug("imei=", imei)

	if strings.EqualFold(id, imei) {
		ulog.Ul.Debug("phone_type tab find imei")
		return true
	} else {
		ulog.Ul.Debug("phone_type tab not find imei")
		return false
	}

}
