package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *App) routes()http.Handler{

	// this is for session management
	dynamic:= alice.New(app.sessionManager.LoadAndSave,app.authenticate)

	// protected routes:
	protected:= dynamic.Append(app.requireAuthentication)


	mux:= http.NewServeMux()

	fileServer:= http.FileServer(http.Dir("./ui/static"))

	// this fileserver for service worker offline capability

		mux.Handle("/", http.FileServer(http.Dir(".")))


	mux.Handle("/static/",http.StripPrefix("/static",fileServer))



	mux.Handle("GET /{$}",dynamic.ThenFunc(app.home))
	mux.Handle("GET /signup",dynamic.ThenFunc(app.signup))

	mux.Handle("POST /signup",dynamic.ThenFunc(app.signupPost))

	mux.Handle("GET /login",dynamic.ThenFunc(app.login))

	mux.Handle("POST /login",dynamic.ThenFunc(app.loginPost))

		mux.Handle("GET /contact",dynamic.ThenFunc(app.contact))

	// authentication chara clips post korte parbo nah
	mux.Handle("POST /{$}",protected.ThenFunc(app.codeClipsPost))

	// authentication chara clips dekhteo  parbo nah

	mux.Handle("GET /clips",protected.ThenFunc(app.clips))

	// authentication chara logout o korte parbo nah

	mux.Handle("GET /logout",protected.ThenFunc(app.logoutPost))





	standard:= alice.New(app.logRequest,app.commonHeader)


	return standard.Then(mux)



}