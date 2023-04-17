package mysql

import (
	g "Shopping_System/global"
	"Shopping_System/model"
)

const (
	AddLikeStr             = "insert into likes (gid,uid) value (?,?)"
	FindGoodInLikeByUidStr = "select lid,gid from likes where lid>0 and uid=?"
)

func SearchGoodInLike(uid string) (interface{}, error) {
	var MyLikeGoods []model.MyLike
	if err := g.Xdb.Select(&MyLikeGoods, FindGoodInLikeByUidStr, uid); err != nil {
		return nil, err
	}
	return MyLikeGoods, nil
}

func LikeAGood(gid int64, uid string) error {
	if _, err := g.Xdb.Exec(AddLikeStr, gid, uid); err != nil {
		return err
	}
	return nil
}
