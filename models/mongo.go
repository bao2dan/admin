package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"errors"
	"strings"
	"time"
)

type (
	M bson.M
)

const (
	NOTFOUND   string = "not found" //not found one from mongo
	SOMI              = "somi"
	ADMIN_USER        = "admin.user"
	CATEGORY          = "category"
)

var (
	MgoCon *mgo.Session
)

//connect mongodb
func ConnectMgo(confName string) (session *mgo.Session, err error) {
	if "" == confName {
		err = errors.New("mgo config name is empty")
		return session, err
	}

	//get mongo config
	url_ir, _ := beego.GetConfig("string", confName)
	url, _ := url_ir.(string)
	if "" == url {
		err = errors.New("mgo url is empty")
		return session, err
	}
	if strings.HasPrefix(url, "mongodb://") {
		url = url[10:]
	}
	if c := strings.Index(url, "?"); c != -1 {
		url = url[:c]
	}

	//connect + timeout
	session, err = mgo.DialWithTimeout(url, 5*time.Second)
	if nil != err {
		return session, err
	} else {
		session.SetSyncTimeout(5 * time.Second)
		session.SetSocketTimeout(5 * time.Second)
	}

	session.SetMode(mgo.Monotonic, true)
	return session, err
}
