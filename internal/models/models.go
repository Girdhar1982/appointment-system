package models
// TemplateData holds data sent from template to  templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

type JsonAvailability struct{
	Ok bool `json:"ok"`
	Message string `json:"message"`
}