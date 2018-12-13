// Usys_user
package usql

import (
	"fmt"
	"log"
)

//用户信息表
type SysUser struct {
	id        [11]byte //   NOT NULL COMMENT 'id',
	user_name [20]byte //   DEFAULT NULL COMMENT '用户名',
	password  [20]byte //   DEFAULT NULL COMMENT '密码',
	level     byte     //   DEFAULT NULL COMMENT '状态',
}

func (su *SysUser) QueryAllSysUser() {
	//	fmt.Printf("22222222=%p, %T\n", U_db, U_db)

	rows, err := u_db.Query("select * from user_info")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}
