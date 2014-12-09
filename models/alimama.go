package models

import (
	"gopkg.in/mgo.v2/bson"

	"errors"
	"reflect"
)

//获取商品列表
func AlimamaList(table map[string]interface{}) (list []map[string]interface{}, count int, err error) {
	connect := MgoCon.DB(SOMI).C(ALIMAMA)
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

//修改商品信息
func UpdateAlimama(id, catid, name, oldPrice, price, sort, status, url, img, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ALIMAMA)
	if !bson.IsObjectIdHex(id) {
		err = errors.New("商品ID有误")
		return err
	}

	where := M{"_id": bson.ObjectIdHex(id)}
	count, err := connect.Find(where).Count()
	if 0 == count {
		err = errors.New("商品不存在")
		return err
	}
	set := M{"catid": catid, "name": name, "oldPrice": oldPrice, "price": price, "sort": sort, "status": status, "url": url, "img": img, "updateTime": nowTime}
	err = connect.Update(where, M{"$set": set})
	return err
}

//添加商品
func AddAlimama(catid, name, oldPrice, price, sort, status, url, img, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ALIMAMA)
	count, err := connect.Find(M{"name": name}).Count()
	if 0 != count {
		err = errors.New("商品已存在")
		return err
	}

	info := M{"catid": catid, "name": name, "oldPrice": oldPrice, "price": price, "sort": sort, "status": status, "url": url, "img": img, "addTime": nowTime, "updateTime": nowTime}
	err = connect.Insert(info)
	return err
}

//获取商品信息
func GetAlimama(id string) (info map[string]interface{}, err error) {
	connect := MgoCon.DB(SOMI).C(ALIMAMA)
	if !bson.IsObjectIdHex(id) {
		err = errors.New("商品ID有误")
		return info, err
	}
	err = connect.Find(M{"_id": bson.ObjectIdHex(id)}).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]interface{})
	} else {
		cid, _ := info["_id"].(bson.ObjectId)
		info["_id"] = cid.Hex()
	}
	return info, err
}

//删除商品
func DelAlimama(id string) (err error) {
	connect := MgoCon.DB(SOMI).C(ALIMAMA)
	if !bson.IsObjectIdHex(id) {
		err = errors.New("商品ID有误")
		return err
	}
	err = connect.Remove(M{"_id": bson.ObjectIdHex(id)})
	return err
}

//上线商品
func OnlineAlimama(id, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ALIMAMA)
	if !bson.IsObjectIdHex(id) {
		err = errors.New("商品ID有误")
		return err
	}
	err = connect.Update(M{"_id": id, "status": "1"}, M{"$set": M{"status": "0", "updateTime": nowTime}})
	return err
}

//下线商品
func OfflineAlimama(id, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ALIMAMA)
	if !bson.IsObjectIdHex(id) {
		err = errors.New("商品ID有误")
		return err
	}
	err = connect.Update(M{"_id": id, "status": "0"}, M{"$set": M{"status": "1", "updateTime": nowTime}})
	return err
}
