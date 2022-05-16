package service

import (
	"Todoapp/helper"
	"Todoapp/model/entity"
	"Todoapp/repository"
	"time"
)

type Tdlsvc interface {
	SelectAll() ([]entity.Todolist, error)
	SelectById(id string) (entity.Todolist, error)
	SelectByCreator(creator int) ([]entity.Todolist, error)
	SelectDeletedDataByCreator(creator int)([]entity.Todolist, error)
	Create(tdlent entity.Todolist) (entity.Todolist, error)
	Update(tdlent entity.Todolist, id string) (entity.Todolist, error)
	Delete(id string) (entity.Todolist, error)
}

type tdlsvc struct{
	tdlrepo repository.Tdlrepo
}

func NewTdlsvc(Tdlrepo repository.Tdlrepo) Tdlsvc{
	return &tdlsvc{
		tdlrepo: Tdlrepo,
	}
}

func(s *tdlsvc) SelectAll() ([]entity.Todolist, error){
	res, err := s.tdlrepo.SelectAll()
	return res, err
}

func(s *tdlsvc) SelectById(id string) (entity.Todolist, error){
	res, err := s.tdlrepo.SelectById(id)
	return res, err
}

func(s *tdlsvc) SelectByCreator(creator int) ([]entity.Todolist, error){
	res, err := s.tdlrepo.SelectByCreator(creator)
	return res, err
}

func(s *tdlsvc) SelectDeletedDataByCreator(creator int)([]entity.Todolist, error){
	res, err := s.tdlrepo.SelectDeletedDataByCreator(creator)
	return res, err
}

func(s *tdlsvc) Create(tdlent entity.Todolist) (entity.Todolist, error){
	res, err := s.tdlrepo.Save(tdlent)
	return res, err
}

func(s *tdlsvc) Update(tdlent entity.Todolist, id string) (entity.Todolist, error){
	rg, errget := s.tdlrepo.SelectById(id)
	if errget != nil{
		return rg, errget
	}
	rg.Task = helper.Ifelse(tdlent.Task, rg.Task).(string)
	rg.DueDate = helper.Ifelse(tdlent.DueDate, rg.DueDate).(time.Time)
	rg.Isdone = helper.Ifelse(tdlent.Isdone, rg.Isdone).(bool)
	res, err := s.tdlrepo.Save(rg)
	return res, err
}

func(s *tdlsvc) Delete(id string) (entity.Todolist, error){
	rg, errg := s.tdlrepo.SelectById(id)
	if errg != nil{
		return rg, errg
	}
	res, err := s.tdlrepo.Delete(rg)
	return res, err
}