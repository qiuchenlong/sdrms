package controllers

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
	"encoding/json"
	"github.com/lhtzbj12/sdrms/models"
	"strings"
)

type SmsController struct {
	BaseController
}


type MessageResult struct {
	Code    	int `json:"code"`
	Msgid 		int `json:"msgid"`
	Message     string `json:"message"`
}

//Prepare 参考beego官方文档说明
func (c *SmsController) Prepare() {
	//先执行
	//c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}


func (c *SmsController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
}


// 发送消息
func (c *SmsController) SendMessage() {
	username := c.curUser.UserName//"CS008"//c.GetString("username")
	userpwd := c.curUser.UserPwd//"123456"//c.GetString("userpwd")

	mobiles := c.GetString("mobiles")
	content := c.GetString("content")
	sign := c.GetString("sign")
	datetime := c.GetString("datetime")

	var timeNow = time.Now().Format("2006-01-02 15:04:05")

	// 如果没有选择定时的时间，将立即发送
	if datetime == "" {
		datetime = timeNow
	}


	mobilelist := strings.Split(mobiles, "\n")
	fmt.Println(mobilelist)

	mobiles = strings.Replace(mobiles, "\n", ",", len(mobiles))


	//pwdmd5 := md5.New()
	//pwdmd5.Write([]byte(userpwd)) // 需要加密的字符串为 sharejs.com

	req := httplib.Post("http://47.244.240.84:88/api/sms/massSend")
	req.Param("username", username)
	req.Param("userpwd", userpwd)
	req.Param("mobiles", mobiles)
	req.Param("content", content)
	req.Param("sign", "")
	req.Param("extid", "")
	req.Param("timing_time", datetime)

	fmt.Println(username, userpwd, mobiles, content, sign, datetime)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--->" + str)


	//str = "{\"code\":0,\"msgid\":118106,\"message\":\"发送成功，本次发送短信 1 条！\"}"


	// 写入消息主体表格
	sms := make([]models.Sms, 1)

	var smsState = 0
	var messageResult MessageResult
	if err := json.Unmarshal([]byte(str), &messageResult); err == nil {
		//c.Data["balance"] = balance["balance"]
		fmt.Println("-->>", messageResult.Code)
		if messageResult.Code == 0 {
			// 发送成功

			smsState = 1//(发送中，完成）

			models.BalanceDec(c.curUser.Id)

		} else {
			// 发送失败
			smsState = 0//(发送中，完成）

		}
	} else {
		fmt.Println("--error-->", err)

		smsState = 0//(发送中，完成）

	}


	sms[0].Mobile = mobiles
	sms[0].Content = content
	sms[0].Sign = sign
	sms[0].State = smsState
	sms[0].Datetime = datetime
	sms[0].Creator = &c.curUser
	sms[0].Channel, _ = models.ChannelOne(1) // 使用默认通道
	sms[0].Price = 0.33 // 使用默认价格

	sms[0].Code = messageResult.Code
	sms[0].Message = messageResult.Message
	sms[0].Msgid = messageResult.Msgid


	if models.SmsBatchInsert(sms) == 1 && messageResult.Code == 0 {

		smsObject, _ := models.SmsOne(messageResult.Msgid) //&sms[0]

		// 写入消息明细表格
		smsDetail := make([]models.SmsDetail, 1)
		smsDetail[0].Id = messageResult.Msgid
		smsDetail[0].Mobile = mobiles
		smsDetail[0].Number = 1
		smsDetail[0].Corporator = ""
		smsDetail[0].Location = ""
		smsDetail[0].SubmitState = smsState
		smsDetail[0].SubmitDatetime = timeNow
		smsDetail[0].SendState = smsState
		smsDetail[0].SendDatetime = datetime
		smsDetail[0].ReportState = 0
		smsDetail[0].ReportDatetime = ""

		fmt.Println("--------->>>>>>", smsObject)
		if (smsObject != nil) {
			smsDetail[0].Sms = smsObject
		} else {
			smsDetail[0].Sms = &sms[0]
		}

		models.SmsDetailBatchInsert(smsDetail)

	}







	// 发送成功
	//sms := make([]models.Sms, 1)
	//
	////var smslist []models.Sms
	//sms[0].Msgid = messageResult.Msgid
	//sms[0].Mobile = mobiles
	//sms[0].Content = content
	//sms[0].Sign = sign
	//sms[0].State = 0//(发送中，完成）
	//sms[0].Datetime = timing_time
	//sms[0].Creator = &c.curUser
	//id := c.curUser.Id
	//models.SmsBatchInsert(sms, id)


	// {"code":0,"msgid":110943,"message":"发送成功，本次发送短信 1 条！"}
	// {"code":0,"msgid":110947,"message":"发送成功，本次发送短信 1 条！"}

	//var str = "{'id':1}"
	c.Ctx.WriteString(str)
}


func (c *SmsController) SmsList() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "sms/sms_index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("BackendUserController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("BackendUserController", "Delete")
}


func (c *SmsController) SmsDataGrid() {
	id := c.curUser.Id

	//var params models.SmsQueryParam
	var params models.SmsDetailQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params) //
	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println(params)
	//获取数据列表和总数
	//data, total := models.SmsListPageList(&params, &id)
	data, total := models.SmsDetailListPageList(&params, &id)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}


func (c *SmsController) SmsDataGridAll() {
	//var params models.SmsQueryParam
	var params models.SmsDetailQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params) //
	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println(params)
	//获取数据列表和总数
	//data, total := models.SmsListPageList(&params, &id)
	data, total := models.SmsDetailAllListPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}




// 接收回调通知
func (c *SmsController) PostMessageState() {

	var smsDetailState[]models.SmsDetailState

	data := c.Ctx.Input.RequestBody
	//json数据封装到MessageState对象中
	err := json.Unmarshal(data, &smsDetailState)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	//MessageState
	//for _,  := range MessageState {

	
	if len(smsDetailState) > 0 {
		fmt.Println("data=",  smsDetailState[0].Msgid)


		fmt.Println("更新是否成功？", models.SmsDetailStateBatchUpdate(smsDetailState))

		//c.Ctx.WriteString("{'id':"+ strconv.Itoa(len(MessageState)) +"}")

		responsedata, _ := json.Marshal(smsDetailState)
		c.Ctx.WriteString(string(responsedata))
	} else {
		c.Ctx.WriteString("{'code': -1}")
	}

}

