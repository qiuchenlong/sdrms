package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"database/sql"
	"strconv"
)

func (a *MessageState) TableName() string {
	return MessageStateTBName()
}

// CourseQueryParam 用于搜索的类
type MessageStateQueryParam struct {
	BaseQueryParam
	//NameLike string
}

type MessageState struct {
	Id					int
	Msgid    			string `json:"msgid"`
	Mobile	 			string `json:"mobile"`
	State    			int `json:"state"`
	Datetime 			string `json:"datetime"`

	Msg					*Sms `orm:"rel(one)"` //`orm:"reverse(one)"`

	//Creator				*BackendUser `orm:"rel(fk)"` //设置一对多关系
}



func BatchInsert(messageState []MessageState) (int64) { //, id *int

	db_type := beego.AppConfig.String("db_type")
	db_name := beego.AppConfig.String(db_type + "::db_name")
	db_user := beego.AppConfig.String(db_type + "::db_user")
	db_pwd := beego.AppConfig.String(db_type + "::db_pwd")
	db_host := beego.AppConfig.String(db_type + "::db_host")
	db_port := beego.AppConfig.String(db_type + "::db_port")
	db_charset := beego.AppConfig.String(db_type + "::db_charset")

	dbconfig := db_user + ":" + db_pwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=" + db_charset

	//db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test_db?charset=utf8")
	db, err := sql.Open(db_type, dbconfig)
	fmt.Println(err)

	execstring := "INSERT INTO rms_messagestate (msgid, mobile, state, datetime) VALUES" //creator_id
	data := ""

	fmt.Println(time.Now().Unix())
	for i := 0; i < 1; i++ {

		//if i < len(MessageState) - 1 {
		//	onedata := "('1', '1', '1', '1', '1'), "
		//	data = data + onedata
		//} else {
		//	onedata := "('1', '1', '1', '1', '1')"
		//	data = data + onedata
		//}


		//fmt.Println("----->", &id)

		for j := i * 10; j < i*10+len(messageState); j++ {
			if j < i*10+(len(messageState)-1) {
				//id := strconv.Itoa(j)
				//onedata := "(" + id + ", '0', '20180103002930'), "
				onedata := "('"+ messageState[j - i*10].Msgid +"', "+ messageState[j - i*10].Mobile + ", "+ strconv.Itoa(messageState[j - i*10].State) +", "+ messageState[j - i*10].Datetime + "), "
				data = data + onedata
			} else {
				//id := strconv.Itoa(j)
				//onedata := "(" + id + ",'0', '20180103002930')"
				//onedata := "('"+ messageState[j - i*10].Msgid +"', '1', '1', '1', '1')"
				onedata := "('"+ messageState[j - i*10].Msgid +"', "+ messageState[j - i*10].Mobile + ", "+ strconv.Itoa(messageState[j - i*10].State) +", "+ messageState[j - i*10].Datetime + ") "
				data = data + onedata
			}
		}

		//fmt.Println(execstring + data)
		_, err := db.Exec(execstring + data)
		if err != nil {
			fmt.Println(err)
			return 0
		}
	}
	fmt.Println(time.Now().Unix())

	return 1
}