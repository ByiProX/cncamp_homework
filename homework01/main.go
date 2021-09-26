/*
1⃣️  9.25课后作业
内容：编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把1都做完

1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回200
提交链接🔗：https://jinshuju.net/f/PlZ3xg
截止时间：10月7日晚23:59前
提示💡：
1、自行选择做作业的地址，只要提交的链接能让助教老师打开即可
2、自己所在的助教答疑群是几组，提交作业就选几组
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthzHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	response, _ := http.Get("http://localhost:8000")
	fmt.Fprintf(w, response.Status)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Add(k, strings.Join(v, " "))
	}

	VERSION := os.Getenv("VERSION")
	w.Header().Add("VERSION", VERSION)

	log.Printf("RemoteAddr = %s", r.RemoteAddr)

}

//!-handler
