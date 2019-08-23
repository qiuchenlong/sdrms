package controllers

import (
	// "fmt"
	"encoding/json"
	"github.com/lhtzbj12/sdrms/models"
	// "strings"
)

type DepositController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *DepositController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

func (c *DepositController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
}



func (c *DepositController) DepositDataGrid() {
	id := c.curUser.Id

	var params models.DepositQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params) //
	//获取数据列表和总数
	//data, total := models.SmsListPageList(&params, &id)
	data, total := models.DepositListPageList(&params, &id)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}


func (c *DepositController) Save() {
	m := models.Deposit{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	
	
}