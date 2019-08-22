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
	NameLike string
}


// 通道管理
type Channel struct {
	Id					int
	Code    			string `json:"code"`    // 通道编码
	Name	 			string `json:"name"`    // 通道名称
	Price				float32 `json:"price"`   // 默认价格
	//Creator				*BackendUser `orm:"rel(fk)"` //设置一对多关系
	//Msg					*Sms `orm:"rel(fk)"` //设置一对多关系
}


// CoursePageList 获取分页数据
func ChannelListPageList(params *ChannelQueryParam) ([]*Channel, int64) {
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
	query = query.Filter("name__istartswith", params.NameLike)
	total, _ := query.Count()
	fmt.Println(params.Limit, params.Offset)
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}


// CourseOne 获取单条
func ChannelOne(id int) (*Channel, error) {
	o := orm.NewOrm()
	m := Channel{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}


// CourseBatchDelete 批量删除
func ChannelBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(ChannelTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}