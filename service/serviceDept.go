package service

import (
	"my_blog/Db"
	"my_blog/dao"
	"my_blog/models"
)

type deptService interface {
	SelectByPid(Pid int) *models.Dept
	SelectAll() *[]models.Dept
}

type DeptService struct {
	dao *dao.DeptDao
}

func NewDeptService() *DeptService {
	return &DeptService{
		dao: dao.NewDeptDao(Db.MariaConf),
	}
}

func (s *DeptService) SelectByPid(Pid int) *models.Dept {
	return s.dao.SelectByPid(Pid)
}

func (s *DeptService) SelectAll() *[]models.Dept {
	return s.dao.SelectAll()
}
