package controllers

import (
	"fmt"
	"net/http"

	"github.com/mazeem91/trackman-poc/domain/models"
	"github.com/mazeem91/trackman-poc/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repository *repository.SQLite
}

type AddLocationRequestBody struct {
	Name string `json:"name" binding:"required"`
	Area string `json:"area" binding:"required"`
}

func (h Handler) GetLocations(ctx *gin.Context) {
	var location []*models.Location
	h.Repository.GetWith(&location, "Area")
	ctx.JSON(http.StatusOK, &location)
}

func (h Handler) AddLocation(ctx *gin.Context) {
	body := AddLocationRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var location models.Location
	var area models.Area

	if err := h.Repository.GetBy(&area, &models.Area{Name: body.Area}); err != nil {
		area = models.Area{Name: body.Area}
		if result := h.Repository.Save(&area); result != nil {
			ctx.AbortWithError(http.StatusNotFound, result)
			return
		}
		fmt.Println(location.AreaID)
	}
	location.Area = area

	location.Name = body.Name

	if result := h.Repository.Save(&location); result != nil {
		ctx.AbortWithError(http.StatusNotFound, result)
		return
	}

	ctx.JSON(http.StatusCreated, &location)
}

type AddSkuRequestBody struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

func (h Handler) GetSkus(ctx *gin.Context) {
	var sku []*models.Sku
	h.Repository.Get(&sku)
	ctx.JSON(http.StatusOK, &sku)
}

func (h Handler) AddSku(ctx *gin.Context) {
	body := AddSkuRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var sku models.Sku

	sku.Name = body.Name
	sku.Code = body.Code
	sku.Quantity = body.Quantity

	if result := h.Repository.Save(&sku); result != nil {
		ctx.AbortWithError(http.StatusNotFound, result)
		return
	}

	ctx.JSON(http.StatusCreated, &sku)
}

type AssignSkuLocationRequestUri struct {
	SkuID      uint `uri:"sku_id" binding:"required"`
	LocationID uint `uri:"location_id" binding:"required"`
}

type AssignSkuLocationRequestBody struct {
	Quantity int `json:"quantity" binding:"required"`
}

func (h Handler) AssignSkuLocation(ctx *gin.Context) {
	uri := AssignSkuLocationRequestUri{}
	body := AssignSkuLocationRequestBody{}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var sku models.Sku

	if err := h.Repository.GetBy(&sku, models.Sku{ID: uri.SkuID}); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var location models.Location

	if err := h.Repository.GetBy(&location, models.Location{ID: uri.LocationID}); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var sku_location models.SkuLocation

	sku_location.Sku = sku
	sku_location.Location = location
	sku_location.Quantity = body.Quantity

	if result := h.Repository.Save(&sku_location); result != nil {
		ctx.AbortWithError(http.StatusNotFound, result)
		return
	}

	ctx.JSON(http.StatusCreated, &sku_location)
}
