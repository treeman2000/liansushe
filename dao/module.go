package dao

import "log"

type LoginReq struct {
	UserID   string
	Password string
}

type LoginRsp struct {
	Result   string
	Token    string
	UserName string
}

type RegisterReq struct {
	UserName string
	UserID   string
	Password string
}

type HouseSearchReq struct {
	PageSize int32
	PageNum  int32
	RoomNum  []int32
	MinPrice int32
	MaxPrice int32
	Center   []string
	MinTerm  int32
	MaxTerm  int32
	State    string
	UserID   string
}

type HouseInfo struct {
	HouseID   int32
	ImgURL    string
	VRURL     string
	Place     string
	Center    string
	Area      int32
	Price     int32
	Deposit   int32
	Room      int32
	Hall      int32
	Elevator  bool
	Storey    int32
	Term      int32
	Direction string
	Facility  int32
	Note      string
	IsOnline  bool
}

type H []*HouseInfo

func (o H) Price(MinPrice, MaxPrice int32) H {
	res := make(H, 0)
	for _, hInfo := range o {
		if hInfo.Price >= MinPrice && (MaxPrice == 0 || hInfo.Price <= MaxPrice) {
			res = append(res, hInfo)
		}
	}
	return res
}

func (o H) IsOnline(State string) H {
	if State == "all" {
		return o
	}
	res := make(H, 0)
	for _, hInfo := range o {
		if (State == "online" || State == "") && hInfo.IsOnline {
			res = append(res, hInfo)
		} else if State == "offline" && !hInfo.IsOnline {
			res = append(res, hInfo)
		}
	}
	return res
}

func (o H) HouseIDs(houseID ...int32) H {
	idMap := make(map[int32]bool)
	for _, hID := range houseID {
		log.Printf("need %v\n", hID)
		idMap[hID] = true
	}
	res := make(H, 0)
	for _, hInfo := range o {
		log.Printf("has %v\n", hInfo.HouseID)
		if idMap[hInfo.HouseID] {
			res = append(res, hInfo)
		}
	}
	return res
}

type HouseAddReq struct {
	Token     string
	Image     string
	VRFile    string
	Place     string
	Center    string
	Area      int32
	Price     int32
	Deposit   int32
	Room      int32
	Hall      int32
	Elevator  bool
	Storey    int32
	Term      int32
	Direction string
	Facility  int32
	Note      string
}

type SetOnlineReq struct {
	Token   string
	HouseID int32
}

type SetOfflineReq struct {
	Token   string
	HouseID int32
}

type VerifyReq struct {
	EmailAddress string
}

type VerifyRsp struct {
	Result           string
	VerificationCode string
}

type RegisterV2Req struct {
	UserName         string
	EmailAddress     string
	Password         string
	VerificationCode string
}

type CollectionChangeReq struct {
	Token     string
	UserID    string
	HouseID   int32
	SetOnline bool
}

type CollectionSearchReq struct {
	Token    string
	PageSize int32
	PageNum  int32
	UserID   string
}

type CollectionSearchRsp struct {
	Result     string
	Number     int32
	HouseInfos H
}

type Profile struct {
	UserID      string
	Age         int32
	Gender      string
	Major       string
	Hometown    string
	Character   string
	Hobby       string
	SleepTime   string
	AwakeTime   string
	Expectation string
	WantHouse   int32
	AvatarURL   string
}

type ProfileAddReq struct {
	Token   string
	Profile Profile
}

type ProfileGetReq struct {
	Token  string
	UserID string
}

type ProfileGetRsp struct {
	Result  string
	Profile Profile
}

type HouseGetReq struct {
	Token   string
	HouseID int32
}

type HouseGetRsp struct {
	Result    string
	HouseInfo HouseInfo
}

type ProfileSearchReq struct {
	Token    string
	UserID   string
	Keyword  string
	PageSize int32
	PageNum  int32
}

type ProfileSearchRsp struct {
	Result   string
	Number   int32
	Profiles []Profile
}

type ChatListReq struct {
	Token  string
	UserID string
}

type PersonInfo struct {
	UserID   string
	UserName string
}
