package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
)


type App struct{
	logger *slog.Logger
	templateCache map[string]*template.Template

}

func main(){

	logger:= slog.New(slog.NewTextHandler(os.Stdout,nil))

	// cli arguments
	addr:=flag.String("addr",":4000","HTTP network address")

	flag.Parse()

	templateCache,err:= newTemplateCache()

	if err!=nil{
		logger.Error(err.Error())
		os.Exit(1)
	}

	// dependency injection
	app:= &App{
		logger: logger,
		templateCache:templateCache,
	}

	mux:= http.NewServeMux();

	// static file:

	fileServer:= http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/",http.StripPrefix("/static",fileServer))


	mux.HandleFunc("/",app.home)

	mux.HandleFunc("/clips",app.clips)

	logger.Info(fmt.Sprintf("Starting server on %s",*addr))

	err = http.ListenAndServe(*addr, mux)

	if err!=nil{
		logger.Error("Unable to start server", "error", err)
		os.Exit(1)
	}




	


}