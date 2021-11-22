package dao

func init() {
	Tokens = make(map[string]string)
	UserID2NameMap = make(map[string]string)
	UserID2PwdMap = make(map[string]string)
	HouseInfos = make(H, 0)
	UUID2ImagePath = make(map[string]string)
	UserID2VerifyCode = make(map[string]string)
	Collection = make(map[string][]int32)
	Profiles = make(map[string]Profile)
	// 便于测试
	UserID2NameMap["defaultID"] = "defaultUser"
	UserID2PwdMap["defaultID"] = "defaultPwd"
}

var Tokens map[string]string

var UserID2PwdMap map[string]string

var UserID2NameMap map[string]string

var HouseInfos H

var UUID2ImagePath map[string]string

var UserID2VerifyCode map[string]string

var Collection map[string][]int32

var Profiles map[string]Profile
