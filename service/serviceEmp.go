package service

import (
	"my_blog/Db"
	"my_blog/dao"
	"my_blog/models"
)

type EmpService interface {
	SelectByPId(pid int) *models.Emp
	SelectAll() *[]models.Emp
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
