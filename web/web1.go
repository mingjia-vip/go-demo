package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的

	fmt.Println("-----------1----------")
	fmt.Println("url:", r.URL)
	fmt.Println("url.path:", r.URL.Path)
	fmt.Println("method:", r.Method)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("model:",r.Form["model"])
	fmt.Println("-----------2----------")
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "@"))
	}
	fmt.Println("-----------3----------")
	fmt.Fprintf(w, "Hello, let's GO!") //这个写入到w的是输出到客户端的
	fmt.Println("")
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}