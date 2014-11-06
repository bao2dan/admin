package models

import (
	"gopkg.in/mgo.v2"
)

//解锁账号（激活）
func UnlockAdmin(db *mgo.Database, collection, account, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Update(M{"account": account, "lock": "1"}, M{"$set": M{"lock": "0", "update_time": nowTime}})
	return err
}

//锁定账号
func LockAdmin(db *mgo.Database, collection, account, nowTime string) (err error) {
	connect := db.C(collection)
	err = connect.Update(M{"account": account, "lock": "0"}, M{"$set": M{"lock": "1", "update_time": nowTime}})
	return err
}

//获取账号信息[注册或其他]
func GetAdminInfo(db *mgo.Database, collection, account string) (info map[string]string, err error) {
	connect := db.C(collection)
	where := M{"account": account}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}
