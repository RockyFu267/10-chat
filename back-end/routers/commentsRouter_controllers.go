package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	//GetRoomListControllers
	beego.GlobalControllerRouter["Hydrogen/controllers:GetRoomListControllers"] = append(beego.GlobalControllerRouter["Hydrogen/controllers:GetRoomListControllers"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})
	//GetUserListControllers
	beego.GlobalControllerRouter["Hydrogen/controllers:GetUserListControllers"] = append(beego.GlobalControllerRouter["Hydrogen/controllers:GetUserListControllers"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
