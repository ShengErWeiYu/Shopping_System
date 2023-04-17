package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
)

func OperateSizeOfGood(ReviseSize *model.ReviseSize) error {
	if err := mysql.CheckUid2(ReviseSize.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(ReviseSize.Password, ReviseSize.Uid); err != nil {
		return err
	}
	if err := mysql.CheckUidInShoppingCart(ReviseSize.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(ReviseSize.Sid); err != nil {
		return err
	}
	if err := mysql.CheckUid4(ReviseSize.Uid, int64(ReviseSize.Sid)); err != nil {
		return err
	}
	var gid int64
	var err error
	if gid, err = mysql.GetGidBySid(ReviseSize.Sid); err != nil {
		return err
	}
	if err := mysql.CheckSize(ReviseSize.NewSize, gid); err != nil {
		return err
	}
	return mysql.OperateSizeOfGood(ReviseSize.NewSize, ReviseSize.Sid)
}
func OperateNumberOfGood(ReviseNumber *model.ReviseNumber) error {
	if err := mysql.CheckUid2(ReviseNumber.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(ReviseNumber.Password, ReviseNumber.Uid); err != nil {
		return err
	}
	if err := mysql.CheckUidInShoppingCart(ReviseNumber.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(ReviseNumber.Sid); err != nil {
		return err
	}
	if err := mysql.CheckUid4(ReviseNumber.Uid, int64(ReviseNumber.Sid)); err != nil {
		return err
	}
	return mysql.OperateNumberOfGood(ReviseNumber.Sid, ReviseNumber.NewNumber)
}
func DeleteGoodInShoppingCart(DeleteGood *model.DeleteAGoodFormShoppingCart) error {
	if err := mysql.CheckUid2(DeleteGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(DeleteGood.Password, DeleteGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckUidInShoppingCart(DeleteGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(DeleteGood.Sid); err != nil {
		return err
	}
	if err := mysql.CheckUid4(DeleteGood.Uid, int64(DeleteGood.Sid)); err != nil {
		return err
	}
	return mysql.DeleteGoodInShoppingCart(DeleteGood.Sid)
}
func BuyAGoodInShoppingCart(BuyA *model.BuyA) error {
	if err := mysql.CheckUid2(BuyA.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(BuyA.Password, BuyA.Uid); err != nil {
		return err
	}
	if err := mysql.CheckUidInShoppingCart(BuyA.Uid); err != nil {
		return err
	}
	if err := mysql.CheckSidInShoppingCart(BuyA.Sid); err != nil {
		return err
	}
	return mysql.BuyAGoodInShoppingCart(BuyA.Uid, BuyA.Address, BuyA.PhoneNumber, BuyA.RealName, BuyA.Sid)
}
func BuyAllInShoppingCart(BuyAll *model.BuyAll) error {
	if err := mysql.CheckUid2(BuyAll.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(BuyAll.Password, BuyAll.Uid); err != nil {
		return err
	}
	if err := mysql.CheckUidInShoppingCart(BuyAll.Uid); err != nil {
		return err
	}
	if err := mysql.BuyAllInShoppingCart(BuyAll.Uid, BuyAll.Address, BuyAll.PhoneNumber, BuyAll.RealName); err != nil {
		return err
	}
	return mysql.ClearShoppingCart(BuyAll.Uid)
}
func ClearShoppingCart(ClearShoppingCartUser *model.SearchShoppingCart) error {
	if err := mysql.CheckUid2(ClearShoppingCartUser.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(ClearShoppingCartUser.Password, ClearShoppingCartUser.Uid); err != nil {
		return err
	}
	if err := mysql.CheckUidInShoppingCart(ClearShoppingCartUser.Uid); err != nil {
		return err
	}
	return mysql.ClearShoppingCart(ClearShoppingCartUser.Uid)
}
func LoveAGood(LoveGood *model.ShoppingCart) error {
	if err := mysql.CheckUid2(LoveGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(LoveGood.Password, LoveGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckGid(int64(LoveGood.Gid)); err != nil {
		return err
	}
	if err := mysql.CheckSize(LoveGood.Size, int64(LoveGood.Gid)); err != nil {
		return err
	}
	if err := mysql.LoveAGood(LoveGood.Gid, LoveGood.Uid, LoveGood.Number, LoveGood.Size); err != nil {
		return err
	}
	return mysql.AddPrice(int64(LoveGood.Gid))
}

func SearchGoodInShoppingCart(MyLoveGood *model.SearchShoppingCart) (interface{}, error) {
	if err := mysql.CheckUid2(MyLoveGood.Uid); err != nil {
		return nil, err
	}
	if err := mysql.CheckPassword(MyLoveGood.Password, MyLoveGood.Uid); err != nil {
		return nil, err
	}
	return mysql.SearchGoodInShoppingCart(MyLoveGood.Uid)
}
