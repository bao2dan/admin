package models

import (
	"gopkg.in/mgo.v2"
)

//获取账号[登陆]
func LoginGetAdminInfo(db *mgo.Database, collection, account, passwd string) (info map[string]string, err error) {
	connect := db.C(collection)
	where := M{"account": account, "passwd": passwd, "lock": "0"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//获取账号信息[激活]
func GetNotActivateAdmin(db *mgo.Database, collection, account string) (info map[string]string, err error) {
	connect := db.C(collection)
	where := M{"account": account, "lock": "1"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//新增账号信息
func InsertAdminInfo(db *mgo.Database, collection, account, passwd, token, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Insert(M{"account": account, "passwd": passwd, "role": "root", "email": account, "lock": "1", "token": token, "login_time": nowTime, "update_time": nowTime, "create_time": nowTime})
	return err
}

//设置账号最后一次登陆时间
func SetAdminLoginTime(db *mgo.Database, collection, account, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Update(M{"account": account}, M{"$set": M{"login_time": nowTime}})
	return err
}
