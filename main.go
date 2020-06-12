package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Workshop struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LogoUrl     string `json:"preview_url"`
	Level       int    `json:"level"`
}

func main() {
	resp, _ := http.Get("https://api.codemao.cn/web/shops/" + os.Args[1])
	defer resp.Body.Close()
	pageBytes, _ := ioutil.ReadAll(resp.Body)
	var workshop Workshop
	json.Unmarshal(pageBytes, &workshop)

	fmt.Printf("工作室名称：%s\n", workshop.Name)
	fmt.Printf("工作室介绍：%s\n", workshop.Description)
	fmt.Printf("工作室等级：%d\n", workshop.Level)
	fmt.Printf("工作室 Logo：%s\n", workshop.LogoUrl)
}
