package controller

import (
	"Shopping_System/dao/mysql"
	g "Shopping_System/global"
	"Shopping_System/model"
	"Shopping_System/services"
	"Shopping_System/untils"
	"Shopping_System/untils/http"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetGoodDetail(c *gin.Context) {
	GoodDetail := new(model.GetGoodDetail)
	if err := c.ShouldBind(GoodDetail); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, GoodDetail),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if gooddetail, err := services.GetGoodDetail(GoodDetail); err != nil {
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"您查询的Gid为" + strconv.FormatInt(GoodDetail.Gid, 10) + "的商品详情为": gooddetail,
		})
	}
}
func GetAllGoods(c *gin.Context) {
	var get []model.GetAllGoods
	sqlstr := "select gid,goodname,price,introduction,size,uid from goods where gid>0;"
	if err := g.Xdb.Select(&get, sqlstr); err != nil {
		panic(err)
	} else {
		c.JSON(200, gin.H{
			"以下是您的查询结果": get,
		})
	}
}
func ChangeTheSizeOfGood(c *gin.Context) {
	ChangeSizeGood := new(model.ChangeSizeGood)
	if err := c.ShouldBind(ChangeSizeGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangeSizeGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if err := services.ChangeTheSizeOfGood(ChangeSizeGood); err != nil {
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
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品规格成功!",
		})
	}
}
func ChangeTheIntroductionOfGood(c *gin.Context) {
	ChangeIntroductionGood := new(model.ChangeIntroductionGood)
	if err := c.ShouldBind(ChangeIntroductionGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangeIntroductionGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if err := services.ChangeTheIntroductionOfGood(ChangeIntroductionGood); err != nil {
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
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品简介成功!",
		})
	}
}
func ChangeThePriceOfGood(c *gin.Context) {
	ChangePriceGood := new(model.ChangePriceGood)
	if err := c.ShouldBind(ChangePriceGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangePriceGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if err := services.ChangeThePriceOfGood(ChangePriceGood); err != nil {
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
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品价格成功!",
		})
	}
}
func ChangeTheNameOfGood(c *gin.Context) {
	ChangeNameGood := new(model.ChangeNameGood)
	if err := c.ShouldBind(ChangeNameGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, ChangeNameGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if err := services.ChangeTheNameOfGood(ChangeNameGood); err != nil {
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
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "修改商品名称成功!",
		})
	}
}
func DeleteAGood(c *gin.Context) {
	DeleteGood := new(model.GoodsDelete)
	if err := c.ShouldBind(DeleteGood); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, DeleteGood),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if err := services.DeleteAGood(DeleteGood); err != nil {
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
		if errors.Is(err, mysql.ErrorGoodExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
		if errors.Is(err, mysql.ErrorUserOperation) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message": "删除商品成功!",
		})
	}
}
func SearchMyGood(c *gin.Context) {
	Search := new(model.SearchMyGoods)
	if err := c.ShouldBind(Search); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, Search),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if goodlist, err := services.SearchMyGoods(Search); err != nil {
		if errors.Is(err, mysql.ErrorUserNotExist) {
			c.JSON(200, gin.H{
				"ErrorMessage": err.Error(),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"message":         "这是您在售的商品",
			"GoodInformation": goodlist,
		})
	}
}
func PublishAGood(c *gin.Context) {
	Good := new(model.GoodsPublish)
	if err := c.ShouldBind(Good); err != nil {
		http.RespFailed(c, http.CodeFail)
		c.JSON(200, gin.H{
			"ErrorMessage": untils.GetValidMsg(err, Good),
		})
		return
	} else {
		http.RespSuccess(c, nil)
	}
	if err := services.PublishAGood(Good); err != nil {
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
	} else {
		c.JSON(200, gin.H{
			"message":         "发布商品成功!",
			"GoodInformation": Good,
		})
	}
}
