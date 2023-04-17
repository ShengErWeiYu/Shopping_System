package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
)

func SearchGoodInLike(MyLikeGood *model.SearchUserLike) (interface{}, error) {
	if err := mysql.CheckUid2(MyLikeGood.Uid); err != nil {
		return nil, err
	}
	if err := mysql.CheckPassword(MyLikeGood.Password, MyLikeGood.Uid); err != nil {
		return nil, err
	}
	return mysql.SearchGoodInLike(MyLikeGood.Uid)
}
func LikeAGood(LikeGood *model.UserLike) error {
	if err := mysql.CheckUid2(LikeGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(LikeGood.Password, LikeGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckGid(LikeGood.Gid); err != nil {
		return err
	}
	return mysql.LikeAGood(LikeGood.Gid, LikeGood.Uid)
}
