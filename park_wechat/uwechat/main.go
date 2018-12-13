package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"encoding/xml"
	"io/ioutil"
	"park_wechat/ulog"
	"time"

	"crypto/sha1"

	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"

	"park_wechat/usql"
	"sort"
	"strings"
)

const (
	token = "wechat_ut_go"
)

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

var db *sql.DB
var http_post_num int

func OpenUnDb() {

	fmt.Printf("1111111=%p, %T\n", db, db)

	var err error
	db, err = sql.Open("mysql", "root:unitone@/ips")
	if err != nil {
		ulog.Ul.Fatalf("Open database error: %s\n", err)
	}

	fmt.Printf("222222222=%p, %T\n", db, db)
	//	defer db.Close()

	fmt.Println("ping0")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ping1")
}

func QueryUnDb() {

	//	db, err := sql.Open("mysql", "root:unitone@/ips")
	//	if err != nil {
	//		log.Fatalf("Open database error: %s\n", err)
	//	}
	//	defer db.Close()

	//	fmt.Println("ping0")

	//	err = db.Ping()
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	fmt.Println("ping1")

	//	insert(db)
	fmt.Printf("33333333=%p, %T\n", db, db)

	rows, err := db.Query("select * from sys_user")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("query1")

	defer rows.Close()
	var id, name, pwd, state string

	for rows.Next() {

		fmt.Println("row0")

		err := rows.Scan(&id, &name, &pwd, &state)

		fmt.Println("row1")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("row2")
		log.Println(id, name)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//initial listener and run
func main() {

	ulog.CreateLog()
	ulog.Ul.Debug("park_wechat running ...\n")

	fmt.Printf("park_wechat running ...\n")

	/////////******************************//

	ulog.Ul.Debug("park_wechat db opening ...\n")

	usql.OpenUpsDataBase()
	defer usql.CloseUpsDataBase()

	ulog.Ul.Debug("park_wechat http server opening ...\n")

	go func() {
		for _ = range time.Tick(6 * 1e10) {
			//		ulog.Ul.Debugf("%s:cur con num=%d;dev con in lis=%d;pc con in lis=%d;\n", ulog.CreateDateSting(), cur_conn_num, uconnect.CountDevConnList(), uconnect.CountPcConnList())
			fmt.Printf("%s; http post num=%d; sms coded num=%d\n", ulog.CreateDateSting(), http_post_num, usql.Sms_coded_num)
		}
	}()

	//	http.HandleFunc("/", myHandle)
	//	http.ListenAndServe(":8888", nil)

	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", procRequest)
	http.HandleFunc("/q_plate_number_fee", req_query_plate_number_fee) //设置访问的路由

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)
	}
	log.Println("Wechat Service: Stop!")

	fmt.Printf("main end\n")
}

//func myHandle(w http.ResponseWriter, r *http.Request) {
//	defer r.Body.Close()

//	con, _ := ioutil.ReadAll(r.Body)
//	fmt.Println(string(con))
//	ulog.Ul.Debug(string(con))

//	http_post_num++

//	sms_xml(con)

//	w.Write([]byte("success!"))
//}

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func validateUrl(w http.ResponseWriter, r *http.Request) bool {
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	signatureGen := makeSignature(timestamp, nonce)

	signatureIn := strings.Join(r.Form["signature"], "")
	if signatureGen != signatureIn {
		return false
	}
	echostr := strings.Join(r.Form["echostr"], "")
	fmt.Fprintf(w, echostr)
	return true
}

func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = fromUserName
	textResponseBody.ToUserName = toUserName
	textResponseBody.MsgType = "text"
	textResponseBody.Content = content
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

func procRequest(w http.ResponseWriter, r *http.Request) {

	//	defer r.Body.Close()
	//	con, _ := ioutil.ReadAll(r.Body)
	//	fmt.Println(string(con))

	r.ParseForm()
	if !validateUrl(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	log.Println("Wechat Service: validateUrl Ok!")

	var openId string

	if r.Method == "POST" {

		log.Println("r.Method == POST!")

		textRequestBody := parseTextRequestBody(r)
		openId = textRequestBody.FromUserName
		if textRequestBody != nil {
			fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
				textRequestBody.Content,
				textRequestBody.FromUserName)
			responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
				textRequestBody.FromUserName,
				"Hello, "+textRequestBody.FromUserName)
			if err != nil {
				log.Println("Wechat Service: makeTextResponseBody error: ", err)
				return
			}
			fmt.Fprintf(w, string(responseTextBody))
		}
	}

	// Fetch access_token
	accessToken, expiresIn, err := fetchAccessToken()
	if err != nil {
		log.Println("Get access_token error:", err)
		return
	}
	fmt.Println(accessToken, expiresIn)

	// Post custom service message
	msg := "hello wuxianzhineng"
	err = pushCustomMsg(accessToken, openId, msg)
	if err != nil {
		log.Println("Push custom service message err:", err)
		return
	}
}

////////**********send custom msg************/////////////////////////
const (
	//	token               = "wechat4go"
	appID                = "wxf32be3b6e28d4bc7"
	appSecret            = "bd49f144beb3a1586c437b808ac09e24"
	accessTokenFetchUrl  = "https://api.weixin.qq.com/cgi-bin/token"
	customServicePostUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}

func fetchAccessToken() (string, float64, error) {
	requestLine := strings.Join([]string{accessTokenFetchUrl,
		"?grant_type=client_credential&appid=",
		appID,
		"&secret=",
		appSecret}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", 0.0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0.0, err
	}

	fmt.Println(string(body))

	if bytes.Contains(body, []byte("access_token")) {
		//unmarshal to AccessTokenResponse struct
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			return "", 0.0, err
		}
		return atr.AccessToken, atr.ExpiresIn, nil
	} else {
		//unmarshal to AccessTokenErrorResponse struct
		fmt.Println("return err")
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			return "", 0.0, err
		}
		return "", 0.0, fmt.Errorf("%s", ater.Errmsg)
	}
}

type CustomServiceMsg struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	Text    TextMsgContent `json:"text"`
}

type TextMsgContent struct {
	Content string `json:"content"`
}

func pushCustomMsg(accessToken, toUser, msg string) error {
	csMsg := &CustomServiceMsg{
		ToUser:  toUser,
		MsgType: "text",
		Text:    TextMsgContent{Content: msg},
	}

	body, err := json.MarshalIndent(csMsg, " ", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	postReq, err := http.NewRequest("POST",
		strings.Join([]string{customServicePostUrl, "?access_token=", accessToken}, ""),
		bytes.NewReader(body))
	if err != nil {
		return err
	}

	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

func req_query_plate_number_fee(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./opt/query_plate_number_fee.html")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据,那么执行登陆的逻辑判断
		fmt.Println("plate_num:", r.Form["plate_number"])
		for k, v := range r.Form {
			fmt.Fprintf(w, "key:", k)
			fmt.Fprintf(w, "val:", strings.Join(v, ""))
		}

	}

	// Fetch access_token
	accessToken, expiresIn, err := fetchAccessToken()
	if err != nil {
		log.Println("Get access_token error:", err)
		return
	}
	fmt.Println(accessToken, expiresIn)
}
