// Udatabase
package usql

import (
	"database/sql"
	//	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"usc/ulog"
)

const (
	device_id_len = 10
)

var u_db *sql.DB

var create_test_table string

func OpenUpsDataBase() {
	var err error
	u_db, err = sql.Open("mysql", "usc_n:unitone@tcp(localhost:3306)/usc?charset=utf8")

	//db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")

	if err != nil {
		ulog.Ul.Fatalf("Open database error: %s\n", err)
	}
	//	fmt.Printf("1111111111=%p, %T\n", u_db, u_db)

	u_db.SetMaxOpenConns(3000)
	u_db.SetMaxIdleConns(1500)

	err = u_db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//	InitBceDeviceStatus()
	//	create_test_table = fmt.Sprintf("CREATE TRIGGER `ubus`.`tri_bec_device_pos_his_insert_year` AFTER INSERT ON `ubus`.`bce_device_pos_hisY2015` FOR EACH ROW begin insert into `ubus`.`bce_device_pos_his`(`DeviceId`,`Time`,`Longitude`,`Latitude`) values(new.DeviceId, new.Time, new.Longitude,new.Latitude); end$$", 0x0d, 0x0a)
	//	tx, err := u_db.Begin()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer tx.Rollback()
	//	stmt, err := tx.Prepare("DELIMITER $$")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer stmt.Close() // danger!
	//	_, err = stmt.Exec("CREATE TRIGGER `ubus`.`tri_bec_device_pos_his_insert_year` AFTER INSERT ON `ubus`.`bce_device_pos_hisY2015` FOR EACH ROW begin insert into `ubus`.`bce_device_pos_his`(`DeviceId`,`Time`,`Longitude`,`Latitude`) values(new.DeviceId, new.Time, new.Longitude,new.Latitude); end$$")
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	err = tx.Commit()
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	u_db.Exec()
	//	u_db.Exec("CREATE TRIGGER `ubus`.`tri_bec_device_pos_his_insert_year` AFTER INSERT ON `ubus`.`bce_device_pos_hisY2015` FOR EACH ROW begin insert into `ubus`.`bce_device_pos_his`(`DeviceId`,`Time`,`Longitude`,`Latitude`) values(new.DeviceId, new.Time, new.Longitude,new.Latitude); end$$")
}

func CloseUpsDataBase() {

	err := u_db.Close()

	if err != nil {
		ulog.Ul.Fatal(err)
	}
}

func DbcheckErr(err error) {
	if err != nil {
		ulog.Ul.Debug(err)
		//		panic(err)
	}
}
