package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ImgPath     string
	Addr        string
	MailAddress string
	MailPwd     string
	MailHost    string
	MailPort    int32
	AvatarPath  string
}

var C Config

func Init(configFileName string) error {
	if configFileName == "" {
		C = Config{
			ImgPath:     "img",
			Addr:        "127.0.0.1:9000",
			MailAddress: "1254312584@qq.com",
			MailPwd:     "ggbreadypudnhjja",
			MailHost:    "smtp.qq.com",
			MailPort:    465,
			AvatarPath:  "avatar",
		}
	} else {
		initWithConfigFile(configFileName)
	}
	log.Printf("%#v", C)
	return nil
}

func initWithConfigFile(configFile string) error {
	f, err := os.Open(configFile)
	if err != nil {
		log.Println("[config.Init] err", err)
		return err
	}
	defer f.Close()
	cBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("[config.Init] err", err)
		return err
	}
	err = json.Unmarshal(cBytes, &C)
	if err != nil {
		log.Println("[config.Init] err", err)
		return err
	}
	return nil
}
