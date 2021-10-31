package ao

import (
	"liansushe/config"
	"liansushe/dao"
	"log"
	"testing"
)

func init() {
	config.Init("")
}

func TestVerify(t *testing.T) {
	req := dao.VerifyReq{
		EmailAddress: "18307130045@fudan.edu.cn",
	}
	res, err := Ao.Verify(&req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
