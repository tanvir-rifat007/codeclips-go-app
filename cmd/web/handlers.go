package main

import (
	"net/http"
)


func (app *App) home(w http.ResponseWriter, r *http.Request){

  data:= app.newTemplateData(r)

	app.render(w,r,http.StatusOK,"home.tmpl.html",data)

	
}


func (app *App) clips(w http.ResponseWriter, r *http.Request){
	data:= app.newTemplateData(r)

	app.render(w,r,http.StatusOK,"clips.tmpl.html",data)


}