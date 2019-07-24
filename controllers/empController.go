package controllers

import (
	"encoding/json"
	"my_blog/log"
	"my_blog/models"
	"my_blog/service"
	"strconv"
	"time"
)

type EmpController struct {
	baseController
}

// @Title 获取所有Emp对象
// @Description Emp列表
// @Success 200 {object} models.emp
// @Failure 404 Not Found
// @router /all [get]
func (this *EmpController) GetAll() {
	var out DataResponse
	es := service.NewEmpService()
	list := es.SelectAll()
	if list == nil { //查询失败
		out = Reponse(-1, list, "查询失败")
	} else {
		out = Reponse(200, list, "Success")
	}
	this.Data["json"] = out
	this.ServeJSON()
}

// @Title 根据Id获取Emp
// @Description 获取单个Emp
// @Sucess 200 {object} models.emp
// @Failut 404 Not Found
// @router /get [post]
func (this *EmpController) GetOne() {
	var data EmpPid
	var out DataResponse
	defer func() {
		this.Data["json"] = out
		this.ServeJSON()
	}()
	reqByte := this.Ctx.Input.RequestBody
	err := json.Unmarshal(reqByte, &data)
	if err != nil {
		log.ControllerLogger.Error("json解析异常", err.Error())
		out.Code = -1
		out.Msg = "json解析异常"
		return
	}

	es := service.NewEmpService()
	empno, _ := strconv.Atoi(data.Empno)
	emp := es.SelectByPId(empno)
	if emp == nil {
		log.ControllerLogger.Error("emp查询失败")
		out.Code = -1
		out.Msg = "Emp查询失败"
		return
	} else {
		out.Code = 200
		out.Msg = "查询成功"
		out.Data = emp
		return
	}
}

// @Title 获取Emp详情包括Dept
// @router /theEmp [post]
func (this *EmpController) GetEmp() {
	var data EmpPid
	var out DataResponse
	defer func() {
		this.Data["json"] = out
		this.ServeJSON()
	}()
	reqByte := this.Ctx.Input.RequestBody
	err := json.Unmarshal(reqByte, &data)
	if err != nil {
		log.ControllerLogger.Error("json解析异常", err.Error())
		out.Code = -1
		out.Msg = "Json解析异常"
		return
	}

	es := service.NewEmpService()
	empno, _ := strconv.Atoi(data.Empno)
	emp := es.SelectByEmpNoGetDept(empno)
	if emp == nil {
		out.Code = -1
		out.Msg = "Emp查询失败"
		return
	} else {
		out.Code = 200
		out.Msg = "查询成功"
		out.Data = emp
		return
	}
}

// @Title 添加Emp
// @router /insert [post]
func (this *EmpController) InsertEmp() {
	var data AddEmp
	var out DataResponse
	defer func() {
		this.Data["json"] = out
		this.ServeJSON()
	}()

	reqByte := this.Ctx.Input.RequestBody
	err := json.Unmarshal(reqByte, &data)
	if err != nil {
		log.ControllerLogger.Error("Json解析异常", err.Error())
		out.Code = -1
		out.Msg = "Json解析异常"
		return
	}

	es := service.NewEmpService()
	mgr, _ := strconv.Atoi(data.Mgr)
	hirdate, _ := strconv.ParseInt(data.Hiredate, 10, 64)
	sal, _ := strconv.ParseFloat(data.Sal, 64)
	comm, _ := strconv.ParseFloat(data.Comm, 64)
	deptno, _ := strconv.Atoi(data.Deptno)
	emp := models.Emp{
		Ename:    data.Ename,
		Job:      data.Job,
		Mgr:      mgr,
		Hiredate: time.Unix(hirdate, 0),
		Sal:      sal,
		Comm:     comm,
		Deptno:   deptno,
	}
	result := es.InsertEmp(emp)
	if result {
		out.Code = 200
		out.Msg = "添加成功"
		return
	} else {
		out.Code = -1
		out.Msg = "添加失败"
		return
	}
}
