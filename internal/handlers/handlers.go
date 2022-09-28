package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/girdhar1982/appointment-system/config"
	"github.com/girdhar1982/appointment-system/internal/models"
	"github.com/girdhar1982/appointment-system/internal/render"
)

// Repo the repository used by the handlers
var Repo *Respository

// Repository is the repository type
type Respository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{App: a}
}

// New Handlers sets the repository for the handlers
func NewHandlers(r *Respository) {
	Repo = r
}

func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	removeIP:=r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",removeIP)
	render.RenderTemplates(w, r,"home.page.tmpl", &models.TemplateData{})
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there!!"

	removeIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = removeIP
	render.RenderTemplates(w, r,"about.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}

func (m *Respository) Generals(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there!!"

	removeIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = removeIP
	render.RenderTemplates(w,r, "generals.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}

func (m *Respository) Majors(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there!!"

	removeIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = removeIP
	render.RenderTemplates(w, r,"majors.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}

func (m *Respository) Reservations(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there!!"

	removeIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = removeIP
	render.RenderTemplates(w, r,"make-reservation.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}

func (m *Respository) Availability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there!!"

	removeIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = removeIP
	render.RenderTemplates(w, r,"search-availability.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}

func (m *Respository) PostAvailability(w http.ResponseWriter, r *http.Request) {
start:=r.Form.Get("start");
end:=r.Form.Get("end");
	stringMap := make(map[string]string)
	stringMap["start"] = start
	stringMap["end"] = end
	render.RenderTemplates(w,r, "search-availability.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}

func (m *Respository) JsonAvailability(w http.ResponseWriter, r *http.Request) {
reso :=models.JsonAvailability{
	Ok: true,
	Message: "Available!",
}
out,err :=json.MarshalIndent(reso,"","")
if err !=nil{
	log.Println(err)
}
w.Header().Set("Content-type","application/json")
w.Write(out)
	}

func (m *Respository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there!!"

	removeIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = removeIP
	render.RenderTemplates(w,r, "contact.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}