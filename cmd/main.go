package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type AddReq struct {
	OperandA int
	OperandB int
}

type AddReply struct {
	Result int
}

func main() {
	// 处理发往/api/add的http请求
	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		var add_req AddReq
		// 将请求的body作为JSON字符串解码，并存入AddReq结构体内
		json.NewDecoder(r.Body).Decode(&add_req)
		// 打印请求数据
		log.Println("req: ", add_req.OperandA, add_req.OperandB)
		var add_reply AddReply
		// 进行加法计算，并保存结果到结构体内
		add_reply.Result = add_req.OperandA + add_req.OperandB
		// 将结果结构体进行JSON编码，并写入响应
		json.NewEncoder(w).Encode(add_reply)
	})

	// 使用web目录下的文件来响应对/web/路径的http请求，一般用作静态文件服务，例如html、javascript、css等
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("../web/"))))

	// 启动http服务
	log.Fatal(http.ListenAndServe(":8080", nil))

	// 可以使用如下curl命令来访问加法API
	// curl -s -XPOST -d'{"OperandA":10,"OperandB":20}' http://localhost:8080/api/add
}
