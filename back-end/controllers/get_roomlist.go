package controllers

import (
	"github.com/astaxie/beego"
)

//GetRoomListControllers 获取房间列表
type GetRoomListControllers struct {
	beego.Controller
}

//returnRoomListJSON 返回的数据结构
type returnRoomListJSON struct {
	Code  int64    `json:"code"`
	Data  []string `json:"data"`
	Error string   `json:"error"`
}

//Post 获取房间列表
func (p *GetRoomListControllers) Post() {
	var returnData returnRoomListJSON
	returnData.Data = append(returnData.Data, "shit", "shit01")
	returnData.Code = 200
	returnData.Error = ""
	p.Data["json"] = returnData
	p.Ctx.Output.JSON(returnData, true, true)
}
