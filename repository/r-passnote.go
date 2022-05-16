package repository

import (
	"Todoapp/model/entity"

	"gorm.io/gorm"
)

type Pnrepo interface {
	SelectAll()([]entity.Passnote, error)
	SelectById(id string)(entity.Passnote, error)
	SelectByCreator(creator int)([]entity.Passnote, error)
	SelectDeletedDataByCreator(creator int)([]entity.Passnote, error)
	Save(ps entity.Passnote)(entity.Passnote, error)
	Delete(ps entity.Passnote)(entity.Passnote, error)
}

type pnconn struct{
	con *gorm.DB
}

func NewPnrepo(DB *gorm.DB) Pnrepo{
	return &pnconn{
		con: DB,
	}
}

var passnotes []entity.Passnote

func(db *pnconn) SelectAll()([]entity.Passnote, error){
	err := db.con.Find(&passnotes).Error
	return passnotes, err
}

func(db *pnconn) SelectById(id string)(entity.Passnote, error){
	rpn := entity.Passnote{}
	err := db.con.Find(&rpn, id).Error
	return rpn, err
}

func(db *pnconn) SelectByCreator(creator int)([]entity.Passnote, error){
	err := db.con.Where("creator = ?", creator).Find(&passnotes).Error
	return passnotes, err
}

func(db *pnconn) SelectDeletedDataByCreator(creator int)([]entity.Passnote, error){
	err := db.con.Unscoped().Where("deleted is not null and creator = ?", creator).Find(&passnotes).Error
	return passnotes, err
}

func(db *pnconn) Save(ps entity.Passnote)(entity.Passnote, error){
	err := db.con.Save(&ps).Error
	return ps, err
}

func(db *pnconn) Delete(ps entity.Passnote)(entity.Passnote, error){
	err := db.con.Delete(&ps).Error
	return ps, err
}