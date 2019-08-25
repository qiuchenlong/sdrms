package controllers

import (
	"encoding/json"
	"github.com/lhtzbj12/sdrms/models"
	"github.com/lhtzbj12/sdrms/enums"
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
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
	c.checkLogin()
}

func (c *DepositController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "deposit/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "deposit/index_footerjs.html"
}



func (c *DepositController) DataGrid() {
	//id := c.curUser.Id

	var params models.DepositQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params) //
	//获取数据列表和总数
	//data, total := models.DepositListPageList(&params, &id)
	data, total := models.DepositAllListPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *DepositController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Add()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &models.Deposit{}
	//var err error
	//if Id > 0 {
	//	//m, err = models.DepositOne(1)
	//	m, err = models.DepositUserOne(Id)
	//
	//	if err != nil {
	//		fmt.Println(err)
	////		c.pageError("数据无效，请刷新后重试")
	//	}
	//}
	m.BackendUser, _ = models.BackendUserOne(Id)
	fmt.Println("========================", m, Id)
	c.Data["m"] = m
	c.setTpl("deposit/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "deposit/edit_footerjs.html"
}

func (c *DepositController) Add() {
	m := models.Deposit{}
	mf := models.DepositFrom{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&mf); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}


	userId := mf.BackendUserId

	m.BackendUser, _ = models.BackendUserOne(userId)

	m.Money = mf.Money
	m.Created = time.Now()
	m.Type = mf.Type
	m.Remarks = mf.Remarks

	fmt.Println("========================", m, userId)



	// 查找关联的用户通道数据
	oM := models.ChannelBackendUserRel{}
	errs := o.QueryTable(models.ChannelBackendUserRelTBName()).Filter("backenduser__id", userId).One(&oM)
	if errs != nil {
		c.jsonResult(enums.JRCodeFailed, "充值失败", m.Id)
	}


	oM.Hit = 1
	if _, err := o.Update(&oM); err != nil {
		c.jsonResult(enums.JRCodeFailed, "充值失败", m.Id)
	}


	oM.Hit = 0
	m.BeforeMoney = oM.Balance
	m.AfterMoney = oM.Balance + m.Money

	// 对用户通道余额+?
	oM.Balance = oM.Balance + m.Money
	fmt.Println("========================", oM)
	if _, err := o.Update(&oM); err != nil {
		c.jsonResult(enums.JRCodeFailed, "充值失败", m.Id)
	}


	m.Channel, _ = models.ChannelOne(oM.Channel.Id)

	// 写入充值记录中
	if _, err := o.Insert(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "充值失败", m.Id)
	}

	c.jsonResult(enums.JRCodeSucc, "充值成功", 1)
	
}