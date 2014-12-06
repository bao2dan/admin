package models

import (
	"gopkg.in/mgo.v2/bson"

	"errors"
	"strconv"
)

func CategoryTreeData(catid string) (list []map[string]interface{}, err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	err = connect.Find(M{}).Sort("-sort").All(&list)
	if nil == list {
		list = make([]map[string]interface{}, 0)
	}

	newlist := make([]map[string]interface{}, 0)
	newlist = CategoryRecursionTree(list, catid, "0", 0)
	return newlist, err
}

//递归并处理分类树的展示结构数据
//@param list 原始数据
//@param newlist 新数据
//@param catid 要选中的分类ID(如果不需要选中则置空)
//@param f 父分类ID
//@param n 上一级分类的级数
func CategoryRecursionTree(list []map[string]interface{}, catid, f string, n int) []map[string]interface{} {
	n = n + 1
	sonList := make([]map[string]interface{}, 0)
	for _, row := range list {
		ele := make(map[string]interface{})
		levelstr, _ := row["level"].(string)
		level, _ := strconv.Atoi(levelstr)
		fid, _ := row["fid"].(string)
		if n == level && fid == f {
			//分类ID的处理
			cid, _ := row["_id"].(bson.ObjectId)
			rowcatid := cid.Hex()
			ele["catid"] = rowcatid
			ele["level"] = levelstr

			//分类名称的处理
			name, _ := row["name"].(string)
			ele["name"] = name

			//递归
			resList := CategoryRecursionTree(list, catid, rowcatid, n)

			addParam := make(map[string]interface{})
			if len(resList) > 0 {
				addParam["children"] = resList
				ele["type"] = "folder"
			} else {
				ele["type"] = "item"
				if rowcatid == catid {
					addParam["item-selected"] = true
				}
			}
			if len(addParam) > 0 {
				ele["additionalParameters"] = addParam
			}
			sonList = append(sonList, ele)
		}
	}
	return sonList
}

//获取分类列表
func CategoryList() (list []map[string]interface{}, count int, err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)

	where := M{}
	count, err = connect.Find(where).Count()
	if nil != err {
		count = 0
	}

	err = connect.Find(where).Sort("-sort").All(&list)
	if nil == list {
		list = make([]map[string]interface{}, 0)
	}
	return list, count, err
}

//递归并处理分类列表的展示结构数据
//@param list 原始数据
//@param newlist 新数据
//@param f 父分类ID
//@param prestr 根据级数，要在分类名称前要加的字符串
//@param n 上一级分类的级数
func CategoryRecursionList(list, newlist []map[string]interface{}, f, prestr string, n int) []map[string]interface{} {
	n = n + 1
	for _, row := range list {
		fid, _ := row["fid"].(string)
		if fid == f {
			//分类ID的处理
			cid, _ := row["_id"].(bson.ObjectId)
			catId := cid.Hex()
			row["_id"] = catId

			//分类名称的处理
			name, _ := row["name"].(string)
			for i := 1; i < n; i++ {
				name = prestr + name
			}
			row["name"] = name
			newlist = append(newlist, row)

			//递归
			newlist = CategoryRecursionList(list, newlist, catId, prestr, n)
		}
	}
	return newlist
}

//修改分类信息
func UpdateCategory(catid, fid, level, name, sort, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if !bson.IsObjectIdHex(catid) {
		err = errors.New("分类ID有误")
		return err
	}
	set := M{"fid": fid, "level": level, "name": name, "sort": sort, "update_time": nowTime}
	err = connect.Update(M{"_id": bson.ObjectIdHex(catid)}, M{"$set": set})
	return err
}

//添加分类
func AddCategory(fid, level, name, sort, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if "0" != fid {
		if !bson.IsObjectIdHex(fid) {
			err = errors.New("分类ID有误")
			return err
		}
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
func GetCategory(catid string) (info map[string]interface{}, err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if !bson.IsObjectIdHex(catid) {
		err = errors.New("分类ID有误")
		return info, err
	}
	err = connect.Find(M{"_id": bson.ObjectIdHex(catid)}).One(&info)
	if nil == err {
		cid, _ := info["_id"].(bson.ObjectId)
		info["_id"] = cid.Hex()
	}
	return info, err
}

//获取某分类的子分类信息
func GetSonCategory(fid string) (list []map[string]interface{}, err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	err = connect.Find(M{"fid": fid}).All(&list)
	if nil == list {
		list = make([]map[string]interface{}, 0)
	}
	return list, err
}

//删除分类
func DelCategory(catid string) (err error) {
	connect := MgoCon.DB(SOMI).C(CATEGORY)
	if !bson.IsObjectIdHex(catid) {
		err = errors.New("分类ID有误")
		return err
	}
	err = connect.Remove(M{"_id": bson.ObjectIdHex(catid)})
	return err
}
