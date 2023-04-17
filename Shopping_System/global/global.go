package global

import (
	"github.com/jmoiron/sqlx"
)

// 定义全局变量
var (
	Xdb *sqlx.DB
)
