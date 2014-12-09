package models

import (
	"reflect"
)

//获取管理员账号列表
func AdminList(table map[string]interface{}) (list []map[string]interface{}, count int, err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	sWhere, _ := table["sWhere"]
	skip, _ := table["iDisplayStart"].(int)
	limit, _ := table["iDisplayLength"].(int)
	sort, _ := table["sSort"].(string)

	where := M{}
	rv := reflect.ValueOf(sWhere)
	rk := rv.MapKeys()
	for i := 0; i < len(rk); i++ {
		key := rk[i].String()
		where[key] = rv.MapIndex(rk[i]).Interface()
	}
	where["role"] = M{"$ne": "root"}

	count, err = connect.Find(where).Count()
	if nil != err {
		count = 0
	}

	err = connect.Find(where).Select(M{"_id": 0, "passwd": 0}).Skip(skip).Limit(limit).Sort(sort).All(&list)
	if nil == list {
		list = make([]map[string]interface{}, 0)
	}
	return list, count, err
}

//修改管理员信息
func UpdateAdmin(account, passwd, name, phone, email, sex, role, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	set := M{"name": name, "phone": phone, "email": email, "sex": sex, "updateTime": nowTime}
	if "" != passwd {
		set["passwd"] = passwd
	}
	if "" != role {
		set["role"] = role
	}
	err = connect.Update(M{"account": account}, M{"$set": set})
	return err
}

//解锁(激活)管理员账号
func UnlockAdmin(account, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	err = connect.Update(M{"account": account, "lock": "1"}, M{"$set": M{"lock": "0", "updateTime": nowTime}})
	return err
}

//锁定管理员账号
func LockAdmin(account, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	err = connect.Update(M{"account": account, "lock": "0"}, M{"$set": M{"lock": "1", "updateTime": nowTime}})
	return err
}

//删除管理员账号
func DelAdmin(account string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	err = connect.Remove(M{"account": account})
	return err
}

//获取管理员账号信息[注册或其他]
func GetAdmin(account string) (info map[string]string, err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	where := M{"account": account}
	err = connect.Find(where).One(&info)
	return info, err
}
