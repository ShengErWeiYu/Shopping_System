package controller

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
	"Shopping_System/services"
	"Shopping_System/untils"
	"Shopping_System/untils/http"
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckOrder(c *gin.Context) {
	Order := new(model.CheckOrder)
	if err := c.ShouldBind(Order); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, Order),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if orderlist, err := services.CheckOrder(Order); err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorPassword) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorOrderExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"您的订单如下": orderlist,
		})
	}
}
