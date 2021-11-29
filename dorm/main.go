package main

import (
	"net/http"
	"time"

	"dorm/controller"
)

func main() {
	// 每10秒处理一次订单
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for _ = range ticker.C {
			m, _ := time.ParseDuration("-5s") //当前时间减5s，防止读写冲突
			timeStr := time.Now().Add(m).Format("2006-01-02 15:04:05")
			controller.ProcessOrder(timeStr)
		}
	}()

	http.HandleFunc("/choose", controller.ChooseDorm)
	http.HandleFunc("/getDormList", controller.GetDormList)
	http.HandleFunc("/getResult", controller.GetResult)
	http.ListenAndServe(":10713", nil)

}
