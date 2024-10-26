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

func (app *App) clientError(w http.ResponseWriter, status int) {
    http.Error(w, http.StatusText(status), status)
}





// construct the TemplateData
// and Used in the template to show some data
func (app *App) newTemplateData(r *http.Request)TemplateData{

	return TemplateData{
		CurrentYear: time.Now().Year(),
		 Toast : app.sessionManager.PopString(r.Context(), "toast"),

		 IsAuthenticated: app.isAuthenticated(r),

		 


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


func (app *App) decodePostForm(w http.ResponseWriter, r *http.Request,dst any)error{
	err:=r.ParseForm()

	if err!=nil{
		return err
 	}

	err = app.formDecoder.Decode(dst,r.PostForm)

	if err!=nil{
		app.serverError(w,r,err)
		
	}
 return nil

}

func (app *App) isAuthenticated(r *http.Request)bool{
	  isAuthenticated,ok:=r.Context().Value(isAuthenticatedContextKey).(bool)

		if !ok{
			return false
		}

		return isAuthenticated
}


