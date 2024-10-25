package main

import (
	"errors"
	"html/template"
	"net/http"

	"codeclips.tanvirrifat.io/internal/models"
	"codeclips.tanvirrifat.io/internal/validator"
)

type CreateCodeClips struct{
  Title string `form:"title"`
	Language string `form:"language"`
	Content template.HTML `form:"content"`
    // ignore this when parsing using the go-playground/form decoder package
  validator.Validator `form:"-"`


}

type UserSignupForm struct{
    Name string `form:"name"`
    Email string `form:"email"`
    Password string `form:"password"`
    validator.Validator `form:"-"`
}

type UserLoginForm struct{
    Email string `form:"email"`
    Password string `form:"password"`
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


func (app *App) contact(w http.ResponseWriter, r *http.Request){
	data:= app.newTemplateData(r)

	app.render(w,r,http.StatusOK,"contact.tmpl.html",data)
}


func (app *App) signup(w http.ResponseWriter, r *http.Request){
	data:= app.newTemplateData(r)
	data.Form = UserSignupForm{
		Name: "",
		Email: "",
		Password:"",
	}

	app.render(w,r,http.StatusOK,"signup.tmpl.html",data)
}



func (app *App) signupPost(w http.ResponseWriter, r *http.Request){

	var form UserSignupForm

	err:=app.decodePostForm(w,r,&form)

	if err!=nil{
		app.clientError(w,http.StatusBadRequest)
		return
	}
// validation check:
	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
    form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
    form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
    form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
    form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

		// if there is any error:
		 if !form.Valid() {
        data := app.newTemplateData(r)
        data.Form = form
        app.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl.html", data)
        return
    }

		err = app.users.Insert(form.Name,form.Email,form.Password)

		// if duplicate email appear's
		if err!=nil{
			if errors.Is(err,models.ErrDuplicateEmail){
				 form.AddFieldError("email", "Email address is already in use")

            data := app.newTemplateData(r)
            data.Form = form
            app.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl.html", data)
			} else{
				app.serverError(w,r,err)
			}

			return
		}

		app.sessionManager.Put(r.Context(), "toast", "Your signup was successful. Please log in.")

    
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}



func (app *App) login(w http.ResponseWriter,r *http.Request){
	data:= app.newTemplateData(r)

	data.Form= UserLoginForm{
		Email: "",
		Password: "",
	}

	app.render(w,r,http.StatusOK,"login.tmpl.html",data)

}