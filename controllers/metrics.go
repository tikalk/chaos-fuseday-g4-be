package controllers

import (
	"net/http"
	"github.com/tikalk/chaos_fuze_2019/model"
	"encoding/json"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	res := &model.OverallMetrics{}
	ser, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(ser)

}
