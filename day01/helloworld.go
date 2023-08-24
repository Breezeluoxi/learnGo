package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/goproxy"
)

func main() {
	// 定义命令行参数
	port := flag.String("port", "1314", "Proxy server port")
	// path := flag.String("path", "./", "work dir log out path")

	// 解析命令行参数
	flag.Parse()

	// 创建一个反向代理
	proxy := goproxy.NewProxyHttpServer()

	// 创建日志记录器
	logFile := &lumberjack.Logger{
		Filename: "proxy.log", // 日志文件名

		MaxSize:    100, // 单个日志文件最大尺寸，单位 MB
		MaxBackups: 3,   // 最多保留的旧日志文件数
		MaxAge:     7,   // 旧日志文件保留的天数
	}

	// 将日志输出重定向到日志记录器
	log.SetOutput(logFile)

	// 添加日志打印功能
	proxy.Verbose = true

	// 修改代理处理函数，记录客户端信息
	proxy.OnRequest().DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		// 打印请求信息
		log.Printf("Incoming Request from: %s - %s %s", ctx.Req.RemoteAddr, r.Method, r.URL.Path)

		// 返回 nil 表示不修改请求和响应
		return r, nil
	})

	// 启动代理服务器
	addr := ":" + strings.TrimLeft(*port, ":") // 确保端口前面有冒号
	log.Printf("Proxy server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
