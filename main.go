package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Delims("{!{", "}!}")
	r.LoadHTMLFiles("./dist/index.html")
	// r.StaticFS("/", http.Dir("./dist/")) 这么写会和后面路由冲突
	r.StaticFS("/css", http.Dir("./dist/css"))
	r.StaticFS("/js", http.Dir("./dist/js"))
	r.StaticFS("/img", http.Dir("./dist/img"))
	r.StaticFS("/fonts", http.Dir("./dist/fonts"))
	r.StaticFS("/src", http.Dir("./src"))
	r.GET("/", getIndex)
	r.POST("/login", login)
	r.POST("/register", register)
	r.POST("/house/search", houseSearch)

	r.Run("127.0.0.1:9000")
}

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func parseReq(c *gin.Context, p interface{}) error {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[parseReq]", err)
		c.Status(http.StatusInternalServerError)
		return err
	}
	err = json.Unmarshal(bodyBytes, p)
	if err != nil {
		log.Println("[parseReq]", err)
		c.Status(http.StatusInternalServerError)
		return err
	}
	log.Printf("%#v", p)
	return nil
}

func login(c *gin.Context) {
	type Req struct {
		UserID   string
		Password string
	}
	req := Req{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	if req.UserID == "777" {
		c.JSON(http.StatusOK, gin.H{
			"Result":   "OK",
			"Token":    "token123",
			"UserName": "nanami",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Result": "该用户不存在",
		})
	}
	log.Println("[login] success")
}

func register(c *gin.Context) {
	type Req struct {
		UserName string
		UserID   string
		Password string
	}
	req := Req{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	if req.UserName == "" || req.UserID == "" || req.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"Result": "请填写完整",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Result": "OK",
		})
	}
}

func houseSearch(c *gin.Context) {
	type Req struct {
		PageSize int
		PageNum  int
		RoomNum  []int
		MinPrice int
		MaxPrice int
		Center   []string
		MinTerm  int
		MaxTerm  int
	}
	type Rsp struct {
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
	req := Req{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	houseInfo := Rsp{
		HouseID:   1,
		ImgURL:    "./src/assets/logo.png",
		VRURL:     "vurl",
		Place:     "我家",
		Center:    "环球港",
		Area:      200,
		Price:     5000,
		Deposit:   9000,
		Room:      2,
		Hall:      5,
		Elevator:  true,
		Storey:    3,
		Term:      6,
		Direction: "东南",
		Facility:  15,
		Note:      "test",
		IsOnline:  true,
	}
	houseInfos := make([]Rsp, 0)
	for i := 0; i < req.PageSize; i++ {
		houseInfo.HouseID++
		houseInfo.Price = 1000*req.PageNum + i
		houseInfos = append(houseInfos, houseInfo)
	}
	c.JSON(http.StatusOK, gin.H{
		"Result":     "OK",
		"Number":     50,
		"HouseInfos": houseInfos,
	})
}
