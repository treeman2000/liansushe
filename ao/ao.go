package ao

import (
	"liansushe/dao"
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
	houseInfo := &dao.HouseInfo{
		HouseID:   req.HouseID,
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

	dao.HouseInfos = append(dao.HouseInfos, *houseInfo)
	return "OK", nil
}
