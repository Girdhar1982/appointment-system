package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/girdhar1982/appointment-system/config"
	"github.com/girdhar1982/appointment-system/pkg/handlers"
	"github.com/girdhar1982/appointment-system/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig; //global variable available to every part
var session *scs.SessionManager;
// each file must have main
func main() {
  app.InProduction =false
  session = scs.New()
  session.Lifetime= 24 * time.Hour;
  session.Cookie.Persist = true;
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
  app.Session = session;
 //get the TemplateCache for the app config
  tc, err := render.CreateTemplateCache();
	app.TemplateCache=tc;
	app.UseCache=false;
	repo:=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	if err != nil {log.Fatal(err)}
	fmt.Println("Starting Application on Port: ",portNumber)
	srv := &http.Server{Addr: portNumber,Handler: routes(&app)}
  err=srv.ListenAndServe()
	if err != nil {log.Fatal(err)}
}
