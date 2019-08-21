package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"database/sql"
	"github.com/astaxie/beego/orm"
	"strings"
	"strconv"
)

func (a *Sms) TableName() string {
	return SmsTBName()
}

// CourseQueryParam 用于搜索的类
type SmsQueryParam struct {
	BaseQueryParam
	//NameLike string
}

type Sms struct {
	Id					int
	Msgid    			string `json:"msgid"`
	Mobile	 			string `json:"mobile"`
	Content				string `json:"content"`
	Sign				string `json:"sign"`
	State    			int `json:"state"`
	Datetime 			string `json:"datetime"`
	Creator				*BackendUser `orm:"rel(fk)"` //设置一对多关系

	//MsgState 			*MessageState `orm:"rel(one)"` // 设置一对多的反向关系

}


// CoursePageList 获取分页数据
func SmsListPageList(params *SmsQueryParam, id *int) ([]*Sms, int64) {
	query := orm.NewOrm().QueryTable(SmsTBName())
	data := make([]*Sms, 0)
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
	query = query.Filter("Creator", &id).RelatedSel()//"name__istartswith", params.NameLike
	total, _ := query.Count()
	fmt.Println(params.Limit, params.Offset)
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}


//func SmsListPageList2(params *SmsQueryParam, id *int) ([]*MessageState, int64) {
//	query := orm.NewOrm().QueryTable(MessageStateTBName())
//	data := make([]*MessageState, 0)
//	//默认排序
//	sortorder := "Id"
//	switch params.Sort {
//	case "Id":
//		sortorder = "Id"
//		//case "Seq":
//		//	sortorder = "Seq"
//	}
//	if params.Order == "desc" {
//		sortorder = "-" + sortorder
//	}
//	query = query.RelatedSel()//"name__istartswith", params.NameLike
//	total, _ := query.Count()
//	fmt.Println(params.Limit, params.Offset)
//	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
//	return data, total
//}



func SmsBatchInsert(sms []Sms, id int) (int64) {

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


	mobilelist := strings.Split(sms[0].Mobile, ",")
	fmt.Println(mobilelist)


	execstring := "INSERT INTO rms_sms (msgid, mobile, content, sign, state, datetime, creator_id)VALUES"
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


		for j := i * 10; j < i*10+len(mobilelist); j++ {
			if j < i*10+(len(mobilelist)-1) {
				//id := strconv.Itoa(j)
				//onedata := "(" + id + ", '0', '20180103002930'), "
				onedata := "('"+ sms[0].Msgid +"', '"+ mobilelist[j - i*10] + "', '"+ sms[0].Content + "', '"+ sms[0].Sign + "', '"+ strconv.Itoa(sms[0].State) +"', '"+ sms[0].Datetime + "', '" + strconv.Itoa(id) + "'), "
				//onedata := "('1', '1', '1', '1', '1'), "
				data = data + onedata
			} else {
				//id := strconv.Itoa(j)
				//onedata := "(" + id + ",'0', '20180103002930')"
				//onedata := "('"+ sms[j - i*10].Msgid +"', '1', '1', '1', '1')"
				onedata := "('"+ sms[0].Msgid +"', '"+ mobilelist[j - i*10] + "', '"+ sms[0].Content + "', '"+ sms[0].Sign + "', '"+ strconv.Itoa(sms[0].State) +"', '"+ sms[0].Datetime + "', '" + strconv.Itoa(id) + "') "
				//onedata := "('1', '1', '1', '1', '1')"
				data = data + onedata
			}
		}


		//for j := i * 10; j < i*10+len(sms); j++ {
		//	if j < i*10+(len(sms)-1) {
		//		//id := strconv.Itoa(j)
		//		//onedata := "(" + id + ", '0', '20180103002930'), "
		//		onedata := "('"+ sms[j - i*10].Msgid +"', '"+ sms[j - i*10].Mobile + "', '"+ sms[j - i*10].Content + "', '"+ sms[j - i*10].Sign + "', '"+ strconv.Itoa(sms[j - i*10].State) +"', '"+ sms[j - i*10].Datetime + "', '" + strconv.Itoa(id) + "'), "
		//		//onedata := "('1', '1', '1', '1', '1'), "
		//		data = data + onedata
		//	} else {
		//		//id := strconv.Itoa(j)
		//		//onedata := "(" + id + ",'0', '20180103002930')"
		//		//onedata := "('"+ sms[j - i*10].Msgid +"', '1', '1', '1', '1')"
		//		onedata := "('"+ sms[j - i*10].Msgid +"', '"+ sms[j - i*10].Mobile + "', '"+ sms[j - i*10].Content + "', '"+ sms[j - i*10].Sign + "', '"+ strconv.Itoa(sms[j - i*10].State) +"', '"+ sms[j - i*10].Datetime + "', '" + strconv.Itoa(id) + "') "
		//		//onedata := "('1', '1', '1', '1', '1')"
		//		data = data + onedata
		//	}
		//}

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