package dao

func init() {
	Tokens = make(map[string]string)
	UserID2NameMap = make(map[string]string)
	UserID2PwdMap = make(map[string]string)
	HouseInfos = make([]HouseInfo, 0)
	UUID2ImagePath = make(map[string]string)
}

var Tokens map[string]string

var UserID2PwdMap map[string]string

var UserID2NameMap map[string]string

var HouseInfos []HouseInfo

var UUID2ImagePath map[string]string
