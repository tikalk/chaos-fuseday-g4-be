package main

import (
	"net/http"
	"github.com/tikalk/chaos-fuseday-g4-be/controllers"
)

func main(){
	http.HandleFunc("/metrics", controllers.GetMetrics)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

