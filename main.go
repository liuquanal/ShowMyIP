package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func myIpHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	ip = ip[0:strings.LastIndex(ip, ":")]
	str := ip
	io.WriteString(w, str)
}

func main() {
	port := "12345" // 端口号
	ht := http.HandlerFunc(myIpHandler)
	if ht != nil {
		http.Handle("/myip", ht)
	}
	fmt.Println("service is running on port: " + port)
	fmt.Println("访问http://127.0.0.1:" + port + "/myip 查看您的IP地址")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
