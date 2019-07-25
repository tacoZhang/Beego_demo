package controllers

/*Emp主键*/
type EmpPid struct {
	Empno string //员工编号
}

/*添加新Emp*/
type AddEmp struct {
	Ename    string `json:"ename"`
	Job      string `json:"job"`
	Mgr      string `json:"mgr"`
	Hiredate string `json:"hiredate"`
	Sal      string `json:"sal"`
	Comm     string `json:"comm"`
	Deptno   string `json:"deptno"`
}

/*更新Emp*/
type UpdateEmp struct {
	Empno  string `json:"empno"`
	Ename  string `json:"ename"`
	Job    string `json:"job"`
	Mgr    string `json:"mgr"`
	Sal    string `json:"sal"`
	Comm   string `json:"comm"`
	Deptno string `json:"deptno"`
}

/*部门主键*/
type DeptPid struct {
	Deptno string `json:"deptno"`
}

/*添加新部门*/
type AddDept struct {
	Dname string `json:"dname"`
	Loc   string `json:"loc"`
}
