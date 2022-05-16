package main

import (
	"Todoapp/config"
	"Todoapp/controller"
	"Todoapp/middleware"
	"Todoapp/repository"
	"Todoapp/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db 						*gorm.DB					= config.SetUpDatabaseConnection()
	//todorepo				repository.Tdlrepo			= repository.NewTdlrepo(db)
	passnoterepo			repository.Pnrepo			= repository.NewPnrepo(db)
	todolistrepo			repository.Tdlrepo			= repository.NewTdlrepo(db)

	passnotesvc				service.Pnsvc				= service.NewPnsvc(passnoterepo)
	todolistsvc				service.Tdlsvc				= service.NewTdlsvc(todolistrepo)

	pass					controller.Pnctrl			= controller.NewPnctrl(passnotesvc)
	todo 					controller.Tdlctrl			= controller.NewTdlctrl(todolistsvc)
)

func main(){
	defer config.CloseDatabaseConnection(db)

	r := gin.Default()
	version := r.Group("/v0.1", middleware.AuthJWT())
	{
		pn := version.Group("/passnote")
		{
			pn.GET("/", pass.GetByCreator)
			pn.GET("/getall", pass.GetAll)
			pn.GET("/deleteddata", pass.GetDeletedDataByCreator)
			pn.POST("/", pass.Post)
			pn.PUT("/:id", pass.Put)
			pn.DELETE("/:id", pass.Delete)
		}
		tdl := version.Group("/todolist")
		{
			tdl.GET("/", todo.GetByCreator)
			tdl.GET("/getall", todo.GetAll)
			tdl.GET("/deleteddata", todo.GetDeletedDataByCreator)
			tdl.POST("/", todo.Post)
			tdl.PUT("/:id", todo.Put)
			tdl.DELETE("/:id", todo.Delete)
		}
	}

	r.Run()
}