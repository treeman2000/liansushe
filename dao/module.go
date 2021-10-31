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
	PageSize int
	PageNum  int
	RoomNum  []int
	MinPrice int
	MaxPrice int
	Center   []string
	MinTerm  int
	MaxTerm  int
	State    string
	UserID   string
}

type HouseInfo struct {
	HouseID   int
	ImgURL    string
	VRURL     string
	Place     string
	Center    string
	Area      int
	Price     int
	Deposit   int
	Room      int
	Hall      int
	Elevator  bool
	Storey    int
	Term      int
	Direction string
	Facility  int
	Note      string
	IsOnline  bool
}

type H []*HouseInfo

func (o H) Price(MinPrice, MaxPrice int) H {
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

func (o H) HouseIDs(houseID ...int) H {
	idMap := make(map[int]bool)
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
	Area      int
	Price     int
	Deposit   int
	Room      int
	Hall      int
	Elevator  bool
	Storey    int
	Term      int
	Direction string
	Facility  int
	Note      string
}

type SetOnlineReq struct {
	Token   string
	HouseID int
}

type SetOfflineReq struct {
	Token   string
	HouseID int
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
	HouseID   int
	SetOnline bool
}

type CollectionSearchReq struct {
	Token    string
	PageSize int
	PageNum  int
	UserID   string
}

type CollectionSearchRsp struct {
	Result     string
	Number     int
	HouseInfos H
}
