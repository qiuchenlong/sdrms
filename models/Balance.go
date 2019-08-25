package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
	"database/sql"
)


// BalanceOne 获取单条
func BalanceOne(userid int) (*ChannelBackendUserRel, error) {
	//o := orm.NewOrm()
	//m := Sms{Msgid: msgid}
	//err := o.Read(&m)
	//if err != nil {
	//	return nil, err
	//}
	//return &m, nil

	m := ChannelBackendUserRel{}
	err := orm.NewOrm().QueryTable(ChannelBackendUserRelTBName()).Filter("backend_user_id__in", userid).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}


// 余额 被 通道价格减
func BalanceDec(userid int, count int) (bool) {
	//o := orm.NewOrm()
	//m := Sms{Msgid: msgid}
	//err := o.Read(&m)
	//if err != nil {
	//	return nil, err
	//}
	//return &m, nil

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

	execstring := "UPDATE rms_channel_backenduser_rel SET balance = ( balance - price * " + strconv.Itoa(count) + ") WHERE hit = 0 and backend_user_id = " + strconv.Itoa(userid)
	result, err := db.Exec(execstring)
	affectCount, _ := result.RowsAffected()
	fmt.Println("==========", affectCount)
	if err != nil || affectCount == 0 {
		fmt.Println(err)
		return false
	}

	//channelBackendUser, _ := ChannelBackendUserByUserOne(userid)
	//if (channelBackendUser.Price == )


	return true





	//m := ChannelBackendUserRel{}
	//err := orm.NewOrm().QueryTable(ChannelBackendUserRelTBName()).Filter("backend_user_id__in", userid).One(&m)
	//if err != nil {
	//	return nil, err
	//}
	//return &m, nil
}

// 没成功 余额 加回来
func BalanceAdd(userid int, count int) (bool) {
	//o := orm.NewOrm()
	//m := Sms{Msgid: msgid}
	//err := o.Read(&m)
	//if err != nil {
	//	return nil, err
	//}
	//return &m, nil

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

	execstring := "UPDATE rms_channel_backenduser_rel SET balance = ( balance + price * " + strconv.Itoa(count) + ") WHERE backend_user_id = " + strconv.Itoa(userid)
	_, err = db.Exec(execstring)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

	//m := ChannelBackendUserRel{}
	//err := orm.NewOrm().QueryTable(ChannelBackendUserRelTBName()).Filter("backend_user_id__in", userid).One(&m)
	//if err != nil {
	//	return nil, err
	//}
	//return &m, nil
}