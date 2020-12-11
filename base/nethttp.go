package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	//getDemo()
	postDemo()

}

func getDemo() {
	apiUrl := "http://www.baidu.com"
	data := url.Values{}
	data.Set("name","zs")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Println(err)
	}
	u.RawQuery = data.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func postDemo() {
	apiUrl := "http://www.baidu.com"
	contentType := "application/json"
	data := `{"name":"zs"}`
	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
