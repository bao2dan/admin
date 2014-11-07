package models

import (
	"gopkg.in/mgo.v2"
)

//获取账号[登陆]
func LoginGetAdminInfo(account, passwd string) (info map[string]string, err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	where := M{"account": account, "passwd": passwd, "lock": "0"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//获取账号信息[激活]
func GetNotActivateAdmin(mgocon *mgo.Session, account string) (info map[string]string, err error) {
	connect := mgocon.DB(SOMI).C(ADMIN_USER)
	where := M{"account": account, "lock": "1"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//新增账号信息
func InsertAdminInfo(mgocon *mgo.Session, account, passwd, token, nowTime string) (err error) {
	connect := mgocon.DB(SOMI).C(ADMIN_USER)
	err = connect.Insert(M{"account": account, "passwd": passwd, "role": "root", "email": account, "lock": "1", "token": token, "login_time": nowTime, "update_time": nowTime, "create_time": nowTime})
	return err
}

//设置账号最后一次登陆时间
func SetAdminLoginTime(mgocon *mgo.Session, account, nowTime string) (err error) {
	connect := mgocon.DB(SOMI).C(ADMIN_USER)
	err = connect.Update(M{"account": account}, M{"$set": M{"login_time": nowTime}})
	return err
}
