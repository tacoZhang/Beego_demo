package dao

import (
	"github.com/jinzhu/gorm"
	"my_blog/log"
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
		log.DaoLogger.Error(err.Error())
		return nil
	}
	return &dept
}

func (d *DeptDao) SelectAll() *[]models.Dept {
	var list []models.Dept
	err := d.engine.Find(&list).Error
	if err != nil {
		log.DaoLogger.Error(err.Error())
		return nil
	}
	return &list
}

func (d *DeptDao) InsertDept(dept models.Dept) bool {
	thing := d.engine.Begin()
	err := thing.Create(&dept).Error
	if err != nil {
		thing.Rollback()
		log.DaoLogger.Error(err.Error())
		return false
	}
	thing.Commit()
	return true
}

func (d *DeptDao) UpdateDept(dept models.Dept) bool {
	thing := d.engine.Begin()
	err := thing.Save(&dept).Error
	if err != nil {
		thing.Rollback()
		log.DaoLogger.Error(err.Error())
		return false
	}
	thing.Commit()
	return true
}

func (d *DeptDao) DeleteDept(deptno int) bool {
	thing := d.engine.Begin()
	var dept models.Dept
	err := thing.Where("deptno=?", deptno).Delete(&dept).Error
	if err != nil {
		thing.Rollback()
		log.DaoLogger.Error(err.Error())
		return false
	}
	thing.Commit()
	return true
}

//根据deptno查询emp列表(一对多)
func (d *DeptDao) SelectByDeptnoGetEmps(deptno int) *models.Dept {
	dept := models.Dept{Deptno: deptno}
	err := d.engine.First(&dept).Model(&dept.Emps).Find(&dept.Emps).Error
	if err != nil {
		log.DaoLogger.Error(err.Error())
		return nil
	}
	return &dept
}
