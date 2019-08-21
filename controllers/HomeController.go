package controllers

import (
	"strings"

	"github.com/lhtzbj12/sdrms/enums"
	"github.com/lhtzbj12/sdrms/models"
	"github.com/lhtzbj12/sdrms/utils"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//判断是否登录
	c.checkLogin()
	c.setTpl()
}
func (c *HomeController) Page404() {
	c.setTpl()
}
func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error.html", "shared/layout_pullbox.html")
}
func (c *HomeController) Login() {

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "home/login_footerjs.html"
	c.setTpl("home/login.html", "shared/layout_base.html")
}
func (c *HomeController) DoLogin() {
	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	if len(username) == 0 || len(userpwd) == 0 {
		c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	userpwd = utils.String2md5(userpwd)
	user, err := models.BackendUserOneByUserName(username, userpwd)
	if user != nil && err == nil {
		if user.Status == enums.Disabled {
			c.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		c.setBackendUser2Session(user.Id)


		o := orm.NewOrm()
		// 登录日志
		m := models.LoginLog{}
		m.Ip = c.BaseController.getClientIp()
		m.Address = ""
		//h, _ := time.ParseDuration("1h")
		//now := time.Now() //.Add(8 * h)
		//local, _ := time.LoadLocation("Asia/Chongqing") //等同于"UTC"
		//m.CreateTime = now.In(local)
		//fmt.Println(now.In(local))
		m.Creator = user
		if _, err = o.Insert(&m); err == nil {
			//c.jsonResult(enums.JRCodeSucc, "添加成功", m.Id)
			fmt.Println(m.Id)
		} else {
			//c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因："+err.Error(), m.Id)
			fmt.Println(err.Error())
		}


		//获取用户信息
		c.jsonResult(enums.JRCodeSucc, "登录成功", "")
	} else {
		c.jsonResult(enums.JRCodeFailed, "用户名或者密码错误", "")
	}
}
func (c *HomeController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.pageLogin()
}
func (c *HomeController) DataReset() {
	if ok, err := models.DataReset(); ok {
		c.jsonResult(enums.JRCodeSucc, "初始化成功", "")
	} else {
		c.jsonResult(enums.JRCodeFailed, "初始化失败,可能原因:"+err.Error(), "")
	}

}
