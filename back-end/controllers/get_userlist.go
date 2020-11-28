package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

//GetUserListControllers 获取用户列表
type GetUserListControllers struct {
	beego.Controller
}

//returnUserListJSON 返回的数据结构
type returnUserListJSON struct {
	Code  int64    `json:"code"`
	Data  []string `json:"data"`
	Error string   `json:"error"`
}

//Post 获取用户列表
func (p *GetUserListControllers) Post() {
	var returnData returnUserListJSON
	var PostData PostPubData
	defer p.Ctx.Request.Body.Close()
	//var err error
	err := json.NewDecoder(p.Ctx.Request.Body).Decode(&PostData)
	if err != nil {
		beego.Error(err, PostData)
		returnData.Code = 500
		returnData.Error = err.Error()
		p.Data["json"] = returnData
		p.Ctx.Output.JSON(returnData, true, true)
		return
	}
	returnData.Data = append(returnData.Data, "fuck", "fuck01")
	returnData.Code = 200
	returnData.Error = ""
	p.Data["json"] = returnData
	p.Ctx.Output.JSON(returnData, true, true)
}
