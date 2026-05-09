package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 处理请求
func handle(resp http.ResponseWriter, req *http.Request) {

	// 获取参数
	fmt.Println(req.URL.String())

	var names = []string{"Tom", "Jack", "张三"}
	result := Response{Code: 200, Msg: "处理成功", Data: names}

	byteData, _ := json.Marshal(result)
	resp.Write(byteData)

}

func main() {

	http.HandleFunc("/", handle)

	fmt.Println("http server 已经启动")
	http.ListenAndServe(":8080", nil)
}
