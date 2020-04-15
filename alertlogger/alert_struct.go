package main

import "time"

// Format can be found here:
// https://prometheus.io/docs/alerting/configuration/#webhook_config

// Example
/*
{"receiver":"default-receiver","status":"firing","alerts":[{"status":"firing","labels":{"alertdomain":"exporter","alertname":"Exporter Down","instance":"flakyhost.com","job":"blackbox","priority":"low","severity":"warning"},"annotations":{"console":"Check the Grafana Dashboard at http://grafana:3000","description":"Exporter on flakyhost.com is not reachable.","summary":"Exporter blackbox is down!"},"startsAt":"2020-04-15T01:20:09.971335041Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://68550835a477:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1","fingerprint":"04e7a19784201ec7"},{"status":"firing","labels":{"alertdomain":"exporter","alertname":"Exporter Down","instance":"reliablehost.com","job":"blackbox","priority":"low","severity":"warning"},"annotations":{"console":"Check the Grafana Dashboard at http://grafana:3000","description":"Exporter on reliablehost.com is not reachable.","summary":"Exporter blackbox is down!"},"startsAt":"2020-04-15T01:20:09.971335041Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://68550835a477:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1","fingerprint":"cf074d009b94bf94"}],"groupLabels":{},"commonLabels":{"alertdomain":"exporter","alertname":"Exporter Down","job":"blackbox","priority":"low","severity":"warning"},"commonAnnotations":{"console":"Check the Grafana Dashboard at http://grafana:3000","summary":"Exporter blackbox is down!"},"externalURL":"http://d8f920d4bbc6:9093","version":"4","groupKey":"{}:{}"}
*/

// AlertManagerData is the webhook payload struct
type AlertManagerData struct {
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Alerts   []struct {
		Status string `json:"status"`
		Labels struct {
			Alertname string `json:"alertname"`
			Instance  string `json:"instance"`
			Job       string `json:"job"`
		} `json:"labels"`
		Annotations struct {
			Description string `json:"description"`
			Summary     string `json:"summary"`
		} `json:"annotations"`
		StartsAt     time.Time `json:"startsAt"`
		EndsAt       time.Time `json:"endsAt"`
		GeneratorURL string    `json:"generatorURL"`
	} `json:"alerts"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Instance  string `json:"instance"`
		Job       string `json:"job"`
	} `json:"commonLabels"`
	CommonAnnotations struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
	} `json:"commonAnnotations"`
	ExternalURL string `json:"externalURL"`
	Version     string `json:"version"`
	GroupKey    string `json:"groupKey"`
}
