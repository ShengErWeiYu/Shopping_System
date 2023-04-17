package model

type User struct {
	Uid              string `form:"uid" json:"uid" db:"uid"  binding:"required" msg:"uid不能为空"`
	Username         string `form:"username" json:"username" db:"username" binding:"required" msg:"需要输入用户名"`
	Password         string `form:"password" json:"password" db:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Sex              int    `form:"sex" json:"sex" db:"sex" binding:"required" msg:"性别填写可选择男(1)、女(0)、保密(2)三种"`
	SecurityQuestion string `form:"securityquestion" db:"securityquestion" json:"security-question" binding:"required" msg:"密保问题不能为空"`
	Answer           string `form:"answer" json:"answer" db:"answer" binding:"required" msg:"密保答案不能为空"`
}

type LoginUser struct {
	Uid      string `form:"uid" json:"uid" db:"uid"  binding:"required" msg:"uid不能为空"`
	Password string `form:"password" json:"password" db:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
}
type GetSecurityQuestion struct {
	Uid string `form:"uid" json:"uid" binding:"required" msg:"要重置密码的账户uid不能为空"`
}
type FgtPswUser struct {
	Uid        string `form:"uid" json:"uid" binding:"required" msg:"uid不能为空"`
	Answer     string `form:"answer" json:"answer" binding:"required" msg:"密保答案不能为空"`
	RePassword string `form:"repassword" json:"repassword" binding:"required" msg:"重置的密码不能为空"`
}

type ReviseNameUser struct {
	Uid         string `form:"uid" json:"uid" binding:"required" msg:"uid不能为空"`
	Password    string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	NewUsername string `form:"newusername" json:"newusername" binding:"required" msg:"新的用户名不能为空"`
}
type ReviseSexUser struct {
	Uid      string `form:"uid" json:"uid" binding:"required" msg:"uid不能为空"`
	Password string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	NewSex   int    `form:"newsex" json:"newsex" binding:"required,gte=0,lte=2" msg:"修改的性别不能为空,且性别填写可选择男(1)、女(0)、保密(2)三种"`
}
type RevisePasswordUser struct {
	Uid         string `form:"uid" json:"uid" binding:"required" msg:"uid不能为空"`
	Password    string `form:"password" json:"password" binding:"min=5,max=20" msg:"密码长度不能小于5且不能大于20"`
	Answer      string `form:"answer" json:"answer" binding:"required" msg:"密保答案不能为空"`
	NewPassword string `form:"newpassword" json:"newpassword" binding:"required" msg:"修改的密码不能为空"`
}
