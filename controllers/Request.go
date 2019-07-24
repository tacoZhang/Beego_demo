package controllers

/*Emp主键*/
type EmpPid struct {
	Empno string //员工编号
}

type AddEmp struct {
	Ename    string `json:"ename"`
	Job      string `json:"job"`
	Mgr      string `json:"mgr"`
	Hiredate string `json:"hiredate"`
	Sal      string `json:"sal"`
	Comm     string `json:"comm"`
	Deptno   string `json:"deptno"`
}
