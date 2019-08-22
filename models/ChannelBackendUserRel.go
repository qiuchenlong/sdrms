package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)


// 用户通道
// ChannelBackendUserRel 通道与用户关系
type ChannelBackendUserRel struct {
	Id          		int
	Channel        		*Channel        `orm:"rel(fk)"`  //外键
	BackendUser 		*BackendUser 	`orm:"rel(fk)" ` // 外键
	Price				float32 			`json:"price"`   // 价格
	Balance				float32 			`json:"balance"`   // 余额
	State       		int 			`json:"state"` // 状态
	Created     		time.Time    	`orm:"auto_now_add;type(datetime)"`
}


type ChannelBackendUserFrom struct {
	Id          		int
	Channel        		int        `json:"channel"`  //外键
	BackendUser 		int 		`json:"backenduser"` // 外键
	Price				float32 			`json:"price"`   // 价格
	Balance				float32 			`json:"balance"`   // 余额
	State       		int 			`json:"state"` // 状态
	Created     		time.Time    	`orm:"auto_now_add;type(datetime)"`
}


// TableName 设置表名
func (a *ChannelBackendUserRel) TableName() string {
	return ChannelBackendUserRelTBName()
}


// CourseQueryParam 用于搜索的类
type ChannelBackendUserQueryParam struct {
	BaseQueryParam
	//NameLike string
}


// CoursePageList 获取分页数据
func ChannelBackendUserListPageList(params *ChannelBackendUserQueryParam) ([]*ChannelBackendUserRel, int64) {
	query := orm.NewOrm().QueryTable(ChannelBackendUserRelTBName())
	data := make([]*ChannelBackendUserRel, 0)
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


// CourseOne 获取单条
func ChannelBackendUserOne(id int) (*ChannelBackendUserRel, error) {
	o := orm.NewOrm()
	m := ChannelBackendUserRel{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
