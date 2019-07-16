package models

import (
	"time"
)

type Emp struct {
	//gorm.Model
	Empno    int `gorm:"primary_key"` //设置主键
	Ename    string
	Job      string
	Mgr      int
	Hiredate time.Time
	Sal      float64
	Comm     float64
	Deptno   int
}

func (Emp) TableName() string {
	return "emp"
}
