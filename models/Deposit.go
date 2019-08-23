package models

import (
    "time"
	"github.com/astaxie/beego/orm"
	"fmt"
)


// 充值中心
type Deposit struct {
	Id			int          // 1
	Money		float32      // 4
	BeforeMoney float32      // 5
	AfterMoney  float32      // 6
	Type        int          // 7
	Created     time.Time  `orm:"auto_now_add;type(datetime)"`      // 8
	Remarks     string       // 9

	BackendUser     *BackendUser `orm:"rel(fk)"`      // 2
	Channel        		*Channel        `orm:"rel(fk)"`  //外键   3
}



// TableName 设置表名
func (a *Deposit) TableName() string {
	return DepositTBName()
}



// CourseQueryParam 用于搜索的类
type DepositQueryParam struct {
	BaseQueryParam
	//NameLike string
}



// CoursePageList 获取分页数据
func DepositListPageList(params *DepositQueryParam, id *int) ([]*Deposit, int64) {
	query := orm.NewOrm().QueryTable(DepositTBName())
	data := make([]*Deposit, 0)
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
	//query = query.Filter("name__istartswith", params.NameLike)
	query = query.Filter("backend_user__in", id).RelatedSel()
	total, _ := query.Count()
	fmt.Println(params.Limit, params.Offset)
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}