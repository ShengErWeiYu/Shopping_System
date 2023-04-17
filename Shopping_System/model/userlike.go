package model

type UserLike struct {
	Uid      string `form:"uid" json:"uid" binding:"required" msg:"用户id不能为空"`
	Password string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Gid      int64  `form:"gid" json:"gid" binding:"required" msg:"商品编号不能为空"`
}
type SearchUserLike struct {
	Uid      string `form:"uid" json:"uid" binding:"required" msg:"用户id不能为空"`
	Password string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
}
type MyLike struct {
	Lid int
	Gid int
}
