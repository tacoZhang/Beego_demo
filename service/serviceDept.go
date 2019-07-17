package service

import (
	"my_blog/Db"
	"my_blog/dao"
	"my_blog/models"
)

type deptService interface {
	SelectByPid(Pid int) *models.Dept
	SelectAll() *[]models.Dept
	InsertDept(dept models.Dept) bool
	UpdateDept(dept models.Dept) bool
	DeleteDept(deptno int) bool
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

func (s *DeptService) InsertDept(dept models.Dept) bool {
	return s.dao.InsertDept(dept)
}

func (s *DeptService) UpdateDept(dept models.Dept) bool {
	return s.dao.UpdateDept(dept)
}

func (s *DeptService) DeleteDept(deptno int) bool {
	return s.dao.DeleteDept(deptno)
}
