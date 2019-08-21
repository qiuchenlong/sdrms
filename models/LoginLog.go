package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func (a *LoginLog) TableName() string {
	return LoginLogTBName()
}


// CourseQueryParam 用于搜索的类
type LoginLogQueryParam struct {
	BaseQueryParam
	//NameLike string
}

type LoginLog struct {
	Id					int
	Ip      			string `orm:"size(32)"`
	Address 			string `orm:"size(32)"`
	CreateTime 			time.Time `orm:"auto_now_add;type(datetime)"`
	Creator				*BackendUser `orm:"rel(fk)"` //设置一对多关系
}


// CoursePageList 获取分页数据
func LoginLogPageList(params *LoginLogQueryParam, id *int) ([]*LoginLog, int64) {
	query := orm.NewOrm().QueryTable(LoginLogTBName())
	data := make([]*LoginLog, 0)
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