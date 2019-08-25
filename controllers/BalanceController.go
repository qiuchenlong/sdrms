package controllers

import (
	"github.com/lhtzbj12/sdrms/models"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
)

type BalanceController struct {
	BaseController
}

type Balance struct {
	Code    	int `json:"code"`
	Balance 	int `json:"balance"`
	Message     string `json:"message"`
}


func (c *BalanceController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的少数Action需要权限控制，则将验证放到需要控制的Action里
	//c.checkAuthor("TreeGrid", "UserMenuTree", "ParentTreeGrid", "Select")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//这里注释了权限控制，因此这里需要登录验证
	c.checkLogin()
}


func (c *BalanceController) IndexAll() {
	//判断是否登录
	//c.checkLogin()
	//c.setTpl()
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	username := c.curUser.UserName//"CS008"//c.GetString("username")
	userpwd := c.curUser.UserPwd//"123456"//c.GetString("userpwd")

	fmt.Println(username + "," + userpwd)

	//pwdmd5 := md5.New()
	//pwdmd5.Write([]byte(userpwd)) // 需要加密的字符串为 sharejs.com

	req := httplib.Get("http://47.244.240.84:88/api/sms/queryBalance")
	req.Param("username", username)
	req.Param("userpwd", userpwd)//hex.EncodeToString(pwdmd5.Sum(nil)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--->" + str)

	//c.Ctx.WriteString(str)

	//var balance map[string] interface{}
	var balance Balance
	if err := json.Unmarshal([]byte(str), &balance); err == nil {
		//c.Data["balance"] = balance["balance"]
		c.Data["balance"] = balance.Balance
		if balance.Balance < 0 {
			c.Data["balance"] = 0
		}
	} else {
		c.Data["balance"] = 0
	}

	c.setTpl()

}


func (c *BalanceController) Index() {
	//判断是否登录
	//c.checkLogin()
	//c.setTpl()
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)


	balance, _ := models.BalanceOne(c.curUser.Id)
	if balance != nil {
		c.Data["balance"] = balance.Balance
	} else {
		c.Data["balance"] = 0
	}


	c.setTpl()

}
