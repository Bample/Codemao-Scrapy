package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Scrap returns the page bytes.
func Scrap(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	pageBytes, _ := ioutil.ReadAll(resp.Body)
	return pageBytes
}
