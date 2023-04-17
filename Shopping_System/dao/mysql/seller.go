package mysql

import (
	g "Shopping_System/global"
	"Shopping_System/model"
)

const (
	FindOrderDetailByGidStr = "select oid,uid,gid,number,size,address,phonenumber,realname from orders where gid = ?"
	FindOrderByGidStr       = "select count(oid) from orders where gid = ?"
)

func CheckOrderExist(uid string) error {
	var a []string
	if err := g.Xdb.Select(&a, "select gid from goods where uid=?", uid); err != nil {
		return err
	} else {
		var count []int
		for _, v := range a {
			if err := g.Xdb.Select(&count, FindOrderByGidStr, v); err != nil {
				return err
			}
		}
		if len(count) == 0 {
			return ErrorOrderExist
		} else {
			return nil
		}
	}
}
func CheckOrder(uid string) ([]model.CheckOrderResult, error) {
	var Order []model.CheckOrderResult
	var a []string
	if err := g.Xdb.Select(&a, "select gid from goods where uid = ?", uid); err != nil {
		return nil, err
	} else {
		for _, v := range a {
			var o []model.CheckOrderResult
			if err := g.Xdb.Select(&o, FindOrderDetailByGidStr, v); err != nil {
				return nil, err
			}
			Order = append(Order, o...)
		}
	}
	return Order, nil
}
