package model

// 这里是订单,是买家下订单的地方，也是卖家查看订单的地方
type Orders struct {
	Uid         string `form:"uid" json:"uid" binding:"required" msg:"uid不能为空"`
	Gid         int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
	Number      int    `form:"number" json:"number" binding:"min=1" msg:"商品数量不能为0"`
	Size        string `form:"size" json:"size" binding:"size" msg:"商品规格不能为空"`
	Address     string `form:"address" json:"address" binding:"required" msg:"收货地址不能为空"`
	PhoneNumber string `form:"phonenumber" json:"phonenumber" binding:"required" msg:"电话号码不能为空"`
	RealName    string `form:"realname" json:"realname" binding:"required" msg:"请填写你的真实姓名"`
}
type CheckOrder struct {
	Uid      string `form:"uid" json:"uid" binding:"required" msg:"uid不能为空"`
	Password string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
}
type CheckOrderResult struct {
	Oid         int
	Uid         string
	Gid         int64
	Number      int
	Size        string
	Address     string
	PhoneNumber string
	RealName    string
}
type Eieiei struct {
	num string
}
