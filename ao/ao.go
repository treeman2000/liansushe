package ao

import (
	"fmt"
	"liansushe/chatManager"
	"liansushe/config"
	"liansushe/dao"
	"log"
	"math/rand"
	"time"

	"github.com/go-gomail/gomail"
)

type AO struct{}

var Ao = &AO{}

func (o *AO) Login(req *dao.LoginReq) (rsp *dao.LoginRsp, err error) {
	rsp = &dao.LoginRsp{}
	if dao.UserID2PwdMap[req.UserID] == req.Password {
		rsp.Result = "OK"
		rsp.Token = req.UserID + "_" + req.Password
		rsp.UserName = dao.UserID2NameMap[req.UserID]
	} else {
		rsp.Result = "没有该用户"
		return nil, err
	}
	return rsp, nil
}

func (o *AO) Register(req *dao.RegisterReq) (Result string, err error) {
	if req.UserName == "" || req.UserID == "" || req.Password == "" {
		return "请填写完整", nil
	}
	dao.UserID2PwdMap[req.UserID] = req.Password
	dao.UserID2NameMap[req.UserID] = req.UserName
	return "OK", nil
}

func (i *AO) HouseAdd(req *dao.HouseAddReq) (Result string, err error) {
	//convert 2 houseInfo
	HouseID := int32(time.Now().UnixNano())
	if HouseID < 0 {
		HouseID *= -1
	}
	houseInfo := &dao.HouseInfo{
		HouseID:   HouseID,
		ImgURL:    "/image/" + req.Image,
		VRURL:     req.VRFile,
		Place:     req.Place,
		Center:    req.Center,
		Area:      req.Area,
		Price:     req.Price,
		Deposit:   req.Deposit,
		Room:      req.Room,
		Hall:      req.Hall,
		Elevator:  req.Elevator,
		Storey:    req.Storey,
		Term:      req.Term,
		Direction: req.Direction,
		Facility:  req.Facility,
		Note:      req.Note,
		IsOnline:  true,
	}

	dao.HouseInfos = append(dao.HouseInfos, houseInfo)
	return "OK", nil
}

func (i *AO) SetOnline(req *dao.SetOnlineReq) (Result string, err error) {
	// houseID, err := strconv.Atoi(req.HouseID)
	if err != nil {
		log.Println("[SetOnline]", err)
		return err.Error(), err
	}
	for i, houseInfo := range dao.HouseInfos {
		if houseInfo.HouseID == req.HouseID {
			dao.HouseInfos[i].IsOnline = true
		}
	}
	return "OK", nil
}

func (i *AO) SetOffline(req *dao.SetOfflineReq) (Result string, err error) {
	// houseID, err := strconv.Atoi(req.HouseID)
	if err != nil {
		log.Println("[SetOffline]", err)
		return err.Error(), err
	}
	for i, houseInfo := range dao.HouseInfos {
		if houseInfo.HouseID == req.HouseID {
			dao.HouseInfos[i].IsOnline = false
		}
	}
	return "OK", nil
}

func (i *AO) HouseSearch(req *dao.HouseSearchReq) (dao.H, error) {
	return dao.HouseInfos.Price(req.MinPrice, req.MaxPrice).IsOnline(req.State), nil
}

func (i *AO) HouseGet(req *dao.HouseGetReq) (*dao.HouseInfo, error) {
	return dao.HouseInfos.HouseIDs(req.HouseID)[0], nil
}

func (i *AO) RegisterV2(req *dao.RegisterV2Req) (string, error) {
	if req.UserName == "" || req.EmailAddress == "" || req.Password == "" {
		return "请填写完整", nil
	}
	if req.VerificationCode != dao.UserID2VerifyCode[req.EmailAddress] {
		return "验证码错误", nil
	}
	dao.UserID2PwdMap[req.EmailAddress] = req.Password
	dao.UserID2NameMap[req.EmailAddress] = req.UserName
	return "OK", nil
}

func (i *AO) Verify(req *dao.VerifyReq) (string, error) {
	rand.Seed(time.Now().Unix())
	VerifyCode := fmt.Sprintf("na%v", rand.Intn(100000))

	dao.UserID2VerifyCode[req.EmailAddress] = VerifyCode

	m := gomail.NewMessage()
	m.SetHeader("From", config.C.MailAddress)
	m.SetHeader("To", req.EmailAddress)
	m.SetHeader("Subject", "验证码-租房平台")
	m.SetBody("text/html", fmt.Sprintf("您的验证码为：%v", VerifyCode))

	d := gomail.NewDialer(config.C.MailHost, int(config.C.MailPort), config.C.MailAddress, config.C.MailPwd)
	err := d.DialAndSend(m)
	if err != nil {
		log.Println("[Verify]", err)
		return err.Error(), err
	}
	return "OK", nil
}

func (i *AO) CollectionChange(req *dao.CollectionChangeReq) (string, error) {
	if req.SetOnline {
		dao.Collection[req.UserID] = append(dao.Collection[req.UserID], req.HouseID)
	} else {
		newCollection := make([]int32, 0)
		for _, hID := range dao.Collection[req.UserID] {
			if hID != req.HouseID {
				newCollection = append(newCollection, hID)
			}
		}
		dao.Collection[req.UserID] = newCollection
	}
	return "OK", nil
}

func (i *AO) CollectionSearch(req *dao.CollectionSearchReq) (*dao.CollectionSearchRsp, error) {
	rsp := &dao.CollectionSearchRsp{}
	houseIDs := dao.Collection[req.UserID]
	rsp.HouseInfos = dao.HouseInfos.IsOnline("online").HouseIDs(houseIDs...)
	rsp.Number = int32(len(houseIDs))
	rsp.Result = "OK"
	return rsp, nil
}

func (i *AO) ProfileAdd(req *dao.ProfileAddReq) (string, error) {
	// 维护先前已经添加进来的avatarURL
	req.Profile.AvatarURL = dao.Profiles[req.Profile.UserID].AvatarURL
	dao.Profiles[req.Profile.UserID] = req.Profile
	fmt.Println("[ProfileAdd]", "id", req.Profile.UserID, dao.Profiles[req.Profile.UserID])
	return "OK", nil
}

func (i *AO) ProfileGet(req *dao.ProfileGetReq) (*dao.Profile, error) {
	res := dao.Profiles[req.UserID]
	return &res, nil
}

func (i *AO) ProfileSearch(req *dao.ProfileSearchReq) (*dao.ProfileSearchRsp, error) {
	ps := make([]dao.Profile, 0)
	for _, p := range dao.Profiles {
		ps = append(ps, p)
	}
	rsp := &dao.ProfileSearchRsp{
		Result:   "OK",
		Number:   int32(len(dao.Profiles)),
		Profiles: ps,
	}
	return rsp, nil
}

func (i *AO) ChatList(req *dao.ChatListReq) ([]*dao.PersonInfo, error) {
	return chatManager.ChatList(req.UserID)
}
