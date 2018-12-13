// bce_protocol_code
package usql

//
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"log"
	//	"time"
	"container/list"
	"strconv"
	"ubus/ulog"
)

type BceProtocolCode struct {
	PIndex      int
	ReqDateTime string
	RspDateTime string
	ReqType     byte
	BusNum      string
	CmdIndex    int
	CMD         byte
	DataType    byte
	Data        []byte
	WorkStatus  byte
	RspCode     byte
	Remarks     string
}

var client_req_lists = list.New()

//var bce_protocol_coded_lists = list.New()
func GetClientReqNum() int {
	return client_req_lists.Len()
}

func GetClientReqFirst() BceProtocolCode {

	return client_req_lists.Front().Value.(BceProtocolCode)

}

func RemoveClientReqFirst() {

	client_req_lists.Remove(client_req_lists.Front())

}

func RemoveClientReqByPIndex(pindex int) {
	var c BceProtocolCode
	var n *list.Element

	for e := client_req_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(BceProtocolCode)

		//		fmt.Println("list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("conn nw1=", conn.Uconn.RemoteAddr().String())

		if c.PIndex == pindex {
			client_req_lists.Remove(e)
			fmt.Println(ulog.CreateDateWithNanoSting(), "::", "client req index=", strconv.Itoa(c.PIndex), "removed ok. num=", client_req_lists.Len())
		}
	}
}

func AddClientReqToList(req BceProtocolCode) bool {
	var c BceProtocolCode
	var n *list.Element

	for e := client_req_lists.Front(); e != nil; e = n {

		n = e.Next()

		c = e.Value.(BceProtocolCode)

		//		fmt.Println("list c nw1=", c.Uconn.RemoteAddr().String())
		//		fmt.Println("conn nw1=", conn.Uconn.RemoteAddr().String())

		if c.PIndex == req.PIndex {
			//			fmt.Println("req is exist in list")
			return false
		}

	}

	client_req_lists.PushBack(req)

	fmt.Printf("%s::client req BusNum=%s;CMD=%X add to list ok. req list num=%d\n", ulog.CreateDateWithNanoSting(), req.BusNum, req.CMD, client_req_lists.Len())

	return true
}

func QueryClientReqProtocolCode() int {

	if GetClientReqNum() > 0 {
		//		fmt.Println("GetClientReqNum=", GetClientReqNum())
		return GetClientReqNum()
	}

	rows, err := u_db.Query("select `pIndex`,`BusNum`,`CMD`,`Data` from protocol_code where ReqType = '0' AND WorkStatus = '0'")
	if err != nil {
		DbcheckErr(err)
		return 0
	}

	defer rows.Close()

	columns, err := rows.Columns()
	DbcheckErr(err)

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]

	}

	var req_tmp BceProtocolCode

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		DbcheckErr(err)

		for i, col := range values {

			switch columns[i] {
			case "pIndex":
				req_tmp.PIndex, err = strconv.Atoi(string(col))
			case "BusNum":
				req_tmp.BusNum = string(col)
			case "CMD":
				req_tmp.CMD = []byte(string(col))[0]
			case "Data":
				req_tmp.Data = []byte(string(col))
			default:
				break

			}

		}
		AddClientReqToList(req_tmp)
	}

	return GetClientReqNum()
	//	var value string

	//	var cmd_byte sql.RawBytes
	//	var data_byte []sql.RawBytes
	//	var req_tmp BceProtocolCode
	//	for rows.Next() {
	//		err = rows.Scan(&req_tmp.Index, &req_tmp.BusNum, &cmd_byte, &data_byte)
	////		copy(, cmd_byte[0:1])
	//		AddClientReqToList(req_tmp)
	//	}

}

func (p *BceProtocolCode) InsertBceProtocolCodeReq() {

	stmt, err := u_db.Prepare(`INSERT INTO protocol_code (ReqDateTime,ReqType,BusNum,CMD,Data,WorkStatus) VALUES(?,?,?,?,?,?)`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.ReqDateTime, string('1'), p.BusNum, p.CMD, p.Data, string('0'))
	DbcheckErr(err)
	id, err := res.LastInsertId()
	DbcheckErr(err)
	fmt.Println(id)
}

func (p *BceProtocolCode) UpdateBceProtocolCodeRsp() {

	fmt.Printf("UpdateBceProtocolCodeRsp index==%d\n", p.PIndex)
	fmt.Printf("UpdateBceProtocolCodeRsp RspDateTime==%s\n", p.RspDateTime)
	fmt.Printf("UpdateBceProtocolCodeRsp WorkStatus==%c\n", p.WorkStatus)
	fmt.Printf("UpdateBceProtocolCodeRsp RspCode==%c\n", p.RspCode)
	fmt.Printf("UpdateBceProtocolCodeRsp Remarks==%s\n", p.Remarks)
	//	fmt.Printf("BceDevicePosHis id_adr==%x\n", pHis.DeviceId)
	//	fmt.Printf("BceDevicePosHis id_adr2==%s\n", string(pHis.DeviceId[:]))
	//	fmt.Printf("BceDevicePosHis datetime==%s\n", pHis.DateTime)
	//	CreateBceDevicePosHisTable()

	stmt, err := u_db.Prepare(`UPDATE ubus.protocol_code SET RspDateTime=?,WorkStatus=?,RspCode=? WHERE pIndex=?`)
	//	stmt, err := u_db.Prepare(insert_str)
	defer stmt.Close()

	DbcheckErr(err)
	res, err := stmt.Exec(p.RspDateTime, string(p.WorkStatus), string(p.RspCode), p.PIndex)
	DbcheckErr(err)
	id, err := res.RowsAffected()
	DbcheckErr(err)
	fmt.Println(id)
}
