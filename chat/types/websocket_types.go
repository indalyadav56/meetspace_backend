package types

type Payload struct {
	Event string `json:"event"`
	Data  map[string]interface{} `json:"data"`
}