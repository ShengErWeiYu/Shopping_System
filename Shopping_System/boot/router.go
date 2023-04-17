package boot

//设置路由组，并启动
import (
	"Shopping_System/controller"
	"github.com/gin-gonic/gin"
	"log"
)

// 全部基础功能先写独立的，写完记得商家删除商品也会体现在用户收藏和购物车内
func InitRouters() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/registration", controller.Register)          //Register
		user.POST("/login", controller.Login)                    //Login
		user.POST("/password/forget", controller.ForgetPassword) //ForgetPassword,把密保问题写出来
		user.POST("/username/revise", controller.ReviseUsername) //ReviseUsername
		user.POST("/password/revise", controller.RevisePassword) //RevisePassword
		user.POST("/sex/revise", controller.ReviseSex)           //ReviseSex
		user.POST("/securityquestion/forget", controller.GetSecurityQuestion)
		//user.POST("/deregistration",controller.DisRegister)//DisRegister
		goods := r.Group("/user/goods")
		{
			//买家用户的商品操作
			goods.GET("/get", controller.GetAllGoods)               //GetAllGoods
			goods.POST("/search", controller.GetGoodDetail)         //GetGoodDetail
			goods.POST("/like", controller.LikeAGood)               //LikeAGood
			goods.POST("/like/search", controller.SearchGoodInLike) //SearchGoodInLike
			//goods.POST("/buy")//BuyAGood
			//商家用户的商品操作
			goods.POST("/publish", controller.PublishAGood)      //PublishAGood
			goods.POST("/searchmygood", controller.SearchMyGood) //SearchMyGood
			goods.POST("/delete", controller.DeleteAGood)        //DeleteAGood
			operation := r.Group("/user/goods/operation")
			{
				operation.POST("/goodname", controller.ChangeTheNameOfGood)             //ChangeTheNameOfGood
				operation.POST("/price", controller.ChangeThePriceOfGood)               //ChangeThePriceOfGood
				operation.POST("/introduction", controller.ChangeTheIntroductionOfGood) //ChangeTheIntroductionOfGood
				operation.POST("/size", controller.ChangeTheSizeOfGood)                 //ChangeTheSizeOfGood
			}
		}
		ShoppingCart := r.Group("/user/shoppingcart")
		{
			//用户对购物车的操作
			ShoppingCart.POST("/love", controller.LoveAGood)                  //LoveAGood
			ShoppingCart.POST("/search", controller.SearchGoodInShoppingCart) //SearchGoodInShoppingCart
			ShoppingCart.POST("/clear", controller.ClearShoppingCart)         //ClearShoppingCart
			ShoppingCart.POST("/allbuy", controller.BuyAllInShoppingCart)     //BuyAllInShoppingCart
			ShoppingCart.POST("/buyagood", controller.BuyAGoodInShoppingCart) //
			operation := r.Group("user/shoppingcart/operation")
			{
				operation.POST("/operatenum", controller.OperateNumberOfGood)      //OperateNumberOfGood
				operation.POST("/operatesize", controller.OperateSizeOfGood)       //OperateSizeOfGood
				operation.POST("/deletegood", controller.DeleteGoodInShoppingCart) //DeleteGoodInShoppingCart
			}
		}
		Comment := r.Group("/user/comment")
		{
			Comment.POST("/publish", controller.PublishComment) //PublishComment
			Comment.POST("/look", controller.LookComment)
		}
		user.POST("/checkorder", controller.CheckOrder) //CheckOrder
	}
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Router initialization successful")
}
