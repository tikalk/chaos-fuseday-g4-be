package main

import (
	"net/http"
	"github.com/tikalk/chaos_fuze_2019/controllers"
)

func main(){
	http.HandleFunc("/metrics", controllers.GetMetrics)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

