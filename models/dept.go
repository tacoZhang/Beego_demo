package models

type Dept struct {
	Deptno int `gorm:"primary_key"`
	Dname  string
	Loc    string
}
