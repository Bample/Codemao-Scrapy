package main

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func middleWare(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	params := make([]string, 1)
	params[0] = "https://api.codemao.cn" // ["https://api.co..."]
	url := ""
	for _, v := range c.Params {
		params = append(params, v.Value)
	}
	url = strings.Join(params, "/") // "https://.../p1/p2/p3"
	var data map[string]interface{}
	json.Unmarshal(Scrap(url), &data)
	c.JSON(200, data)
}

// StartServer starts the API proxy server.
func StartServer(port string) {
	r := gin.New()
	r.GET("/:p1", middleWare)
	r.GET("/:p1/:p2", middleWare)
	r.GET("/:p1/:p2/:p3", middleWare)
	r.GET("/:p1/:p2/:p3/:p4", middleWare)
	r.GET("/:p1/:p2/:p3/:p4/:p5", middleWare)
	r.GET("/:p1/:p2/:p3/:p4/:p5/:p6", middleWare)
	r.GET("/:p1/:p2/:p3/:p4/:p5/:p6/:p7", middleWare)
	r.GET("/:p1/:p2/:p3/:p4/:p5/:p6/:p7/:p8", middleWare)
	r.Run(":" + port)
}
