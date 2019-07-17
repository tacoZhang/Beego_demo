package service

import (
	"my_blog/Db"
	"my_blog/dao"
	"my_blog/models"
)

type EmpService interface {
	SelectByPId(pid int) *models.Emp
	SelectAll() *[]models.Emp
	InsertEmp(emp models.Emp) bool
	UpdateEmp(emp models.Emp) bool
	DeleteEmp(empno int) bool
}

type empService struct {
	dao *dao.EmpDao
}

func NewEmpService() EmpService {
	return &empService{
		dao: dao.NewEmpDao(Db.MariaConf),
	}
}

func (s *empService) SelectByPId(pid int) *models.Emp {
	return s.dao.SelectByPId(pid)
}

func (s *empService) SelectAll() *[]models.Emp {
	return s.dao.SelectAll()
}

func (s *empService) InsertEmp(emp models.Emp) bool {
	return s.dao.InsertEmp(emp)
}

func (s *empService) UpdateEmp(emp models.Emp) bool {
	return s.dao.UpdateEmp(emp)
}

func (s *empService) DeleteEmp(empno int) bool {
	return s.dao.DeleteEmp(empno)
}
