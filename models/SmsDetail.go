package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"database/sql"
	"strconv"
	"strings"
	"github.com/astaxie/beego/orm"
)


// 短信明细表
func (a *SmsDetail) TableName() string {
	return SmsDetailTBName()
}

// CourseQueryParam 用于搜索的类
type SmsDetailQueryParam struct {
	BaseQueryParam
	//NameLike string
}



type SmsDetail struct {
	Id					int                      // 1、主键
	Mobile	 			string `json:"mobile"`   // 3、手机号
	Number				int `json:"number"`      // 4、短信条数
	Corporator			string `json:"corporator"`   // 5、运营商
	Location			string `json:"location"`   // 6、归属地
	SubmitState    		int `json:"submit_state"`    // 7、提交状态
	SubmitDatetime 		string `json:"submit_datetime"`  // 8、提交时间
	SendState			int `json:"send_state"`    // 9、发送状态
	SendDatetime		string `json:"send_datetime"`  // 10、发送时间
	ReportState			int `json:"report_state"`  // 11、报告状态
	ReportDatetime		string `json:"report_datetime"` // 12、报告状态

	Sms					*Sms `orm:"rel(one)"` //`orm:"reverse(one)"`   // 2、发短信主体     一条消息 对应 多个明细

	//Creator				*BackendUser `orm:"rel(fk)"` //设置一对多关系
}



func SmsDetailBatchInsert(smsDetail []SmsDetail) (int64) { //, id *int

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


	mobilelist := strings.Split(smsDetail[0].Mobile, ",")
	fmt.Println(mobilelist)

	execstring := "INSERT INTO rms_smsdetail (id, mobile, number, corporator, location, submit_state, submit_datetime, send_state, send_datetime, report_state, report_datetime, sms_id) VALUES" //creator_id
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

		for j := i * 10; j < i*10+len(smsDetail); j++ {
			if j < i*10+(len(smsDetail)-1) {
				//id := strconv.Itoa(j)
				//onedata := "(" + id + ", '0', '20180103002930'), "

				onedata := "('"+ strconv.Itoa(smsDetail[0].Id) +
					"', '"+ mobilelist[j - i*10] +
					"', '"+ strconv.Itoa(smsDetail[0].Number) +
					"', '"+ smsDetail[0].Corporator +
					"', '"+ smsDetail[0].Location +
					"', '"+ strconv.Itoa(smsDetail[0].SubmitState) +
					"', '"+ smsDetail[0].SubmitDatetime +
					"', '"+ strconv.Itoa(smsDetail[0].SendState) +
					"', '"+ smsDetail[0].SendDatetime +
					"', '"+ strconv.Itoa(smsDetail[0].ReportState) +
					"', '"+ smsDetail[0].ReportDatetime +

					"', '"+ strconv.Itoa(smsDetail[0].Sms.Id) +
					"'), "
				data = data + onedata
			} else {
				//id := strconv.Itoa(j)
				//onedata := "(" + id + ",'0', '20180103002930')"
				//onedata := "('"+ messageState[j - i*10].Msgid +"', '1', '1', '1', '1')"

				onedata := "('"+ strconv.Itoa(smsDetail[0].Id) +
					"', '"+ mobilelist[j - i*10] +
					"', '"+ strconv.Itoa(smsDetail[0].Number) +
					"', '"+ smsDetail[0].Corporator +
					"', '"+ smsDetail[0].Location +
					"', '"+ strconv.Itoa(smsDetail[0].SubmitState) +
					"', '"+ smsDetail[0].SubmitDatetime +
					"', '"+ strconv.Itoa(smsDetail[0].SendState) +
					"', '"+ smsDetail[0].SendDatetime +
					"', '"+ strconv.Itoa(smsDetail[0].ReportState) +
					"', '"+ smsDetail[0].ReportDatetime +

					"', '"+ strconv.Itoa(smsDetail[0].Sms.Id) +
					"') "
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


type SmsDetailState struct {
	Msgid    string `json:"msgid"`
	Mobile   string `json:"mobile"`
	State    string    `json:"state"`
	Datetime string `json:"datetime"`
}

func SmsDetailStateBatchUpdate(smsDetailState []SmsDetailState) (bool) {

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

	ids := ""
	execstring := "UPDATE rms_smsdetail SET "
	execstring += "report_state = CASE id"
	for _, value := range smsDetailState {
		execstring += " WHEN " + strconv.Itoa(value.Msgid) + " THEN " + strconv.Itoa(value.State) + " "
		ids += (strconv.Itoa(value.Msgid) + ",")
	}
	execstring += " END, "
	execstring += "report_datetime = CASE id"
	for _, value := range smsDetailState {
		execstring += " WHEN " + strconv.Itoa(value.Msgid) + " THEN " + value.Datetime + " " //time.Now().Format("2006-01-02 15:04:05")
		ids += (strconv.Itoa(value.Msgid) + ",")
	}
	execstring += " END "
	execstring += " WHERE id IN (" + ids[0 : len(ids)-1] +")"

	fmt.Println(execstring)

	_, err = db.Exec(execstring)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}




// CoursePageList 获取分页数据
func SmsDetailAllListPageList(params *SmsDetailQueryParam) ([]*SmsDetail, int64) {
	query := orm.NewOrm().QueryTable(SmsDetailTBName())
	data := make([]*SmsDetail, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
		//case "Seq":
		//	sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	// .Filter("Creator", &id)
	query = query.Filter("submit_datetime__istartswith", time.Now().Format("2006-01-02")).RelatedSel() //"name__istartswith", params.NameLike
	total, _ := query.Count()
	fmt.Println(params.Limit, params.Offset)
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}
