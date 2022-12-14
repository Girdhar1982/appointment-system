package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/girdhar1982/appointment-system/config"
	"github.com/girdhar1982/appointment-system/internal/models"
	"github.com/justinas/nosurf"
)
var app *config.AppConfig
//NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig){
app=a
}

func AddDefaultData(td *models.TemplateData , r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// render template
func RenderTemplates(w http.ResponseWriter, r *http.Request,tmpl string,td *models.TemplateData) {
	var tc  map[string] *template.Template
	//get the template cache from the app config
	if app.UseCache {
	tc = app.TemplateCache;
}else{
	tc,_=CreateTemplateCache()
}
	//get requested template from Cache
  t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get templates")
	}

	buf := new(bytes.Buffer); //additional errro checking optional

	td = AddDefaultData(td,r)
	err := t.Execute(buf,td)
	if err != nil {
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}
	//get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {		return myCache, err	}
	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts
	}
	return myCache, nil
}
