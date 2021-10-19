package controller

import (
	"net/http"
	"strconv"

	"github.com/anish-gyawali/assessment/server/dto"
	"github.com/anish-gyawali/assessment/server/entity"
	"github.com/anish-gyawali/assessment/server/service"
	"github.com/gin-gonic/gin"
)

//DataController is a ...
type DataController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type dataController struct {
	dataService service.DataService
}

type EmptyObj struct{}

//NewDataController create a new instances of DataController
func NewDataController(dataServ service.DataService) DataController {
	return &dataController{
		dataService: dataServ,
	}
}

func (c *dataController) All(context *gin.Context) {
	var datas []entity.Data = c.dataService.All()
	res := datas
	context.JSON(http.StatusOK, res)
}

func (c *dataController) FindByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	var data entity.Data = c.dataService.FindByID(id)
	if (data == entity.Data{}) {
		context.JSON(http.StatusNotFound, EmptyObj{})
	} else {
		context.JSON(http.StatusOK, data)
	}
}

func (c *dataController) Insert(context *gin.Context) {
	var dataCreateDTO dto.DataCreateDTO
	errDTO := context.ShouldBind(&dataCreateDTO)
	if errDTO != nil {
		context.JSON(http.StatusBadRequest, errDTO.Error())
	} else {
		result := c.dataService.Insert(dataCreateDTO)
		context.JSON(http.StatusCreated, result)
	}
}

func (c *dataController) Update(context *gin.Context) {
	var dataUpdateDTO dto.DataUpdateDTO
	errDTO := context.ShouldBind(&dataUpdateDTO)
	if errDTO != nil {
		context.JSON(http.StatusBadRequest, errDTO.Error())
		return
	} else {
		result := c.dataService.Update(dataUpdateDTO)
		context.JSON(http.StatusOK, result)
	}
}

func (c *dataController) Delete(context *gin.Context) {
	var data entity.Data
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		context.JSON(http.StatusBadRequest, EmptyObj{})
	} else {
		data.ID = id
		c.dataService.Delete(data)
		context.JSON(http.StatusOK, EmptyObj{})
	}
}
