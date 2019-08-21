package models

import "time"

// ChannelBackendUserRel 通道与用户关系
type ChannelBackendUserRel struct {
	Id          		int
	Channel        		*Channel        `orm:"rel(fk)"`  //外键
	BackendUser 		*BackendUser 	`orm:"rel(fk)" ` // 外键
	Price				string 			`json:"price"`   // 价格
	Balance				string 			`json:"balance"`   // 余额
	State       		int 			`json:"state"` // 状态
	Created     		time.Time    	`orm:"auto_now_add;type(datetime)"`
}

// TableName 设置表名
func (a *ChannelBackendUserRel) TableName() string {
	return ChannelBackendUserRelTBName()
}
