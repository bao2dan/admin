package models

import (
	"gopkg.in/mgo.v2"
)

//获取管理员信息
func GetAdminInfo(db *mgo.Database, collection, uname, passwd, lock string) (info map[string]string, err error) {
	connect := db.C(collection)
	err = connect.Find(M{"uname": uname, "passwd": passwd, "lock": lock}).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//新增管理员信息
func InsertAdminInfo(db *mgo.Database, collection, uname, passwd, token string) (err error) {
	connect := db.C(collection)
	err = connect.Insert(M{"uname": uname, "passwd": passwd, "role": "administrator", "email": uname, "lock": "1", "token": token})
	return err
}
