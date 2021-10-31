package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"liansushe/ao"
	"liansushe/config"
	"liansushe/dao"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	configFileName := flag.String("c", "", "-c + 配置文件的文件名")
	flag.Parse()
	config.Init(*configFileName)

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

	// sprint1
	r.POST("/login", login)
	r.POST("/register", register)
	r.POST("/house/search", houseSearch)
	r.POST("/house/add", houseAdd)
	r.POST("/image/:uuid", uploadImage)
	r.GET("/image/:uuid", getImage)
	r.POST("/house/set_online", setOnline)
	r.POST("/house/set_offline", setOffline)

	// sprint2
	r.POST("/register/v2", registerV2)
	r.POST("/verify", verify)
	r.POST("/collection/change", collectionChange)
	r.POST("/collection/search", collectionSearch)
	r.GET("/vr/:imageName", vrHandler)
	r.Run(config.C.Addr)
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
		log.Println(string(bodyBytes))
		c.Status(http.StatusInternalServerError)
		return err
	}
	log.Printf("jsonString is %#v", p)
	return nil
}

func login(c *gin.Context) {
	req := dao.LoginReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	rsp, err := ao.Ao.Login(&req)
	if err != nil {
		log.Println("[login]", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result":   rsp.Result,
		"Token":    rsp.Token,
		"UserName": rsp.UserName,
	})
	log.Println("[login] success")
}

func register(c *gin.Context) {
	req := dao.RegisterReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	res, err := ao.Ao.Register(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": res,
	})
}

func houseSearch(c *gin.Context) {
	req := dao.HouseSearchReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	// houseInfo := dao.HouseInfo{
	// 	HouseID:   1,
	// 	ImgURL:    "./src/assets/logo.png",
	// 	VRURL:     "vurl",
	// 	Place:     "我家",
	// 	Center:    "环球港",
	// 	Area:      200,
	// 	Price:     5000,
	// 	Deposit:   9000,
	// 	Room:      2,
	// 	Hall:      5,
	// 	Elevator:  true,
	// 	Storey:    3,
	// 	Term:      6,
	// 	Direction: "东南",
	// 	Facility:  15,
	// 	Note:      "test",
	// 	IsOnline:  true,
	// }
	// for i := 0; i < req.PageSize; i++ {
	// 	houseInfo.HouseID++
	// 	houseInfo.Price = 1000*req.PageNum + i
	// 	houseInfos = append(houseInfos, houseInfo)
	// }
	res, err := ao.Ao.HouseSearch(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result":     "OK",
		"Number":     len(res),
		"HouseInfos": res,
	})
}

func houseAdd(c *gin.Context) {
	req := dao.HouseAddReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}

	res, err := ao.Ao.HouseAdd(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": res,
	})
}

func uploadImage(c *gin.Context) {
	uuid := c.Param("uuid")
	f, err := c.FormFile("file")
	log.Println(c.ContentType(), f.Filename, f.Header, f.Size)
	if err != nil || f.Filename == "" {
		log.Println("[image]", err, " or fileName is empty")
		c.Status(http.StatusInternalServerError)
		return
	}

	// get path
	fNames := strings.Split(f.Filename, ".")
	fName := uuid + "." + fNames[len(fNames)-1]
	path := filepath.Join(config.C.ImgPath, fName)
	dao.UUID2ImagePath[uuid] = path
	log.Println(path)

	os.MkdirAll(config.C.ImgPath, os.FileMode(0777))
	c.SaveUploadedFile(f, path)
	c.Status(http.StatusOK)
}

func getImage(c *gin.Context) {
	uuid := c.Param("uuid")
	fPath := dao.UUID2ImagePath[uuid]
	c.File(fPath)
}

func setOnline(c *gin.Context) {
	req := dao.SetOnlineReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}

	res, err := ao.Ao.SetOnline(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": res,
	})
}

func setOffline(c *gin.Context) {
	req := dao.SetOfflineReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}

	res, err := ao.Ao.SetOffline(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": res,
	})
}

func registerV2(c *gin.Context) {
	req := dao.RegisterV2Req{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	res, err := ao.Ao.RegisterV2(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": res,
	})
}

func verify(c *gin.Context) {
	req := dao.VerifyReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	res, err := ao.Ao.Verify(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": res,
	})
}

func collectionChange(c *gin.Context) {
	req := dao.CollectionChangeReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	res, err := ao.Ao.CollectionChange(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result": res,
	})
}

func collectionSearch(c *gin.Context) {
	req := dao.CollectionSearchReq{}
	err := parseReq(c, &req)
	if err != nil {
		return
	}
	res, err := ao.Ao.CollectionSearch(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Result":     res.Result,
		"Number":     res.Number,
		"HouseInfos": res.HouseInfos,
	})
}

func vrHandler(c *gin.Context) {
	imageName := c.Param("imageName")
	filePath := filepath.Join("vr", imageName)
	c.File(filePath)
}
