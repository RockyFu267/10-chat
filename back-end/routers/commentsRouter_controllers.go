package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	//GetRoomListControllers
	beego.GlobalControllerRouter["back-end/controllers:GetRoomListControllers"] = append(beego.GlobalControllerRouter["back-end/controllers:GetRoomListControllers"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})
	//GetUserListControllers
	beego.GlobalControllerRouter["back-end/controllers:GetUserListControllers"] = append(beego.GlobalControllerRouter["back-end/controllers:GetUserListControllers"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
