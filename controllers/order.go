package controllers

import (
	"assignment2/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func (oc *OrderController) GetAll(c *gin.Context) {
	orders := []models.Order{}

	if err := oc.DB.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "unsuccessfully get all orders",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successfully get all orders",
		"data":    orders,
	})
}

func (oc *OrderController) Create(c *gin.Context) {
	newOrder := models.Order{}

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(400, gin.H{
			"message": "unsuccessfully create new order",
		})
		return
	}

	if err := oc.DB.Create(&newOrder).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "unsuccessfully create new order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "successfully create new order",
		"order_id": newOrder.ID,
	})
}

func (oc *OrderController) Update(c *gin.Context) {
	id := c.Param("orderId")
	updatedOrder := models.Order{}
	order := models.Order{}

	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(400, gin.H{
			"message": "unsuccessfully update order",
		})
		return
	}

	if err := oc.DB.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("unsuccessfully update order, order with id = %s is not found", id),
		})
		return
	}

	order.CustomerName = updatedOrder.CustomerName
	order.OrderedAt = updatedOrder.OrderedAt
	order.Items = updatedOrder.Items

	if err := oc.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "unsuccessfully update order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "successfully update order",
		"order_id": order.ID,
	})
}

func (oc *OrderController) Delete(c *gin.Context) {
	id := c.Param("orderId")

	order := models.Order{}

	if err := oc.DB.First(&order, id).Error; err != nil {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("unsuccessfully delete order, order with id = %s is not found", id),
		})
		return
	}

	if err := oc.DB.Select("Item").Delete(&order).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "unsuccessfully delete order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "successfully delete order",
		"order_id": id,
	})
}
