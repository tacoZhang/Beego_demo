package controllers

import "my_blog/service"

type DeptController struct {
	baseController
}

// @Title 获取所有部门
// @Success 200 {object} models.dept
// @Failure 404 Not Found
// @router /all [get]
func (this *DeptController) GetAll() {
	var out DataResponse
	defer func() {
		this.Data["json"] = out
		this.ServeJSON()
	}()

	ds := service.NewDeptService()
	list := ds.SelectAll()
	if list == nil {
		out.Code = -1
		out.Msg = "Dept查询失败！"
		return
	} else {
		out.Code = 200
		out.Msg = "查询成功"
		out.Data = list
		return
	}
}
