package model

type OverallMetrics struct {
	Shipping		[2]float64		`json:"shipping"`
	Queueing		[2]float64		`json:"queueing"`
	Payments		[2]float64		`json:"payments"`
	Users			[2]float64		`json:"users"`
	Carts			[2]float64		`json:"carts"`
	Orders			[2]float64		`json:"orders"`
	Catalog			[2]float64		`json:"catalog"`

	MongoDB			[2]float64		`json:"mongo_db"`
	RabbitMQ		[2]float64		`json:"rabbit_mq"`
	ElasticSearch	[2]float64		`json:"elastic_search"`
	Kibana			[2]float64		`json:"kibana"`
}

type EsHits struct {
	Total int						`json:"total"`
}
type EsResp struct {
	Hits			EsHits			`json:"hits"`
}
