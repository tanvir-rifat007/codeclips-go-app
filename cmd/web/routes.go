package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *App) routes()http.Handler{

	// this is for session management
	dynamic:= alice.New(app.sessionManager.LoadAndSave)

	mux:= http.NewServeMux()

	fileServer:= http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/",http.StripPrefix("/static",fileServer))


	mux.Handle("GET /{$}",dynamic.ThenFunc(app.home))
	mux.Handle("POST /{$}",dynamic.ThenFunc(app.codeClipsPost))


	mux.Handle("GET /clips",dynamic.ThenFunc(app.clips))

	standard:= alice.New(app.logRequest,app.commonHeader)


	return standard.Then(mux)



}