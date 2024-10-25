package main

import (
	"html/template"
	"path/filepath"
	"time"

	"codeclips.tanvirrifat.io/internal/models"
)


type TemplateData struct{
	CurrentYear int
	CodeClips []models.CodeClips
	Form any
	Toast string
}

// custom template
func humanDate(t time.Time) string{
	    return t.Format("02 Jan 2006 at 15:04")

}

var functions = template.FuncMap{
	"humanDate":humanDate,
}


func newTemplateCache() (map[string]*template.Template,error) {
	cache:= map[string]*template.Template{}

	// get the whole url for each pages:
	pages,err:=filepath.Glob("./ui/html/pages/*.tmpl.html")

	if err!=nil{
		return nil,err
	}

	for _,page:=range pages{
		// get the each pages name(like:clips.tmpl.html)
		 name:=filepath.Base(page)
// injecting the custom function template
		 ts,err:=template.New(name).Funcs(functions).ParseFiles(
			"./ui/html/base.tmpl.html",
			"./ui/html/pages/partials/nav.tmpl.html",
			page,
			"./ui/html/pages/partials/footer.tmpl.html",
		
		)

		if err!=nil{
			return nil,err
		}

		cache[name]=ts

	}

	return cache,nil

}

