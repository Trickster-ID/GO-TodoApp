package repository

import (
	"Todoapp/model/entity"

	"gorm.io/gorm"
)

type Tdlrepo interface {
	SelectAll()([]entity.Todolist, error)
	SelectById(id string)(entity.Todolist, error)
	SelectByCreator(creator int)([]entity.Todolist, error)
	SelectDeletedDataByCreator(creator int)([]entity.Todolist, error)
	Save(tdl entity.Todolist)(entity.Todolist, error)
	Delete(tdl entity.Todolist)(entity.Todolist, error)
}

type tdlcon struct{
	con *gorm.DB
}

func NewTdlrepo(DB *gorm.DB) Tdlrepo{
	return &tdlcon{
		con: DB,
	}
}

var restdl []entity.Todolist

func(db *tdlcon) SelectAll()([]entity.Todolist, error){
	err := db.con.Find(&restdl).Error
	return restdl, err
}

func(db *tdlcon) SelectById(id string)(entity.Todolist, error){
	rtdl := entity.Todolist{}
	err := db.con.Find(&rtdl, id).Error
	return rtdl, err
}

func(db *tdlcon) SelectByCreator(creator int)([]entity.Todolist, error){
	err := db.con.Where("creator = ?", creator).Find(&restdl).Error
	return restdl, err
}

func(db *tdlcon) SelectDeletedDataByCreator(creator int)([]entity.Todolist, error){
	err := db.con.Unscoped().Where("deleted is not null and creator = ?", creator).Find(&restdl).Error
	return restdl, err
}

func(db *tdlcon) Save(tdl entity.Todolist)(entity.Todolist, error){
	err := db.con.Save(&tdl).Error
	return tdl, err
}

func(db *tdlcon) Delete(tdl entity.Todolist)(entity.Todolist, error){
	err := db.con.Delete(&tdl).Error
	return tdl, err
}