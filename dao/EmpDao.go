package dao

import (
	"github.com/jinzhu/gorm"
	"my_blog/models"
)

type EmpDao struct {
	engine *gorm.DB
}

func NewEmpDao(engine *gorm.DB) *EmpDao {
	return &EmpDao{
		engine: engine,
	}
}

//通过主键查询
func (e *EmpDao) SelectByPId(Pid int) *models.Emp {
	var emp = models.Emp{}
	err := e.engine.First(&emp, Pid).Error
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &emp
}

//查询所有
func (e *EmpDao) SelectAll() *[]models.Emp {
	var list []models.Emp
	err := e.engine.Find(&list).Error
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &list
}

//添加员工
func (e *EmpDao) InsertEmp(emp models.Emp) bool {
	thing := e.engine.Begin()
	defer thing.Close()
	err := thing.Create(&emp).Error
	if err != nil {
		thing.Rollback()
		panic(err.Error())
		return false
	}
	thing.Commit()
	return true
}

//修改员工信息
func (e *EmpDao) UpdateEmp(emp models.Emp) bool {
	thing := e.engine.Begin()
	defer thing.Close()
	err := thing.Save(&emp).Error //自动识别主键修改
	if err != nil {
		thing.Rollback()
		panic(err.Error())
		return false
	}
	thing.Commit()
	return true
}

//删除员工
func (e *EmpDao) DeleteEmp(empno int) bool {
	var emp models.Emp
	thing := e.engine.Begin()
	defer thing.Close()
	err := thing.Where("empno=?", empno).Delete(&emp).Error
	if err != nil {
		thing.Rollback()
		panic(err.Error())
		return false
	}
	thing.Commit()
	return true
}
