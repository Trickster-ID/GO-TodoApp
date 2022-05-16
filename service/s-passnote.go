package service

import (
	"Todoapp/helper"
	"Todoapp/model/entity"
	"Todoapp/repository"
)

type Pnsvc interface {
	SelectAll() ([]entity.Passnote, error)
	SelectById(id string) (entity.Passnote, error)
	SelectByCreator(creator int) ([]entity.Passnote, error)
	SelectDeletedDataByCreator(creator int)([]entity.Passnote, error)
	Create(psent entity.Passnote) (entity.Passnote, error)
	Update(psent entity.Passnote, id string) (entity.Passnote, error)
	Delete(id string) (entity.Passnote, error)
}

type pnsvc struct{
	pnrepo repository.Pnrepo
}

func NewPnsvc(Pnrepo repository.Pnrepo) Pnsvc{
	return &pnsvc{
		pnrepo: Pnrepo,
	}
}

func(s *pnsvc) SelectAll() ([]entity.Passnote, error){
	res, err := s.pnrepo.SelectAll()
	for i, s := range res {
		decryptpass, _ := helper.Decrypt(s.Password)
		res[i].Password = decryptpass
	}
	return res, err
}

func(s *pnsvc) SelectById(id string) (entity.Passnote, error){
	res, err := s.pnrepo.SelectById(id)
	decryptpass, _ := helper.Decrypt(res.Password)
	res.Password = decryptpass
	return res, err
}

func(s *pnsvc) SelectByCreator(creator int) ([]entity.Passnote, error){
	res, err := s.pnrepo.SelectByCreator(creator)
	for i, s := range res {
		decryptpass, _ := helper.Decrypt(s.Password)
		res[i].Password = decryptpass
	}
	return res, err
}

func(s *pnsvc) SelectDeletedDataByCreator(creator int)([]entity.Passnote, error){
	res, err := s.pnrepo.SelectDeletedDataByCreator(creator)
	for i, s := range res {
		decryptpass, _ := helper.Decrypt(s.Password)
		res[i].Password = decryptpass
	}
	return res, err
}

func(s *pnsvc) Create(psent entity.Passnote) (entity.Passnote, error){
	encryptpass, _ := helper.Encrypt(psent.Password)
	psent.Password = encryptpass
	res, err := s.pnrepo.Save(psent)
	decryptpass, _ := helper.Decrypt(res.Password)
	res.Password = decryptpass
	return res, err
}

func(s *pnsvc) Update(psent entity.Passnote, id string) (entity.Passnote, error){
	rg, errg := s.pnrepo.SelectById(id)
	if errg != nil{
		return rg, errg
	}
	rg.Username = helper.Ifelse(psent.Username, rg.Username).(string)
	rg.Note = helper.Ifelse(psent.Note, rg.Note).(string)
	if psent.Password != ""{
		encryptpass, _ := helper.Encrypt(psent.Password)
		rg.Password = encryptpass
	}
	res, err := s.pnrepo.Save(rg)
	decryptpass, _ := helper.Decrypt(res.Password)
	res.Password = decryptpass
	return res, err
}

func(s *pnsvc) Delete(id string) (entity.Passnote, error){
	rg, errg := s.pnrepo.SelectById(id)
	if errg != nil{
		return rg, errg
	}
	res, err := s.pnrepo.Delete(rg)
	decryptpass, _ := helper.Decrypt(res.Password)
	res.Password = decryptpass
	return res, err
}