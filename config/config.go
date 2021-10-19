package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ImgPath string
	Addr    string
}

var C Config

func Init() error {
	useConfigFile := flag.String("c", "", "-c + 配置文件的文件名")
	flag.Parse()
	if *useConfigFile == "" {
		C = Config{
			ImgPath: "img",
			Addr:    "127.0.0.1:9000",
		}
	} else {
		initWithConfigFile(*useConfigFile)
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
