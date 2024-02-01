package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func res2string(res *http.Response, err error) string {
	if err != nil {
		fmt.Println(err)
		return ""
	}
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(d)
}

func get(url string) string {
	res, err := http.Get(url)
	return res2string(res, err)
}

func do(uri string) string {
	client := &http.Client{}

	u := &url.URL{
		Scheme: "http",
		Host:   uri,
	}
	req := &http.Request{
		URL: u,
	}
	client.Do(req)

	return res2string(client.Do(req))
}

func main() {
	fmt.Println(get("https://www.baidu.com/"))
	fmt.Println("--------------------------------")

	fmt.Println(do("www.baidu.com"))
}
