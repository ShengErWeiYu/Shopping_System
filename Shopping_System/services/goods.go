package services

import (
	"Shopping_System/dao/mysql"
	"Shopping_System/model"
)

func GetGoodDetail(GoodDetail *model.GetGoodDetail) (interface{}, error) {
	if err := mysql.CheckGid(GoodDetail.Gid); err != nil {
		return nil, err
	}
	return mysql.GetGoodDetail(GoodDetail.Gid)
}
func ChangeTheSizeOfGood(ChangeSizeGood *model.ChangeSizeGood) error {
	if err := mysql.CheckUid2(ChangeSizeGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(ChangeSizeGood.Password, ChangeSizeGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckGid(ChangeSizeGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(ChangeSizeGood.Uid, ChangeSizeGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeTheSizeOfGood(ChangeSizeGood.Gid, ChangeSizeGood.NewSize)
}
func ChangeTheIntroductionOfGood(ChangeIntroductionGood *model.ChangeIntroductionGood) error {
	if err := mysql.CheckUid2(ChangeIntroductionGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(ChangeIntroductionGood.Password, ChangeIntroductionGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckGid(ChangeIntroductionGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(ChangeIntroductionGood.Uid, ChangeIntroductionGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeTheIntroductionOfGood(ChangeIntroductionGood.Gid, ChangeIntroductionGood.NewIntroduction)
}
func ChangeThePriceOfGood(ChangePriceGood *model.ChangePriceGood) error {
	if err := mysql.CheckUid2(ChangePriceGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(ChangePriceGood.Password, ChangePriceGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckGid(ChangePriceGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(ChangePriceGood.Uid, ChangePriceGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeThePriceOfGood(ChangePriceGood.Gid, ChangePriceGood.NewPrice)
}
func ChangeTheNameOfGood(ChangeNameGood *model.ChangeNameGood) error {
	if err := mysql.CheckUid2(ChangeNameGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(ChangeNameGood.Password, ChangeNameGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckGid(ChangeNameGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(ChangeNameGood.Uid, ChangeNameGood.Gid); err != nil {
		return err
	}
	return mysql.ChangeTheNameOfGood(ChangeNameGood.Gid, ChangeNameGood.NewGoodName)
}
func DeleteAGood(DeleteGood *model.GoodsDelete) error {
	if err := mysql.CheckUid2(DeleteGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(DeleteGood.Password, DeleteGood.Uid); err != nil {
		return err
	}
	if err := mysql.CheckGid(DeleteGood.Gid); err != nil {
		return err
	}
	if err := mysql.CheckUid3(DeleteGood.Uid, DeleteGood.Gid); err != nil {
		return err
	}
	return mysql.DeleteAGood(DeleteGood.Gid)
}
func SearchMyGoods(Search *model.SearchMyGoods) (interface{}, error) {
	if err := mysql.CheckUid2(Search.Uid); err != nil {
		return nil, err
	}
	return mysql.SearchMyGoods(Search.Uid)
}
func PublishAGood(Good *model.GoodsPublish) error {
	if err := mysql.CheckUid2(Good.Uid); err != nil {
		return err
	}
	if err := mysql.CheckPassword(Good.Password, Good.Uid); err != nil {
		return err
	}
	return mysql.PublishAGood(Good.GoodName, Good.Price, Good.Introduction, Good.Size, Good.Uid)
}
