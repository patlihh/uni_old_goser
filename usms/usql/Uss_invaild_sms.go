// Uss_invaild_sms
package usql

import (
	"fmt"
	//	"strings"
	"usms/ulog"
	//	"usms/umisc"
)

type Invaild_sms_record struct {
	Phone_num  string
	Datatime   string
	SmsContent string
}

func (p *Invaild_sms_record) InsertUssInvaildSmsRecord() {

	ulog.Ul.Debug("Invaild sms record insert; num=", p.Phone_num)

	insert_str := fmt.Sprintf("INSERT INTO `InvaildSmsRecord` (PhoneNum,DateTime,SmsContent) VALUES(?,?,?)")

	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.Phone_num, p.Datatime, p.SmsContent)
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}
