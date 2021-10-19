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

type HouseAddReq struct {
	Token     string
	HouseID   int
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
	HouseID string
}

type SetOfflineReq struct {
	Token   string
	HouseID string
}
