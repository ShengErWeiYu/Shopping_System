package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
	"log"
)

func CheckOrder(CheckOder *model.CheckOrder) (interface{}, error) {
	if err := mysql.CheckUid2(CheckOder.Uid); err != nil {
		return nil, err
	}
	if err := mysql.CheckPassword(CheckOder.Password, CheckOder.Uid); err != nil {
		return nil, err
	}
	if err := mysql.CheckOrderExist(CheckOder.Uid); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return mysql.CheckOrder(CheckOder.Uid)
}
