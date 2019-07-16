package controllers

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/tikalk/chaos-fuseday-g4-be/model"
	"io/ioutil"
)

var es_url = "http://localhost:9200/_search"
var es_body = []byte(`{
    "query" : {
        "bool" : {
            "must" : [
                {"match" : { "log": "Posting Order" } },
                {
                    "range" : {
                        "@timestamp": {
                            "gte": "2019-07-16T07:48:45.579412449+00:00",
                            "lte": "2019-07-16T08:48:45.579412449+00:00"
                        }
                    }
                }
            ]
        }
    }
}`)
func GetMetrics(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Post(es_url, "application/json", bytes.NewBuffer(es_body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	esRes := &model.EsResp{}
	err = json.Unmarshal(body, &esRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(esRes.Hits.Total)

	res := &model.OverallMetrics{}
	ser, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(ser)

}
