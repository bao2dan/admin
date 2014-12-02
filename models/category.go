package models

import (
	"gopkg.in/mgo.v2/bson"

	"errors"
	"reflect"
)

//获取分类列表
func CategoryList(table map[string]interface{}) (list []map[string]interface{}, count int, err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	sWhere, _ := table["sWhere"]
	sort, _ := table["sSort"].(string)

	where := M{}
	rv := reflect.ValueOf(sWhere)
	rk := rv.MapKeys()
	for i := 0; i < len(rk); i++ {
		key := rk[i].String()
		where[key] = rv.MapIndex(rk[i]).Interface()
	}

	count, err = connect.Find(where).Count()
	if nil != err {
		count = 0
	}

	err = connect.Find(where).Sort(sort).All(&list)
	if nil == list {
		list = make([]map[string]interface{}, 0)
	}
	return list, count, err
}

//修改分类信息
func UpdateCategory(_id, name, sort, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if !bson.IsObjectIdHex(_id) {
		err = errors.New("分类ID有误")
		return err
	}
	set := M{"name": name, "sort": sort, "update_time": nowTime}
	err = connect.Update(M{"_id": bson.ObjectIdHex(_id)}, M{"$set": set})
	return err
}

//添加分类
func AddCategory(fid, level, name, sort, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if !bson.IsObjectIdHex(fid) {
		err = errors.New("分类ID有误")
		return err
	}
	if "0" != fid {
		fcount, err := connect.Find(M{"_id": bson.ObjectIdHex(fid)}).Count()
		if 0 == fcount {
			err = errors.New("父分类ID不存在")
			return err
		}
	}

	count, err := connect.Find(M{"name": name, "level": level}).Count()
	if 0 != count {
		err = errors.New("分类已存在")
		return err
	}

	err = connect.Insert(M{"name": name, "fid": fid, "level": level, "sort": sort, "add_time": nowTime, "update_time": nowTime})
	return err
}

//获取分类信息
func GetCategory(_id string) (info M, err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if !bson.IsObjectIdHex(_id) {
		err = errors.New("分类ID有误")
		return info, err
	}
	err = connect.Find(M{"_id": bson.ObjectIdHex(_id)}).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(M)
	} else {
		cid, _ := info["_id"].(bson.ObjectId)
		info["_id"] = cid.Hex()
	}
	return info, err
}

//删除分类
func DelCategory(_id string) (err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if !bson.IsObjectIdHex(_id) {
		err = errors.New("分类ID有误")
		return err
	}
	err = connect.Remove(M{"_id": bson.ObjectIdHex(_id)})
	return err
}
