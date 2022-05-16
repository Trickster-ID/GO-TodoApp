package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todolist struct {
	ID        int			`gorm:"autoIncrement;primaryKey"`
	Task      string		`gorm:"type:varchar;null" json:"task,omitempty" form:"task"`
	DueDate   time.Time		`gorm:"null" json:"duedate,omitempty" form:"duedate"`
	Creator   int			
	Isdone    bool			`gorm:"type:bool;null" json:"isdone" form:"isdone"`
	CreatedAt time.Time
	UpdatedAt time.Time		
	Deleted gorm.DeletedAt	
}