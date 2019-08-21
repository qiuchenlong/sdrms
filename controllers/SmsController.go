package controllers

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"time"
	"encoding/json"
	"github.com/lhtzbj12/sdrms/models"
	"strings"
	"strconv"
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
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	c.checkLogin()
}


func (c *SmsController) Index() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
}

func (c *SmsController) SendMessage() {
	username := c.curUser.UserName//"CS008"//c.GetString("username")
	userpwd := c.curUser.UserPwd//"123456"//c.GetString("userpwd")

	mobiles := c.GetString("mobiles")
	content := c.GetString("content")
	sign := c.GetString("sign")
	datetime := c.GetString("datetime")

	var timing_time = time.Now().Format("2006-01-02 15:04:05")

	if datetime != "" {
		timing_time = datetime
	}


	mobilelist := strings.Split(mobiles, "\n")
	fmt.Println(mobilelist)

	mobiles = strings.Replace(mobiles, "\n", ",", len(mobiles))


	//pwdmd5 := md5.New()
	//pwdmd5.Write([]byte(userpwd)) // 需要加密的字符串为 sharejs.com

	req := httplib.Post("http://47.244.240.84:88/api/sms/massSend")
	//req.Param("username", username)
	//req.Param("userpwd", userpwd)
	//req.Param("mobiles", mobiles)
	//req.Param("content", content)
	//req.Param("sign", "")
	//req.Param("extid", "")
	//req.Param("timing_time", "")

	fmt.Println(username, userpwd, mobiles, content, sign, timing_time)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--->" + str)


	str = "{\"code\":0,\"msgid\":118106,\"message\":\"发送成功，本次发送短信 1 条！\"}"

	var messageResult MessageResult
	if err := json.Unmarshal([]byte(str), &messageResult); err == nil {
		//c.Data["balance"] = balance["balance"]
		fmt.Println("-->>", messageResult.Code)
		if messageResult.Code == 0 {
			// 发送成功
			sms := make([]models.Sms, 1)

			//var smslist []models.Sms
			sms[0].Msgid = strconv.Itoa(messageResult.Msgid)
			sms[0].Mobile = mobiles
			sms[0].Content = content
			sms[0].Sign = sign
			sms[0].State = 1//(发送中，完成）
			sms[0].Datetime = timing_time
			sms[0].Creator = &c.curUser
			id := c.curUser.Id
			models.SmsBatchInsert(sms, id)
		} else {
			// 发送失败
			sms := make([]models.Sms, 1)

			//var smslist []models.Sms
			sms[0].Msgid = strconv.Itoa(messageResult.Msgid)
			sms[0].Mobile = mobiles
			sms[0].Content = content
			sms[0].Sign = sign
			sms[0].State = 0//(发送中，完成）
			sms[0].Datetime = timing_time
			sms[0].Creator = &c.curUser
			id := c.curUser.Id
			models.SmsBatchInsert(sms, id)
		}
	} else {
		fmt.Println("--error-->", err)

		// 发送失败
		sms := make([]models.Sms, 1)

		//var smslist []models.Sms
		sms[0].Msgid = strconv.Itoa(messageResult.Msgid)
		sms[0].Mobile = mobiles
		sms[0].Content = content
		sms[0].Sign = sign
		sms[0].State = 0//(发送中，完成）
		sms[0].Datetime = timing_time
		sms[0].Creator = &c.curUser
		id := c.curUser.Id
		models.SmsBatchInsert(sms, id)
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

	var params models.SmsQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params) //
	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println(params)
	//获取数据列表和总数
	data, total := models.SmsListPageList(&params, &id)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}





func (c *SmsController) PostMessageState() {

	var MessageState []models.MessageState

	data := c.Ctx.Input.RequestBody
	//json数据封装到MessageState对象中
	err := json.Unmarshal(data, &MessageState)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	//MessageState
	//for _,  := range MessageState {


	fmt.Println("data=",  MessageState[0].Msgid)
	fmt.Println("", len(MessageState))


	models.BatchInsert(MessageState)

	//c.Ctx.WriteString("{'id':"+ strconv.Itoa(len(MessageState)) +"}")

	responsedata, _ := json.Marshal(MessageState)
	c.Ctx.WriteString(string(responsedata))

}

