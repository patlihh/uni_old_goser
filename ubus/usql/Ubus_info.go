// Ubus_info
package usql

import (
	"fmt"
)

type BusInfo struct {
	BusId          string
	BusNum         string
	FirmId         string
	CodeNum        string
	OperatorNum    string
	NuclearLoad    int
	ChassisNum     string
	Manufacturer   string
	EngineNum      string
	PurchaseDate   string
	EngineType     string
	FrameNum       string
	EngineVender   string
	BusType        byte
	BusLength      int
	FuelType       byte
	DischargeLevel byte
	BatteryType    byte
	Capacity       int
	RefConsumption string
	BatteryLife    string
	OutlineColor   string
	CheckDate      string
	AirConBus      byte
	Remarks        string
}

func (p *BusInfo) QueryBusNumFromBusInfoByBusId() {

	query_str := fmt.Sprintf("select `BusNum` from bus_info where BusId =%s", p.BusId)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.BusNum)
	}
}

func (p *BusInfo) QueryBusIdFromBusInfoByBusNum() {

	query_str := fmt.Sprintf("select `BusId` from bus_info where BusNum =%s", p.BusNum)
	rows, err := u_db.Query(query_str)
	if err != nil {
		DbcheckErr(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.BusId)
	}
}
