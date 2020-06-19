package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User type.
type User struct {
	Nickname    string  `json:"nickname"`
	Description string  `json:"description"`
	Doing       string  `json:"doing"`
	Level       float64 `json:"level"`
	Avatar      string  `json:"avatar"`
}

// GetUserInfo is a function.
func GetUserInfo(id string) User {
	resp, _ := http.Get("https://api.codemao.cn/api/user/info/detail/" + id)
	defer resp.Body.Close()
	pageBytes, _ := ioutil.ReadAll(resp.Body)
	data := make(map[string]interface{})
	json.Unmarshal(pageBytes, &data)
	userInfo := data["data"].(map[string]interface{})["userInfo"].(map[string]interface{})["user"].(map[string]interface{})
	user := User{
		Nickname:    userInfo["nickname"].(string),
		Description: userInfo["description"].(string),
		Doing:       userInfo["doing"].(string),
		Level:       userInfo["level"].(float64),
		Avatar:      userInfo["avatar"].(string),
	}
	return user
}

// NotifyUserInfo can print the info.
func NotifyUserInfo(u User) {
	fmt.Printf("Name:\t%s\n", u.Nickname)
	fmt.Println("Description: --------------------------------------")
	fmt.Printf("%s\n", u.Description)
	fmt.Println("---------------------------------------------------")
	fmt.Printf("Doing:\t%s\n", u.Doing)
	fmt.Printf("Level:\t%f\n", u.Level)
	fmt.Printf("Logo URL:\t%s\n", u.Avatar)
}
