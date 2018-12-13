// sms_xml
package main

import (
	"encoding/xml"
	"fmt"
	//	"io/ioutil"
	//	"os"
	"usms/uptl"
)

type Recurlyservers struct {
	Request     RequestConext `xml:"Request"`
	Description string        `xml:",innerxml"`
}

//为了正确解析，go语言的xml包要求struct定义中的所有字段必须是可导出的（即首字母大写）
type RequestConext struct {
	XMLName    xml.Name `xml:"Request"`
	Action     string   `xml:"action"`
	ApiVersion string   `xml:"apiVersion"`
	AppendCode string   `xml:"appendCode"`
	SubAppend  string   `xml:"subAppend"`
	Content    string   `xml:"content"`
	FromNum    string   `xml:"fromNum"`
	RecvTime   string   `xml:"recvTime"`
	SmsType    int      `xml:"smsType"`
}

func sms_xml(str_xml []byte) {
	//	file, err := os.Open("./uss_sms.xml") // For read access.
	//	if err != nil {
	//		fmt.Printf("sms_xml file open error: %v\n", err)
	//		return
	//	}
	//	defer file.Close()
	//	data, err := ioutil.ReadAll(file)
	//	if err != nil {
	//		fmt.Printf("sms_xml file read error: %v\n", err)
	//		return
	//	}
	//		fmt.Println(data)
	fmt.Println(string(str_xml))
	v := RequestConext{}
	err := xml.Unmarshal(str_xml, &v)
	//	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("sms_xml Unmarshal error: %v\n", err)
		return
	}

	fmt.Println(v)
	fmt.Printf("RecvTime: %v\n", v.RecvTime)
	fmt.Printf("FromNum: %v\n", v.FromNum)
	fmt.Printf("Content: %v\n", v.Content)

	p_10 := new(uptl.Cmd10_UP_INFO_Packet)
	p_10.UnPack(v.Content, v.FromNum)

}
