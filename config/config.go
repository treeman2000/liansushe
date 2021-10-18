package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ImgPath string
}

var C Config

func Init() error {
	f, err := os.Open("config.json")
	if err != nil {
		log.Println("[config.Init] err", err)
		return err
	}
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
	log.Printf("%#v", C)
	return nil
}
