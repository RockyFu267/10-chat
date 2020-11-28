package routers

import (
	"back-end/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	chatNS := beego.NewNamespace("/chat/",
		//GetRoomListControllers 获取房间列表
		beego.NSNamespace("/getroomlist",
			beego.NSInclude(
				&controllers.GetRoomListControllers{},
			),
		),
		//GetUserListControllers 获取用户列表
		beego.NSNamespace("/getuserlist",
			beego.NSInclude(
				&controllers.GetUserListControllers{},
			),
		),
	)
	beego.AddNamespace(chatNS)
}
