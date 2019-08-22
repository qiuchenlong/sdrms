package routers

import (
	"github.com/lhtzbj12/sdrms/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//课程路由
	//beego.Router("/course/index", &controllers.CourseController{}, "*:Index")
	//beego.Router("/course/datagrid", &controllers.CourseController{}, "Get,Post:DataGrid")
	//beego.Router("/course/edit/?:id", &controllers.CourseController{}, "Get,Post:Edit")
	//beego.Router("/course/delete", &controllers.CourseController{}, "Post:Delete")
	//beego.Router("/course/updateseq", &controllers.CourseController{}, "Post:UpdateSeq")
	//beego.Router("/course/uploadimage", &controllers.CourseController{}, "Post:UploadImage")

	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")
	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")
	beego.Router("/home/datareset", &controllers.HomeController{}, "Post:DataReset")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")


	beego.Router("/sms/index", &controllers.SmsController{}, "*:Index")
	beego.Router("/sms/send", &controllers.SmsController{}, "*:SendMessage")
	beego.Router("/sms/sendlist", &controllers.SmsController{}, "*:SmsList")
	beego.Router("/sms/smsdatagrid", &controllers.SmsController{}, "Post:SmsDataGrid")
	beego.Router("/sms/postmessagestatus", &controllers.SmsController{}, "Post:PostMessageState")

	beego.Router("/sms/queryBalance", &controllers.BalanceController{}, "*:Index")
	beego.Router("/backenduser/loginlog", &controllers.BackendUserController{}, "*:LoginLog")
	beego.Router("/backenduser/loginlogdatagrid", &controllers.BackendUserController{}, "Post:LoginLogDataGrid")
	beego.Router("/backenduser/modifypassword", &controllers.BackendUserController{}, "*:ModifyPassword")
	//beego.Router("/backenduser/postmodifypassword", &controllers.BackendUserController{}, "Post:PostModifyPassword")


	beego.Router("/channel/index", &controllers.ChannelController{}, "*:Index")
	beego.Router("/channel/datagrid", &controllers.ChannelController{}, "Get,Post:DataGrid")
	beego.Router("/channel/edit/?:id", &controllers.ChannelController{}, "Get,Post:Edit")
	beego.Router("/channel/delete", &controllers.ChannelController{}, "Post:Delete")



	beego.Router("/channelbackenduser/index", &controllers.ChannelBackendUserController{}, "*:Index")
	beego.Router("/channelbackenduser/datagrid", &controllers.ChannelBackendUserController{}, "Get,Post:DataGrid")
	beego.Router("/channelbackenduser/edit/?:id", &controllers.ChannelBackendUserController{}, "Get,Post:Edit")
	beego.Router("/channelbackenduser/delete", &controllers.ChannelBackendUserController{}, "Post:Delete")


	beego.Router("/sms/smsdatagridall", &controllers.SmsController{}, "Get:SmsDataGridAll")



	beego.Router("/", &controllers.HomeController{}, "*:Index")

}
