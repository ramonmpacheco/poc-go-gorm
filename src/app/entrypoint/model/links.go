package model

type link struct {
	Rel  string `json:"relationship"`
	Href string `json:"href"`
	Type string `json:"type"`
}
