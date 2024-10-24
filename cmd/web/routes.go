package main

import "net/http"

func (app *App) routes()*http.ServeMux{

	mux:= http.NewServeMux()

	fileServer:= http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/",http.StripPrefix("/static",fileServer))


	mux.HandleFunc("GET /{$}",app.home)
	mux.HandleFunc("POST /{$}",app.codeClipsPost)


	mux.HandleFunc("GET /clips",app.clips)


	return mux



}