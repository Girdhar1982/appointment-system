package handlers

import (
	"net/http"
	"github.com/girdhar1982/appointment-system/config"
	"github.com/girdhar1982/appointment-system/pkg/models"
	"github.com/girdhar1982/appointment-system/pkg/render"
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
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there!!"

	removeIP:=m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = removeIP
	render.RenderTemplates(w, "about.page.tmpl",&models.TemplateData{StringMap: stringMap,})
}
