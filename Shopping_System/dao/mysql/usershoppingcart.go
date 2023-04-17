package mysql

import (
	g "Shopping_System/global"
	"Shopping_System/model"
	"log"
)

const (
	AddLoveStr                        = "insert into shoppingcart (gid,uid,price,number,size) value (?,?,'0',?,?)"
	FindSizeByGidStr                  = "select count(goodname) from goods where size like ? and gid =?"
	FindPriceByGidStr                 = "select price from goods where gid = ?"
	AddPriceStr                       = "update shoppingcart set price = ? where gid =?"
	FindGoodInShoppingCartByUid       = "select sid,gid,number,size,price from shoppingcart where sid>0 and uid = ?"
	ClearShoppingCartStr              = "delete from shoppingcart where uid=?"
	FindSidByUidStr                   = "select count(sid) from shoppingcart where uid =?"
	BuyAllGoodsInUserShoppingCartStr  = "insert into orders(gid, uid,  number, size, address, phonenumber, realname) values (?,?,?,?,?,?,?)"
	BuyAGoodsInUserShoppingCartStr    = "insert into orders(gid, uid,  number, size, address, phonenumber, realname) values (?,?,?,?,?,?,?)"
	FindGoodDetailByUidStr            = "select gid,number,size from shoppingcart where uid=?"
	FindUidBySidStr                   = "select uid from shoppingcart where sid =?"
	ReviseGoodInShoppingCartNumberStr = "update shoppingcart set number=? where sid=?"
	FindThisGoodInShoppingCartStr     = "select count(uid) from shoppingcart where sid=?"
	GetGidBySidStr                    = "select gid from shoppingcart where sid=?"
	ReviseGoodInShoppingCartSizeStr   = "update shoppingcart set size=? where sid=?"
	FindGoodDetailBySidStr            = "select gid,number,size from shoppingcart where sid=?"
	DeleteAGoodFormShoppingCartStr    = "delete from shoppingcart where sid=?"
)

func OperateSizeOfGood(newsize string, sid int) error {
	if _, err := g.Xdb.Exec(ReviseGoodInShoppingCartSizeStr, newsize, sid); err != nil {
		return err
	}
	return nil
}
func GetGidBySid(sid int) (int64, error) {
	var gid int64
	if err := g.Xdb.Get(&gid, GetGidBySidStr, sid); err != nil {
		return 0, err
	}
	return gid, nil
}
func CheckSidInShoppingCart(sid int) error {
	var count int
	if err := g.Xdb.Get(&count, FindThisGoodInShoppingCartStr, sid); err != nil {
		return err
	}
	if count == 0 {
		return ErrorThisGoodInShoppingCart
	}
	return nil
}
func OperateNumberOfGood(sid int, newnumber int) error {
	if _, err := g.Xdb.Exec(ReviseGoodInShoppingCartNumberStr, newnumber, sid); err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
func CheckUid4(uid string, sid int64) error {
	var realuid string
	if err := g.Xdb.Get(&realuid, FindUidBySidStr, sid); err != nil {
		return err
	}
	if uid != realuid {
		return ErrorUserOperation
	}
	return nil
}
func DeleteGoodInShoppingCart(sid int) error {
	if _, err := g.Xdb.Exec(DeleteAGoodFormShoppingCartStr, sid); err != nil {
		return err
	}
	return nil
}
func BuyAGoodInShoppingCart(uid string, address string, phonenumber string, realname string, sid int) error {
	var GoodDetail []model.GoodGids
	if err := g.Xdb.Select(&GoodDetail, FindGoodDetailBySidStr, sid); err != nil {
		return err
	}
	for _, v := range GoodDetail {
		if _, err := g.Xdb.Exec(BuyAGoodsInUserShoppingCartStr, v.Gid, uid, v.Number, v.Size, address, phonenumber, realname); err != nil {
			return err
		}
	}
	return nil
}
func BuyAllInShoppingCart(uid string, address string, phonenumber string, realname string) error {
	var GoodDetail []model.GoodGids
	if err := g.Xdb.Select(&GoodDetail, FindGoodDetailByUidStr, uid); err != nil {
		log.Fatalln(err)
		return err
	}
	for _, v := range GoodDetail {
		if _, err := g.Xdb.Exec(BuyAllGoodsInUserShoppingCartStr, v.Gid, uid, v.Number, v.Size, address, phonenumber, realname); err != nil {
			return err
		}
	}
	return nil
}
func CheckUidInShoppingCart(uid string) error {
	var count int
	if err := g.Xdb.Get(&count, FindSidByUidStr, uid); err != nil {
		log.Fatalln(err)
		return err
	}
	if count == 0 {
		return ErrorGoodExistInShoppingCart
	}
	return nil
}
func ClearShoppingCart(uid string) error {
	if _, err := g.Xdb.Exec(ClearShoppingCartStr, uid); err != nil {
		return err
	}
	return nil
}
func SearchGoodInShoppingCart(uid string) (interface{}, error) {
	var MyLoveGoods []model.SearchShoppingCartResult
	if err := g.Xdb.Select(&MyLoveGoods, FindGoodInShoppingCartByUid, uid); err != nil {
		return nil, err
	}
	return MyLoveGoods, nil
}
func AddPrice(gid int64) error {
	var price string
	if err := g.Xdb.Get(&price, FindPriceByGidStr, gid); err != nil {
		return err
	}
	if _, err := g.Xdb.Exec(AddPriceStr, price, gid); err != nil {
		return err
	}
	return nil
}
func CheckSize(size string, gid int64) error {
	var count int
	if err := g.Xdb.Get(&count, FindSizeByGidStr, "%"+size+"%", gid); err != nil {
		return err
	}
	if count == 0 {
		return ErrorSizeExist
	}
	return nil
}
func LoveAGood(gid int, uid string, number int, size string) error {
	if _, err := g.Xdb.Exec(AddLoveStr, gid, uid, number, size); err != nil {
		return err
	}
	return nil
}
