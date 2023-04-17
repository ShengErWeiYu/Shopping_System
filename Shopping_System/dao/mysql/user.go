package mysql

import (
	g "Shopping_System/global"
	"Shopping_System/model"
)

const (
	FindUserByUsernameStr       = "select count(`rank`) from users where username = ?"
	FindUserByUidStr            = "select count(`rank`) from users where uid = ?"
	FindPasswordByUidStr        = "select password from users where uid = ?"
	FindAnswerByUidstr          = "select answer from users where uid = ?"
	ReviseNameByUidStr          = "update users set username=? where uid =?"
	RevisePasswordByUidStr      = "update users set password=? where uid =?"
	ReviseSexByUidStr           = "update users set sex=? where uid=?"
	ResetPasswordByUidStr       = "update users set password=? where uid =?"
	AddUserStr                  = "insert into users (uid, username, password, sex, securityquestion, answer) values(?,?,?,?,?,?)"
	GetUsernameByUidstr         = "select username from users where uid=?"
	GetSecurityQuestionByUidstr = "select securityquestion from users where uid=?"
)

func GetSecurityQuestion(uid string) (interface{}, error) {
	var SecurityQuestion string
	if err := g.Xdb.Get(&SecurityQuestion, GetSecurityQuestionByUidstr, uid); err != nil {
		return "在获取密保问题时出错啦", err
	}
	return SecurityQuestion, nil
}
func GetUsernameByUid(uid string) string {
	var username string
	if err := g.Xdb.Get(&username, GetUsernameByUidstr, uid); err != nil {
		panic(err)
	}
	return username
}
func ReviseSex(uid string, newsex int) error {
	if _, err := g.Xdb.Exec(ReviseSexByUidStr, newsex, uid); err != nil {
		return err
	}
	return nil
}
func RevisePassword(uid string, newpassword string) error {
	if _, err := g.Xdb.Exec(RevisePasswordByUidStr, newpassword, uid); err != nil {
		return err
	}
	return nil
}

func ReviseUsername(uid string, newusername string) error {
	if _, err := g.Xdb.Exec(ReviseNameByUidStr, newusername, uid); err != nil {
		return err
	}
	return nil
}
func ResetPassword(uid string, repassword string) error {
	if _, err := g.Xdb.Exec(ResetPasswordByUidStr, repassword, uid); err != nil {
		return err
	}
	return nil
}

func CheckAnswer(uid string, answer string) error {
	var realanswer string
	if err := g.Xdb.Get(&realanswer, FindAnswerByUidstr, uid); err != nil {
		return err
	}
	if answer != realanswer {
		return ErrorAnswer
	} else {
		return nil
	}
}

func CheckUid2(uid string) error {
	var count int
	if err := g.Xdb.Get(&count, FindUserByUidStr, uid); err != nil {
		return err
	}
	if count == 0 {
		return ErrorUserNotExist

	} else {
		return nil
	}
}
func CheckPassword(password string, uid string) error {
	var realpassword string
	if err := g.Xdb.Get(&realpassword, FindPasswordByUidStr, uid); err != nil {
		return err
	}
	if password != realpassword { //这里有问题
		return ErrorPassword
	} else {
		return nil
	}
}
func CheckUid1(uid string) error {
	var count int
	if err := g.Xdb.Get(&count, FindUserByUidStr, uid); err != nil { //这里有问题
		return err
	}
	if count > 0 {
		return ErrorUserExist //这里留个心眼子
	} else {
		return nil
	}
}
func CheckUsername(username string) error {
	var count int
	if err := g.Xdb.Get(&count, FindUserByUsernameStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist //这里留个心眼子

	} else {
		return nil
	}
}

func AddUser(user *model.User) error {
	_, err := g.Xdb.Exec(AddUserStr, user.Uid, user.Username, user.Password, user.Sex, user.SecurityQuestion, user.Answer)
	return err
}
