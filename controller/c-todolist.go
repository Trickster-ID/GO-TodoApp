package controller

import (
	"Todoapp/helper"
	"Todoapp/model/entity"
	"Todoapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tdlctrl interface {
	GetAll(cx *gin.Context)
	GetByCreator(cx *gin.Context)
	GetDeletedDataByCreator(cx *gin.Context)
	Post(cx *gin.Context)
	Put(cx *gin.Context)
	Delete(cx *gin.Context)
}

type tdlctrl struct{tdlsvc service.Tdlsvc}

func NewTdlctrl(Tdlsvc service.Tdlsvc) Tdlctrl{
	return &tdlctrl{
		tdlsvc: Tdlsvc,
	}
}

func(c *tdlctrl) GetAll(cx *gin.Context){
	res, err :=c.tdlsvc.SelectAll()
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to get all data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to get data", res))
}

func(c *tdlctrl) GetByCreator(cx *gin.Context){
	creator := helper.Getjwtdata(cx.GetHeader("Authorization")).(int)
	res, err :=c.tdlsvc.SelectByCreator(creator)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to get all your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to get data", res))
}

func(c *tdlctrl) GetDeletedDataByCreator(cx *gin.Context){
	creator := helper.Getjwtdata(cx.GetHeader("Authorization")).(int)
	res, err :=c.tdlsvc.SelectDeletedDataByCreator(creator)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to get all your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to get data", res))
}

func(c *tdlctrl) Post(cx *gin.Context){
	var tdlent entity.Todolist
	errSb := cx.ShouldBind(&tdlent)
	if errSb != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to binding body data", errSb.Error(), helper.EmptyObject{}))
		return
	}
	tdlent.Creator = helper.Getjwtdata(cx.GetHeader("Authorization")).(int)
	res, err :=c.tdlsvc.Create(tdlent)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to save your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to input data", res))
}

func(c *tdlctrl) Put(cx *gin.Context){
	tdlent := entity.Todolist{}
	id := cx.Param("id")
	errSb := cx.ShouldBind(&tdlent)
	if errSb != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to binding body data", errSb.Error(), helper.EmptyObject{}))
		return
	}
	res, err :=c.tdlsvc.Update(tdlent, id)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to save your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to update data", res))
}

func(c *tdlctrl) Delete(cx *gin.Context){
	id := cx.Param("id")
	res, err := c.tdlsvc.Delete(id)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to delete your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to delete data", res))
}