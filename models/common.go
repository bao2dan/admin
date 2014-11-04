package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

type (
	M bson.M

	city struct {
		City_id   string `json:"city_id"`
		City_name string `json:"city_name"`
	}

	OrderStatus struct {
		AppKey           string `bson:"AppKey"`
		Method           string `bson:"Method"`
		MsgCode          string `bson:"MsgCode"`
		MsgCont          string `bson:"MsgCont"`
		OrderItemGroupId string `bson:"OrderItemGroupId"`
		OrderNo          string `bson:"OrderNo"`
		OrderStatus      string `bson:"OrderStatus"`
		OrderStatusTime  string `bson:"OrderStatusTime"`
		ReqTime          string `bson:"ReqTime"`
		TranscationID    string `bson:"TranscationID"`
	}
)

const (
	NOTFOUND string = "not found"
)

//connect mongodb
func ConnectMgo(url string) (db *mgo.Database, err error) {
	if "" == url {
		err = errors.New("mongo url is empty")
		return db, err
	}

	//get dbname
	dbname := ""
	if strings.HasPrefix(url, "mongodb://") {
		url = url[10:]
	}
	if c := strings.Index(url, "?"); c != -1 {
		url = url[:c]
	}
	if c := strings.Index(url, "/"); c != -1 {
		dbname = url[c+1:]
		dbname = strings.TrimSpace(dbname)
	}

	if "" == dbname {
		err = errors.New("mongo db name is empty")
		return db, err
	}

	session, err := mgo.DialWithTimeout(url, 3*time.Second)
	if nil != err {
		return db, err
	} else {
		session.SetSyncTimeout(3 * time.Second)
		session.SetSocketTimeout(3 * time.Second)
	}

	session.SetMode(mgo.Monotonic, true)
	db = session.DB(dbname)
	return db, err
}
