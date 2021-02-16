package parser

type MetricItem struct {
	ResultType string                 `json:"resultType"`
	Metric     map[string]interface{} `json:"metric"`
	Value      []interface{}          `json:"value"`
}

type Metrics = *[]*MetricItem
