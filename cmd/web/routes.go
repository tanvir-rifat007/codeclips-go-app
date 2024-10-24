package main

import "net/http"

func (app *App) routes()*http.ServeMux{

	mux:= http.NewServeMux()

	fileServer:= http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/",http.StripPrefix("/static",fileServer))


	mux.HandleFunc("/",app.home)

	mux.HandleFunc("/clips",app.clips)


	return mux



}