package models

//get admin list
func AdminList(account string) (list []interface{}, err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	where := M{}
	if "" != account {
		where = M{"account": account}
	}
	err = connect.Find(where).Select(M{"_id": 0, "passwd": 0}).All(&list)
	if nil == list {
		list = make([]interface{}, 0)
	}
	return list, err
}

//解锁账号（激活）
func UnlockAdmin(account, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	err = connect.Update(M{"account": account, "lock": "1"}, M{"$set": M{"lock": "0", "update_time": nowTime}})
	return err
}

//锁定账号
func LockAdmin(account, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	err = connect.Update(M{"account": account, "lock": "0"}, M{"$set": M{"lock": "1", "update_time": nowTime}})
	return err
}

//获取账号信息[注册或其他]
func GetAdminInfo(account string) (info map[string]string, err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	where := M{"account": account}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}
