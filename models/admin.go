package models

import (
	"labix.org/v2/mgo"
)

//get admin info
func AGetAdminInfo(db *mgo.Database, collection, uname, passwd string) (info map[string]string, err error) {
	connect := db.C(collection)
	err = connect.Find(M{"uname": uname, "passwd": passwd}).One(&info)

	//connect.Insert(M{"uname": uname, "passwd": passwd, "age": "30", "work": "development", "address": "beijing qing he wu cai cheng NO56", "role": "admin"})
	return info, err
}
