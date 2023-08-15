package controllers

import (
	"net/http"
	"project2/pkg/models"
	interfaces "project2/pkg/service/iservice"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MarketController struct {
	service interfaces.MarketServiceInterface
}

func NewMarketController(service interfaces.MarketServiceInterface) *MarketController {
	return &MarketController{service: service}
}

func (cont *MarketController) CreateProduct(ctx *gin.Context) {
	var product models.ProductInput
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "invalid input",
			},
		)
		return
	}

	id, err := cont.service.CreateProduct(&product)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"created_product_id": id,
		},
	)
}

func (cont *MarketController) GetAllProducts(ctx *gin.Context) {
	products, err := cont.service.GetAllProducts()

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		products,
	)

}

func (cont *MarketController) GetAllProductsWithName(ctx *gin.Context) {

	name := ctx.Param("name")

	products, err := cont.service.GetProductByName(name)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		products,
	)
}

func (cont *MarketController) GetProductWithId(ctx *gin.Context) {
	s := ctx.Param("id")

	id, err := strconv.Atoi(s)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "id must be unsigned integer",
			},
		)
		return
	}

	product, err := cont.service.GetProductById(uint(id))

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "No product wih this id",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		product,
	)
}

func (cont *MarketController) DeleteProductById(ctx *gin.Context) {
	s := ctx.Param("id")

	id, err := strconv.Atoi(s)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "id must be unsigned integer",
			},
		)
		return
	}

	if err := cont.service.DeleteProductById(uint(id)); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"deleted_product_id": id,
		},
	)
}
