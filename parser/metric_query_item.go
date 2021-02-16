package parser

type MetricQueryItem struct {
	Query string `json:"query"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Step  int64  `json:"step"`
	Time  int64  `json:"time"`
	From  int64  `json:"from"`
}

type MetricsMap = map[string]*MetricQueryItem

type MetricQueryItemResult struct {
	Metric interface{} `json:"metric"`
	Value  interface{} `json:"value"`
}
