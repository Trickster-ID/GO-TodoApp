package controller

import (
	"Todoapp/helper"
	"Todoapp/model/entity"
	"Todoapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pnctrl interface {
	GetAll(cx *gin.Context)
	GetByCreator(cx *gin.Context)
	GetDeletedDataByCreator(cx *gin.Context)
	Post(cx *gin.Context)
	Put(cx *gin.Context)
	Delete(cx *gin.Context)
}

type pnctrl struct{pnsvc service.Pnsvc}

func NewPnctrl(Pnsvc service.Pnsvc) Pnctrl{
	return &pnctrl{
		pnsvc: Pnsvc,
	}
}

func(c *pnctrl) GetAll(cx *gin.Context){
	res, err :=c.pnsvc.SelectAll()
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to get all data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to get data", res))
}

func(c *pnctrl) GetByCreator(cx *gin.Context){
	creator := helper.Getjwtdata(cx.GetHeader("Authorization")).(int)
	res, err :=c.pnsvc.SelectByCreator(creator)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to get all your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to get data", res))
}

func(c *pnctrl) GetDeletedDataByCreator(cx *gin.Context){
	creator := helper.Getjwtdata(cx.GetHeader("Authorization")).(int)
	res, err :=c.pnsvc.SelectDeletedDataByCreator(creator)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to get all your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to get deleted data", res))
}

func(c *pnctrl) Post(cx *gin.Context){
	var psent entity.Passnote
	errSb := cx.ShouldBind(&psent)
	if errSb != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to binding body data", errSb.Error(), helper.EmptyObject{}))
		return
	}
	psent.Creator = helper.Getjwtdata(cx.GetHeader("Authorization")).(int)
	res, err :=c.pnsvc.Create(psent)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to save your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to input data", res))
}

func(c *pnctrl) Put(cx *gin.Context){
	psent := entity.Passnote{}
	id := cx.Param("id")
	errSb := cx.ShouldBind(&psent)
	if errSb != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to binding body data", errSb.Error(), helper.EmptyObject{}))
		return
	}
	res, err := c.pnsvc.Update(psent, id)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to update data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to update data", res))
}

func(c *pnctrl) Delete(cx *gin.Context){
	id := cx.Param("id")
	res, err := c.pnsvc.Delete(id)
	if err != nil{
		cx.JSON(http.StatusBadRequest, helper.BuildErrorResponse("fail to delete your data", err.Error(), res))
		return
	}
	cx.JSON(http.StatusOK, helper.BuildResponse("success to delete data", res))
}