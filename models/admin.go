package models

import (
	"gopkg.in/mgo.v2"
)

//获取管理员[登陆]
func LoginGetAdminInfo(db *mgo.Database, collection, uname, passwd string) (info map[string]string, err error) {
	connect := db.C(collection)
	where := M{"uname": uname, "passwd": passwd, "lock": "0"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//获取管理员信息[激活]
func GetNotActivateAdmin(db *mgo.Database, collection, uname string) (info map[string]string, err error) {
	connect := db.C(collection)
	where := M{"uname": uname, "lock": "1"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//获取管理员信息[注册或其他]
func GetAdminInfo(db *mgo.Database, collection, uname string) (info map[string]string, err error) {
	connect := db.C(collection)
	where := M{"uname": uname}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//新增管理员信息
func InsertAdminInfo(db *mgo.Database, collection, uname, passwd, token, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Insert(M{"uname": uname, "passwd": passwd, "role": "administrator", "email": uname, "lock": "1", "token": token, "login_time": nowTime, "update_time": nowTime, "create_time": nowTime})
	return err
}

//设置管理员最后一次登陆时间
func SetAdminLoginTime(db *mgo.Database, collection, uname, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Update(M{"uname": uname}, M{"$set": M{"login_time": nowTime}})
	return err
}

//解锁管理员（激活）
func UnlockAdmin(db *mgo.Database, collection, uname, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Update(M{"uname": uname, "lock": "1"}, M{"$set": M{"lock": "0", "update_time": nowTime}})
	return err
}

//锁定管理员
func LockAdmin(db *mgo.Database, collection, uname, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Update(M{"uname": uname, "lock": "0"}, M{"$set": M{"lock": "1", "update_time": nowTime}})
	return err
}
