package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/lhtzbj12/sdrms/enums"
	"github.com/lhtzbj12/sdrms/models"
	"github.com/lhtzbj12/sdrms/utils"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/httplib"
)

type BackendUserController struct {
	BaseController
}

func (c *BackendUserController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	c.checkLogin()

}
func (c *BackendUserController) Index() {
	//是否显示更多查询条件的按钮弃用，前端自动判断
	//c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	//页面模板设置
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backenduser/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("BackendUserController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("BackendUserController", "Delete")
}
func (c *BackendUserController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.BackendUserQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.BackendUserPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

// Edit 添加 编辑 页面
func (c *BackendUserController) Edit() {
	//如果是Post请求，则由Save处理
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := &models.BackendUser{}
	var err error
	if Id > 0 {
		m, err = models.BackendUserOne(Id)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
		o := orm.NewOrm()
		o.LoadRelated(m, "RoleBackendUserRel")
	} else {
		//添加用户时默认状态为启用
		m.Status = enums.Enabled
	}
	c.Data["m"] = m
	//获取关联的roleId列表
	var roleIds []string
	for _, item := range m.RoleBackendUserRel {
		roleIds = append(roleIds, strconv.Itoa(item.Role.Id))
	}
	c.Data["roles"] = strings.Join(roleIds, ",")
	c.setTpl("backenduser/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "backenduser/edit_footerjs.html"
}
func (c *BackendUserController) Save() {
	m := models.BackendUser{}
	o := orm.NewOrm()
	var err error
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	//删除已关联的历史数据
	if _, err := o.QueryTable(models.RoleBackendUserRelTBName()).Filter("backenduser__id", m.Id).Delete(); err != nil {
		c.jsonResult(enums.JRCodeFailed, "删除历史关系失败", "")
	}
	if m.Id == 0 {
		//对密码进行加密
		m.UserPwd = utils.String2md5(m.UserPwd)
		if _, err := o.Insert(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
		}
	} else {
		if oM, err := models.BackendUserOne(m.Id); err != nil {
			c.jsonResult(enums.JRCodeFailed, "数据无效，请刷新后重试", m.Id)
		} else {
			m.UserPwd = strings.TrimSpace(m.UserPwd)
			if len(m.UserPwd) == 0 {
				//如果密码为空则不修改
				m.UserPwd = oM.UserPwd
			} else {
				m.UserPwd = utils.String2md5(m.UserPwd)
			}
			//本页面不修改头像和密码，直接将值附给新m
			m.Avatar = oM.Avatar
		}
		if _, err := o.Update(&m); err != nil {
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}
	//添加关系
	var relations []models.RoleBackendUserRel
	for _, roleId := range m.RoleIds {
		r := models.Role{Id: roleId}
		relation := models.RoleBackendUserRel{BackendUser: &m, Role: &r}
		relations = append(relations, relation)
	}
	if len(relations) > 0 {
		//批量添加
		if _, err := o.InsertMulti(len(relations), relations); err == nil {
			c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "保存失败", m.Id)
		}
	} else {
		c.jsonResult(enums.JRCodeSucc, "保存成功", m.Id)
	}
}
func (c *BackendUserController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	query := orm.NewOrm().QueryTable(models.BackendUserTBName())
	if num, err := query.Filter("id__in", ids).Delete(); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}


func (c *BackendUserController) LoginLog() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backenduser/loginlog_index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("BackendUserController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("BackendUserController", "Delete")
}

func (c *BackendUserController) LoginLogDataGrid() {
	id := c.curUser.Id

	var params models.LoginLogQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params) //
	fmt.Println(c.Ctx.Input.RequestBody)
	fmt.Println(params)
	//获取数据列表和总数
	data, total := models.LoginLogPageList(&params, &id)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

func (c* BackendUserController) ModifyPassword() {
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "backenduser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "backenduser/modifypassword_index_footerjs.html"
}



type ModifyPasswordResult struct {
	Oldpwd    	string `json:"oldpwd"`
	Newpwd 		string `json:"newpwd"`
}

type PasswordResult struct {
	Code    	int `json:"code"`
	Message     string `json:"message"`
}

func (c* BackendUserController) PostModifyPassword() {
	var modifyPasswordResult ModifyPasswordResult

	data := c.Ctx.Input.RequestBody
	//json数据封装到MessageState对象中
	err := json.Unmarshal(data, &modifyPasswordResult)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	oldPwd := modifyPasswordResult.Oldpwd
	newPwd := modifyPasswordResult.Newpwd

	oldPwd = utils.String2md5(oldPwd)
	newPwd = utils.String2md5(newPwd)

	oldPwd = strings.ToUpper(oldPwd)
	newPwd = strings.ToUpper(newPwd)

	req := httplib.Post("http://47.244.240.84:88/api/sms/updatePassword")
	req.Param("username", c.curUser.UserName)
	req.Param("userpwd", oldPwd)
	req.Param("newpwd", newPwd)

	fmt.Println(c.curUser.UserName, oldPwd, newPwd)


	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--->" + str)


	var result PasswordResult
	if err := json.Unmarshal([]byte(str), &result); err == nil {
		//c.Data["balance"] = balance["balance"]
		if result.Code == 0 {
			// 修改成功

			m := &models.BackendUser{}
			//var err error
			if c.curUser.UserName != "" {
				m, err = models.BackendUserOneByUsername(c.curUser.UserName)
				if err != nil {
					c.pageError("数据无效，请刷新后重试")
				}
				m.UserPwd = strings.ToLower(newPwd)

				o := orm.NewOrm()
				if _, err := o.Update(m); err != nil {
					c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
				}
			}

		}
	} else {

	}



	c.Ctx.WriteString(str)

	//responsedata, _ := json.Marshal(modifyPasswordResult)
	//c.Ctx.WriteString(string(responsedata))
}