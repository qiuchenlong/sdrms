package controllers

import (
	"github.com/lhtzbj12/sdrms/models"
	"encoding/json"
	"fmt"
)

type ApiController struct {
	BaseController
}


func (c *ApiController) Prepare() {
	//先执行
	//c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}



// 接收回调通知
func (c *ApiController) PostMessageState() {

	var smsDetailState[]models.SmsDetailState

	data := c.Ctx.Input.RequestBody
	//json数据封装到MessageState对象中
	err := json.Unmarshal(data, &smsDetailState)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}


	fmt.Println(data, string(data))

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