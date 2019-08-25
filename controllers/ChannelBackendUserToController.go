package controllers

import (
	"github.com/lhtzbj12/sdrms/models"
	"github.com/astaxie/beego/orm"
	"github.com/lhtzbj12/sdrms/enums"
	"encoding/json"
	"fmt"
	"github.com/lhtzbj12/sdrms/utils"
)

type ChannelBackendUserToController struct {
	BaseController
}


//type MessageResult struct {
//	Code    	int `json:"code"`
//	Msgid 		int `json:"msgid"`
//	Message     string `json:"message"`
//}

//Prepare 参考beego官方文档说明
func (c *ChannelBackendUserToController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	//c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	c.checkLogin()
}


func (c *ChannelBackendUserToController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "channelbackenduserto/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "channelbackenduserto/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("ChannelBackendUserController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("ChannelBackendUserController", "Delete")

}


// DataGrid 课程管理首页 表格获取数据
func (c *ChannelBackendUserToController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.ChannelBackendUserQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.ChannelBackendUserListPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}


//Edit 添加、编辑课程界面
func (c *ChannelBackendUserToController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt(":id", 0)
	m := models.ChannelBackendUserRel{Id: Id}
	//var m []*models.ChannelBackendUserRel
	if Id > 0 {
		//o := orm.NewOrm()
		//err := o.Read(&m)
		//if err != nil {
		//	c.pageError("数据无效，请刷新后重试")
		//}

		query := orm.NewOrm().QueryTable(models.ChannelBackendUserRelTBName())
		query = query.RelatedSel()
		query.Filter("id__in", Id).One(&m)

		//var sql string
		////联查多张表
		//sql = fmt.Sprintf(`SELECT DISTINCT *
		//FROM %s AS T0
		//INNER JOIN %s AS T1 ON T0.backend_user_id = T1.id
		//INNER JOIN %s AS T2 ON T2.id = T0.channel_id
		//WHERE T0.id = ? Order By T2.seq asc,T2.id asc`, models.ChannelBackendUserRelTBName(), models.BackendUserTBName(), models.ChannelTBName())
		//o.Raw(sql, Id).QueryRows(&m)
	}
	//c.Data["hasImg"] = len(m.Img) > 0


	c.Data["m"] = m


	// 获取通道
	//bp := models.BaseQueryParam{Order: "desc", Sort: "Id", Limit:10, Offset:0}
	//param := models.ChannelQueryParam{BaseQueryParam: bp}
	//channels, _ := models.ChannelListPageList(&param)
	var channels []*models.Channel
	query := orm.NewOrm().QueryTable(models.ChannelTBName())
	query.All(&channels)//.Exclude("id__in", Id)

	c.Data["c"] = channels


	c.setTpl("channelbackenduserto/edit.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "channelbackenduserto/edit_headcssjs.html"
	c.LayoutSections["footerjs"] = "channelbackenduserto/edit_footerjs.html"

	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("ChannelBackendUserToController.Index")
}

func (c* ChannelBackendUserToController) Save() {
	var err error
	m := models.ChannelBackendUserFrom{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "提交表单数据失败，可能原因："+err.Error(), m.Id)
	}

	o := orm.NewOrm()
	if m.Id == 0 {
		if _, err = o.Insert(&m); err == nil {
			c.jsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因："+err.Error(), m.Id)
		}

	} else {
		oM, _ := models.ChannelBackendUserOne(m.Id)
		fmt.Println(m)
		oM.Price = m.Price
		oM.Balance = m.Balance
		oM.Channel, _ = models.ChannelOne(m.Channel)


		if _, err = o.Update(oM); err == nil {
			c.jsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			utils.LogDebug(err)
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}

}


////Delete 批量删除
//func (c *ChannelBackendUserToController) Delete() {
//	strs := c.GetString("ids")
//	ids := make([]int, 0, len(strs))
//	for _, str := range strings.Split(strs, ",") {
//		if id, err := strconv.Atoi(str); err == nil {
//			ids = append(ids, id)
//		}
//	}
//	if num, err := models.ChannelBatchDelete(ids); err == nil {
//		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
//	} else {
//		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
//	}
//}
