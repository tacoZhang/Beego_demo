package dao

import (
	"github.com/jinzhu/gorm"
	"my_blog/log"
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
		log.DaoLogger.Error(err.Error())
		return nil
	}
	return &emp
}

//查询所有
func (e *EmpDao) SelectAll() *[]models.Emp {
	var list []models.Emp
	err := e.engine.Find(&list).Error
	if err != nil {
		log.DaoLogger.Error(err.Error())
		return nil
	}
	return &list
}

//添加员工
func (e *EmpDao) InsertEmp(emp models.Emp) bool {
	thing := e.engine.Begin()
	err := thing.Create(&emp).Error
	if err != nil {
		thing.Rollback()
		log.DaoLogger.Error(err.Error())
		return false
	}
	thing.Commit()
	return true
}

//修改员工信息
func (e *EmpDao) UpdateEmp(emp models.Emp) bool {
	thing := e.engine.Begin()
	err := thing.Save(&emp).Error //自动识别主键修改
	if err != nil {
		thing.Rollback()
		log.DaoLogger.Error(err.Error())
		return false
	}
	thing.Commit()
	return true
}

//删除员工
func (e *EmpDao) DeleteEmp(empno int) bool {
	var emp models.Emp
	thing := e.engine.Begin()
	err := thing.Where("empno=?", empno).Delete(&emp).Error
	if err != nil {
		thing.Rollback()
		log.DaoLogger.Error(err.Error())
		return false
	}
	thing.Commit()
	return true
}

//根据empno查询dept (一对一)
func (d *EmpDao) SelectByEmpNoGetDept(empno int) *models.Emp {
	data := models.Emp{Empno: empno}
	err := d.engine.First(&data).Model(&data.Dept).First(&data.Dept).Error
	if err != nil {
		log.DaoLogger.Error(err.Error())
		return nil
	}
	return &data
}
