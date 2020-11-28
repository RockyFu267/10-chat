package controllers

import (
	"encoding/json"

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
type PostPubData struct {
	ID           string `json:"id,omitempty"`
	ProjectID    string `json:"project_id,omitempty"`
	PipelineName string `json:"pipeline_name,omitempty"`
	UpdateUser   string `json:"update_user,omitempty"`
	CanaryWeight string `json:"canaryweight,omitempty"`
	BaseImage    string `json:"baseimage,omitempty"`
	ClusterName  string `json:"cluster_name,omitempty"`
}

//Post 获取房间列表
func (p *GetRoomListControllers) Post() {
	var returnData returnRoomListJSON
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
	returnData.Data = append(returnData.Data, "shit", "shit01")
	returnData.Code = 200
	returnData.Error = ""
	p.Data["json"] = returnData
	p.Ctx.Output.JSON(returnData, true, true)
}
