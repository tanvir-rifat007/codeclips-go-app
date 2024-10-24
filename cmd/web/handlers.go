package main

import (
	"html/template"
	"net/http"
)

type CreateCodeClips struct{
  Title string `form:"title"`
	Language string `form:"language"`
	Content template.HTML `form:"content"`



}


func (app *App) home(w http.ResponseWriter, r *http.Request){

  data:= app.newTemplateData(r)

	app.render(w,r,http.StatusOK,"home.tmpl.html",data)

	
}

func (app *App) codeClipsPost(w http.ResponseWriter, r *http.Request){
	var form CreateCodeClips
	// decode form and update it to the struct
	err:=app.decodePostForm(w,r,&form)

	if err!=nil{
		app.clientError(w,http.StatusBadRequest)
		return
	}

	err = app.codeClips.Insert(form.Title,form.Language,form.Content)

	if err!=nil{
		 app.serverError(w,r,err)
        return
	}

}


func (app *App) clips(w http.ResponseWriter, r *http.Request){
	data:= app.newTemplateData(r)

	app.render(w,r,http.StatusOK,"clips.tmpl.html",data)


}