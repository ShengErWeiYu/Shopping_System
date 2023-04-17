package model

type GoodsPublish struct {
	Uid          string `form:"uid" json:"uid" binding:"required" msg:"卖家id不能为空"`
	Password     string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	GoodName     string `form:"goodname" json:"goodname" binding:"required" msg:"商品名称不能为空"`
	Price        string `form:"price" json:"price" binding:"required" msg:"价格不能为空"`
	Introduction string `form:"introduction" json:"introduction" binding:"required" msg:"简介不能为空"`
	Size         string `form:"size" json:"size" binding:"required" msg:"商品规格不能为空"`
}
type GetAllGoods struct {
	Gid          int64
	GoodName     string
	Price        string
	Introduction string
	Size         string
	Uid          string
}
type GetGoodDetail struct {
	Gid int64 `form:"gid" json:"gid" binding:"required" msg:"要查询的商品编号不能为空"`
}

//	type GoodsOperation struct {
//		Gid             int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
//		NewGoodName     string `form:"goodname" json:"goodname" `
//		NewPrice        string `form:"price" json:"price" `
//		NewIntroduction string `form:"introduction" json:"introduction" `
//		NewSize         string `form:"size" json:"size" `
//	}
type ChangeSizeGood struct {
	Uid      string `form:"uid" json:"uid" binding:"required" msg:"卖家id不能为空"`
	Password string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Gid      int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
	NewSize  string `form:"newsize" json:"newsize" binding:"required" msg:"修改的商品规格不能为空"`
}
type ChangeIntroductionGood struct {
	Uid             string `form:"uid" json:"uid" binding:"required" msg:"卖家id不能为空"`
	Password        string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Gid             int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
	NewIntroduction string `form:"newintroduction" json:"newintroduction" binding:"required" msg:"修改的商品简介不能为空"`
}
type ChangePriceGood struct {
	Uid      string `form:"uid" json:"uid" binding:"required" msg:"卖家id不能为空"`
	Password string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Gid      int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
	NewPrice string `form:"newprice" json:"newprice" binding:"required" msg:"修改的价格不能为空"`
}
type GoodsDelete struct {
	Uid      string `form:"uid" json:"uid" binding:"required" msg:"卖家id不能为空"`
	Password string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Gid      int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
}
type ChangeNameGood struct {
	Uid         string `form:"uid" json:"uid" binding:"required" msg:"卖家id不能为空"`
	Password    string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Gid         int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
	NewGoodName string `form:"newgoodname" json:"newgoodname" binding:"required" msg:"商品新名称不能为空"`
}
type SearchMyGoods struct {
	Uid string `form:"uid" json:"uid" binding:"required" msg:"用户id不能为空"`
}
type SearchResult struct {
	Gid      int64
	GoodName string
}
