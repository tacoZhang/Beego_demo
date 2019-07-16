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
