package controllers

import (
	"encoding/json"

	"github.com/lflxp/neo4j-api/models"
	"github.com/lflxp/neo4j-api/pkg"
)

// 后端系统认证
type Neo4jController struct {
	BaseController
}

// @Title Search
// @Description 登陆验证
// @Param	body		body 	models.Neo4jPost	true		"查询语句"
// @Success 200 {string} models.Neo4jPost.Cql
// @Failure 403 body is empty
// @router /search [post]
func (u *Neo4jController) Post() {
	// {"code":20000,"data":{"token":"admin"}}
	var info models.Neo4jPost
	json.Unmarshal(u.Ctx.Input.RequestBody, &info)
	data, err := pkg.ReadTran(info.Cql, nil)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = data
	}

	u.ServeJSON()
}
