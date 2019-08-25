package models

import (
    "time"
	"github.com/astaxie/beego/orm"
	"fmt"
)


// 充值记录
type Deposit struct {
	Id			int          // 1  主键
	Money		float32      // 4
	BeforeMoney float32      // 5
	AfterMoney  float32      // 6
	Type        int          // 7
	Created     time.Time  `orm:"auto_now_add;type(datetime)"`      // 8
	Remarks     string       // 9

	BackendUser     *BackendUser `orm:"rel(fk)"`      // 2
	Channel        		*Channel        `orm:"rel(fk)"`  //外键   3
}



type DepositFrom struct {
	Id			int          // 1  主键
	Money		float32      // 4
	BeforeMoney float32      // 5
	AfterMoney  float32      // 6
	Type        int          // 7
	Remarks     string       // 9

	BackendUserId int
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



// CoursePageList 获取分页数据
func DepositAllListPageList(params *DepositQueryParam) ([]*Deposit, int64) {
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
	query = query.RelatedSel()
	total, _ := query.Count()
	fmt.Println(params.Limit, params.Offset)
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}



// BackendUserOne 根据id获取单条
func DepositUserOne(userid int) (*Deposit, error) {
	m := Deposit{}
	err := orm.NewOrm().QueryTable(DepositTBName()).Filter("backend_user_id", userid).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// BackendUserOne 根据id获取单条
func DepositOne(id int) (*Deposit, error) {
	o := orm.NewOrm()
	m := Deposit{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}