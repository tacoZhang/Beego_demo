package dao

import (
	"github.com/jinzhu/gorm"
	"my_blog/models"
)

type DeptDao struct {
	engine *gorm.DB
}

func NewDeptDao(engine *gorm.DB) *DeptDao {
	return &DeptDao{
		engine: engine,
	}
}

func (d *DeptDao) SelectByPid(Pid int) *models.Dept {
	var dept = models.Dept{}
	err := d.engine.First(&dept, Pid).Error
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &dept
}

func (d *DeptDao) SelectAll() *[]models.Dept {
	var list []models.Dept
	err := d.engine.Find(&list).Error
	if err != nil {
		panic(err.Error)
		return nil
	}
	return &list
}
