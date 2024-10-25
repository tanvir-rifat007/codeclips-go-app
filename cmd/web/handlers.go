package main

import (
	"html/template"
	"net/http"

	"codeclips.tanvirrifat.io/internal/validator"
)

type CreateCodeClips struct{
  Title string `form:"title"`
	Language string `form:"language"`
	Content template.HTML `form:"content"`
    // ignore this when parsing using the go-playground/form decoder package
  validator.Validator `form:"-"`


}


func (app *App) home(w http.ResponseWriter, r *http.Request){

  data:= app.newTemplateData(r)

	// for displaying the validator error in t he codeClipsPost handler
	data.Form = CreateCodeClips{
		Title: "",
		Language: "",
		Content: "",
	}

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

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
    form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
    form.CheckField(validator.NotBlank(string(form.Content)), "content", "This field cannot be blank")


    



        if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, r, http.StatusUnprocessableEntity, "home.tmpl.html", data)
        return
    }


    err = app.codeClips.Insert(form.Title,form.Language,form.Content)


    if err!=nil{
        app.serverError(w,r,err)
        return
    }

	// toast message:

	app.sessionManager.Put(r.Context(),"toast","Code clips saved successfully!")

	http.Redirect(w,r,"/clips",http.StatusSeeOther)

}


func (app *App) clips(w http.ResponseWriter, r *http.Request){
	codeClips,err:= app.codeClips.GetAll()

	if err!=nil{
		app.serverError(w,r,err)
		return
	}




	data:= app.newTemplateData(r)
	data.CodeClips = codeClips

	app.render(w,r,http.StatusOK,"clips.tmpl.html",data)


}