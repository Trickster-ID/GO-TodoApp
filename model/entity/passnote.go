package entity

import (
	"time"

	"gorm.io/gorm"
)

type Passnote struct {
	ID       uint		`gorm:"autoIncrement;primaryKey"`
	Username string		`gorm:"type:varchar;null" json:"username" form:"username"`
	Password string		`gorm:"type:varchar;null" json:"password" form:"password"`
	Note     string		`gorm:"type:varchar;null" json:"note" form:"note"`
	Creator  int		
	CreatedAt time.Time	
	UpdatedAt time.Time	
	Deleted gorm.DeletedAt
}