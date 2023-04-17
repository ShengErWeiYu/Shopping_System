package mysql

import "errors"

var (
	ErrorUserExist               = errors.New("用户已存在！")
	ErrorUserNotExist            = errors.New("用户不存在，请先注册！")
	ErrorPassword                = errors.New("密码错误")
	ErrorAnswer                  = errors.New("密保答案错误！")
	ErrorGoodExist               = errors.New("该商品不存在！")
	ErrorUserOperation           = errors.New("您没有本次操作的权限！")
	ErrorSizeExist               = errors.New("该商品的规格没有您要添入购物车的规格！")
	ErrorGoodExistInShoppingCart = errors.New("该用户的购物车暂无商品！")
	ErrorThisGoodInShoppingCart  = errors.New("购物车内无此商品！")
	ErrorCommentPublishUser      = errors.New("非买家不能评论")
	ErrorGoodExistInOrders       = errors.New("订单中无该商品的记录！")
	ErrorOrderExist              = errors.New("您没有订单！")
)
