package controllers

import (
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/tikalk/chaos-fuseday-g4-be/model"
	"io/ioutil"
	"fmt"
	"time"
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
                            "gte": "%s",
                            "lte": "%s"
                        }
                    }
                }
            ]
        }
    }
}`)
func GetMetrics(w http.ResponseWriter, r *http.Request) {

	url := fmt.Sprintf(string(es_body), time.Now().String(), time.Now().Add(time.Duration(-5) * time.Second).String())
	resp, err := http.Post(
		url,
		"application/json", bytes.NewBuffer(es_body))
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


	res := &model.OverallMetrics{ElasticSearch: [...]float64{0.0, float64(esRes.Hits.Total)}}
	ser, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(ser)

}
