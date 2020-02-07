package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func myIpHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	ip = ip[0:strings.LastIndex(ip, ":")]
	str := ip
	io.WriteString(w, str)
}

func main() {
	version := "1.0.0"
	port := "12345" // 端口号
	ht := http.HandlerFunc(myIpHandler)
	if ht != nil {
		http.Handle("/myip", ht)
	}
	fmt.Println("version: " + version)
	fmt.Println("service is running on port: " + port)
	fmt.Println("访问http://127.0.0.1:" + port + "/myip 查看您的IP地址")
	fmt.Println("启动于：" + time.Now().Format("2006-01-02 15:04:05"))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
