package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)


func (app *App) serverError(w http.ResponseWriter, r *http.Request,err error){
	app.logger.Error(err.Error(),"method",r.Method,"uri",r.URL.RequestURI())
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)

}




// construct the TemplateData
// and Used in the template to show some data
func (app *App) newTemplateData(r *http.Request)TemplateData{

	return TemplateData{
		CurrentYear: time.Now().Year(),
	}

}


func (app *App) render(w http.ResponseWriter, r *http.Request, status int, page string, data TemplateData){
	 ts,ok:=app.templateCache[page]

	 if !ok{
        err := fmt.Errorf("the template %s does not exist", page)
		 app.serverError(w,r,err)
		 return
	 }

	 buf:= new(bytes.Buffer)

	 err:=ts.ExecuteTemplate(buf,"base",data)

	 if err!=nil{
		app.serverError(w,r,err)
		return
	 }

	 w.WriteHeader(status)

	 buf.WriteTo(w)



}