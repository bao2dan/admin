package models

import (
	"errors"
)

//获取账号[登陆用]
func LoginGetAdminInfo(account, passwd string) (info map[string]string, err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	where := M{"account": account, "passwd": passwd, "lock": "0"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//获取账号信息[激活用]
func GetNotActivateAdmin(account string) (info map[string]string, err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	where := M{"account": account, "lock": "1"}
	err = connect.Find(where).One(&info)
	if nil != err && NOTFOUND == err.Error() {
		err = nil
		info = make(map[string]string)
	}
	return info, err
}

//新增账号信息
func InsertAdminInfo(account, passwd, token, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	err = connect.Insert(M{"account": account, "passwd": passwd, "role": "guest", "email": account, "lock": "1", "token": token, "login_time": nowTime, "update_time": nowTime, "create_time": nowTime})
	return err
}

//设置账号最后一次登陆时间
func SetAdminLoginTime(account, nowTime string) (err error) {
	connect := MgoCon.DB(SOMI).C(ADMIN_USER)
	err = connect.Update(M{"account": account}, M{"$set": M{"login_time": nowTime}})
	return err
}

//获取导航配置
func GetMenuConfig() (aMenu [][]string, bMenu map[string][]string, urlInfo map[string][]string, err error) {
	//一级导航栏（正序显示）
	aMenu = [][]string{
		{"2", "", "基础"},
		{"3", "", "商家"},
		{"4", "", "活动"},
		{"5", "", "统计"},
		{"6", "/test", "测试"},
		{"1", "", "管理员"},
	}

	//二级导航栏（正序显示）
	bMenu = map[string][]string{
		"2": {"21", "31", "42", "51", "61"},
		"3": {"71"},
		"4": {"81", "82"},
		"5": {"91"},
		"6": {"101"},
		"1": {"11", "12", "13", "14", "15"},
	}

	//所有操作集合（无序）
	urlInfo = map[string][]string{
		"11": {"/admin/list", "管理员列表"},
		"12": {"/admin/update", "编辑管理员"},
		"13": {"/admin/del", "删除管理员"},
		"14": {"/admin/lock", "锁定管理员"},
		"15": {"/admin/unlock", "解锁管理员"},

		"21": {"/country/list", "国家列表"},
		"22": {"/country/add", "新建国家"},
		"23": {"/country/update", "编辑国家"},
		"24": {"/country/del", "删除国家"},

		"31": {"/province/list", "省份列表"},
		"32": {"/province/add", "新建省份"},
		"33": {"/province/update", "编辑省份"},
		"34": {"/province/del", "删除省份"},

		"41": {"/city/list", "城市列表"},
		"42": {"/city/add", "新建城市"},
		"43": {"/city/update", "编辑城市"},
		"44": {"/city/del", "删除城市"},

		"51": {"/region/list", "地区列表"},
		"52": {"/region/add", "新建地区"},
		"53": {"/region/update", "编辑地区"},
		"54": {"/region/del", "删除地区"},

		"61": {"/category/list", "类型列表"},
		"62": {"/category/add", "新建类型"},
		"63": {"/category/update", "编辑类型"},
		"64": {"/category/del", "删除类型"},

		"71": {"/shop/list", "商家列表"},
		"72": {"/shop/add", "新建商家"},
		"73": {"/shop/update", "编辑商家"},
		"74": {"/shop/del", "删除商家"},

		"81": {"/activity/list", "活动列表"},
		"82": {"/activity/add", "新建活动"},
		"83": {"/activity/update", "编辑活动"},
		"84": {"/activity/del", "删除活动"},

		"91": {"/statis/list", "统计列表"},

		"101": {"/test/list", "测试"},
	}

	return aMenu, bMenu, urlInfo, err

}

//获取权限配置
func GetAuthConfig(role string) (auth []string, err error) {
	//权限分配
	auths := map[string][]string{
		"admin1": {
			"3:71", "3:72", "3:73", "3:44",
			"4:81", "4:82", "4:83", "4:84",
			"5:91",
		},
		"admin2": {
			"3:71", "3:72", "3:73", "3:74",
			"4:81", "4:82", "4:83", "4:84",
		},
		"guest": {},
	}
	if v, ok := auths[role]; ok {
		auth = v
	} else {
		err = errors.New("没有此角色")
	}
	return auth, err
}
