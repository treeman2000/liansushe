package dao

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

func (o H) FilterPrice(h *HouseSearchReq) H {
	res := make(H, 0)
	for _, hInfo := range o {
		if hInfo.Price >= h.MinPrice && (h.MaxPrice == 0 || hInfo.Price <= h.MaxPrice) {
			res = append(res, hInfo)
		}
	}
	return res
}

func (o H) FilterIsOnline(h *HouseSearchReq) H {
	if h.State == "all" {
		return o
	}
	res := make(H, 0)
	for _, hInfo := range o {
		if (h.State == "online" || h.State == "") && hInfo.IsOnline {
			res = append(res, hInfo)
		} else if h.State == "offline" && !hInfo.IsOnline {
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
