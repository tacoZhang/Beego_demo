package controllers

import (
	"encoding/json"
	"my_blog/log"
	"my_blog/models"
	"my_blog/service"
	"strconv"
)

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
		if out.Data == nil {
			out.Data = ""
		}
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

// @Title 根据编号获取部门
// @router /getOne [post]
func (this *DeptController) GetOne() {
	var data DeptPid
	var out DataResponse
	defer func() {
		if out.Data == nil {
			out.Data = ""
		}
		this.Data["json"] = out
		this.ServeJSON()
	}()

	reqByte := this.Ctx.Input.RequestBody
	err := json.Unmarshal(reqByte, &data)
	if err != nil {
		log.ControllerLogger.Error("Json解析错误", err.Error())
		out.Code = -1
		out.Msg = "Json解析错误"
		return
	}

	ds := service.NewDeptService()
	deptno, _ := strconv.Atoi(data.Deptno)
	dept := ds.SelectByDeptnoGetEmps(deptno)
	if dept == nil {
		log.ControllerLogger.Error("查询Dept失败")
		out.Code = -1
		out.Msg = "查询Dept失败"
		return
	}
	out.Code = 200
	out.Msg = "查询成功"
	out.Data = dept
	return
}

// @Title 添加部门
// @router /insert [post]
func (this *DeptController) Insert() {
	var data AddDept
	var out DataResponse
	defer func() {
		if out.Data == nil {
			out.Data = ""
		}
		this.Data["json"] = out
		this.ServeJSON()
	}()

	reqByte := this.Ctx.Input.RequestBody
	err := json.Unmarshal(reqByte, &data)
	if err != nil {
		log.ControllerLogger.Error("Json解析错误", err.Error())
		out.Code = -1
		out.Msg = "Json解析错误"
		return
	}

	ds := service.NewDeptService()
	dept := models.Dept{
		Dname: data.Dname,
		Loc:   data.Loc,
	}
	res := ds.InsertDept(dept)
	if !res {
		log.ControllerLogger.Error("Dept添加失败")
		out.Code = -1
		out.Msg = "Dept添加失败"
		return
	}
	out.Code = 200
	out.Msg = "添加成功！"
	return
}

// @Title 删除部门
// @router /del [post]
func (this *DeptController) DelDept() {
	var data DeptPid
	var out DataResponse
	defer func() {
		if out.Data == nil {
			out.Data = ""
		}
		this.Data["json"] = out
		this.ServeJSON()
	}()

	reqByte := this.Ctx.Input.RequestBody
	err := json.Unmarshal(reqByte, &data)
	if err != nil {
		log.ControllerLogger.Error("Json解析错误", err.Error())
		out.Code = -1
		out.Msg = "Json解析错误"
		return
	}

	ds := service.NewDeptService()
	deptno, _ := strconv.Atoi(data.Deptno)
	res := ds.DeleteDept(deptno)
	if !res {
		log.ControllerLogger.Error("删除失败")
		out.Code = -1
		out.Msg = "删除失败"
		return
	}
	out.Code = 200
	out.Msg = "删除成功"
	return
}
