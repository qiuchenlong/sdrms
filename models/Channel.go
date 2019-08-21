package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func (a *Channel) TableName() string {
	return ChannelTBName()
}

// CourseQueryParam 用于搜索的类
type ChannelQueryParam struct {
	BaseQueryParam
	//NameLike string
}

type Channel struct {
	Id					int
	Code    			string `json:"code"`    // 通道编码
	Name	 			string `json:"name"`    // 通道名称
	Price				string `json:"price"`   // 默认价格
	//Creator				*BackendUser `orm:"rel(fk)"` //设置一对多关系
	//Msg					*Sms `orm:"rel(fk)"` //设置一对多关系
}


// CoursePageList 获取分页数据
func ChannelListPageList(params *ChannelQueryParam, id *int) ([]*Channel, int64) {
	query := orm.NewOrm().QueryTable(ChannelTBName())
	data := make([]*Channel, 0)
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
	//query = query.Filter("Creator", &id).RelatedSel()//"name__istartswith", params.NameLike
	total, _ := query.Count()
	fmt.Println(params.Limit, params.Offset)
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}