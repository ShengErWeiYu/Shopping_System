package mysql

import (
	g "Shopping_System/global"
	"Shopping_System/model"
)

const (
	AddGoodStr                = "insert into goods (goodname, price, introduction, size, uid) values (?,?,?,?,?)"
	FindGoodByUidStr          = "select gid,goodname from goods where gid>0 and uid=?"
	CheckGidStr               = "select count(goodname) from goods where gid =?"
	FindUidByGidStr           = "select uid from goods where gid =?"
	DeleteGoodStr             = "delete from goods where gid = ?"
	ChangeGoodNameStr         = "update goods set goodname=? where gid =?"
	ChangeGoodPriceStr        = "update goods set price=? where gid =?"
	ChangeGoodIntroductionStr = "update goods set introduction=? where gid =?"
	ChangeGoodSizeStr         = "update goods set size=? where gid =?"
	FindGoodDetailByGidStr    = "select gid,goodname,price,introduction,size,uid from goods where gid=?"
)

func GetGoodDetail(gid int64) (interface{}, error) {
	var getdetail model.GetAllGoods
	if err := g.Xdb.Get(&getdetail, FindGoodDetailByGidStr, gid); err != nil {
		return nil, err
	}
	return getdetail, nil
}
func ChangeTheSizeOfGood(gid int64, newsize string) error {
	if _, err := g.Xdb.Exec(ChangeGoodSizeStr, newsize, gid); err != nil {
		return err
	}
	return nil
}
func ChangeTheIntroductionOfGood(gid int64, newintroduction string) error {
	if _, err := g.Xdb.Exec(ChangeGoodIntroductionStr, newintroduction, gid); err != nil {
		return err
	}
	return nil
}
func ChangeThePriceOfGood(gid int64, newprice string) error {
	if _, err := g.Xdb.Exec(ChangeGoodPriceStr, newprice, gid); err != nil {
		return err
	}
	return nil
}
func ChangeTheNameOfGood(gid int64, newgoodname string) error {
	if _, err := g.Xdb.Exec(ChangeGoodNameStr, newgoodname, gid); err != nil {
		return err
	}
	return nil
}
func CheckGid(gid int64) error {
	var count int
	if err := g.Xdb.Get(&count, CheckGidStr, gid); err != nil {
		return err
	}
	if count > 0 {
		return nil
	} else {
		return ErrorGoodExist
	}
}
func DeleteAGood(gid int64) error {
	if _, err := g.Xdb.Exec(DeleteGoodStr, gid); err != nil {
		return err
	}
	return nil
}
func CheckUid3(uid string, gid int64) error {
	var realuid string
	if err := g.Xdb.Get(&realuid, FindUidByGidStr, gid); err != nil {
		return err
	}
	if uid != realuid {
		return ErrorUserOperation
	}
	return nil
}
func SearchMyGoods(uid string) (interface{}, error) {
	var goods []model.SearchResult
	if err := g.Xdb.Select(&goods, FindGoodByUidStr, uid); err != nil {
		return nil, err
	} else {
		return goods, err
	}
}
func PublishAGood(goodname string, price string, introduce string, size string, uid string) error {
	if _, err := g.Xdb.Exec(AddGoodStr, goodname, price, introduce, size, uid); err != nil {
		return err
	}
	return nil
}
