package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//	"os"
	//	"os/exec"
	//	"path/filepath"
	"io/ioutil"
	"net/http"
	"time"
	"usms/ulog"
	//	"usms/umisc"
	//	"usms/unet"
	"usms/usql"
)

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
	ulog.Ul.Debug("usms running ...\n")

	fmt.Printf("usms running ...\n")

	////////////////////***********************//

	//	lf, err := os.OpenFile("log_bus.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	//	if err != nil {
	//		os.Exit(1)
	//	}
	//	defer lf.Close()

	//	//log file
	//	l := log.New(lf, "", os.O_APPEND)

	//	if os.Getppid() != 1 { //ÅÐ¶Ïµ±ÆäÊÇ·ñÊÇ×Ó½ø³Ì£¬µ±¸¸½ø³ÌreturnÖ®ºó£¬×Ó½ø³Ì»á±» ÏµÍ³1 ºÅ½ø³Ì½Ó¹Ü

	//		filePath, _ := filepath.Abs(os.Args[0]) //½«ÃüÁîÐÐ²ÎÊýÖÐÖ´ÐÐÎÄ¼þÂ·¾¶×ª»»³É¿ÉÓÃÂ·¾¶
	//		cmd := exec.Command(filePath, os.Args[1:]...)
	//		//½«ÆäËûÃüÁî´«ÈëÉú³É³öµÄ½ø³Ì
	//		cmd.Stdin = os.Stdin //¸øÐÂ½ø³ÌÉèÖÃÎÄ¼þÃèÊö·û£¬¿ÉÒÔÖØ¶¨Ïòµ½ÎÄ¼þÖÐ
	//		cmd.Stdout = os.Stdout
	//		cmd.Stderr = os.Stderr
	//		cmd.Start() //¿ªÊ¼Ö´ÐÐÐÂ½ø³Ì£¬²»µÈ´ýÐÂ½ø³ÌÍË³ö

	//		if err != nil {
	//			l.Printf("%s process start fail! err=%s\n", time.Now().Format("2006-01-02 15:04:05"), err)

	//		}
	//		l.Printf("%s process start up! err=%s\n", time.Now().Format("2006-01-02 15:04:05"), err)
	//		//		err = cmd.Wait()
	//		//		l.Printf("%s process exit!", time.Now().Format("2006-01-02 15:04:05"), err)

	//		return
	//	}

	/////////******************************//

	ulog.Ul.Debug("usms db opening ...\n")

	usql.OpenUpsDataBase()
	defer usql.CloseUpsDataBase()

	//	pSs := new(usql.Dserver_status)
	//	pSs.StartTime = umisc.CreateDateSting()
	//	pSs.Status = '1'
	//	pSs.InsertUssServerStatus()

	//	OpenUnDb()
	//	QueryUnDb()
	ulog.Ul.Debug("usms http server opening ...\n")

	go func() {
		for _ = range time.Tick(6 * 1e10) {
			//		ulog.Ul.Debugf("%s:cur con num=%d;dev con in lis=%d;pc con in lis=%d;\n", ulog.CreateDateSting(), cur_conn_num, uconnect.CountDevConnList(), uconnect.CountPcConnList())
			fmt.Printf("%s; http post num=%d; sms coded num=%d\n", ulog.CreateDateSting(), http_post_num, usql.Sms_coded_num)
		}
	}()

	http.HandleFunc("/", myHandle)
	http.ListenAndServe(":8888", nil)

	fmt.Printf("main end\n")
}

func myHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	con, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(con))
	ulog.Ul.Debug(string(con))

	http_post_num++

	sms_xml(con)

	w.Write([]byte("success!"))
}
