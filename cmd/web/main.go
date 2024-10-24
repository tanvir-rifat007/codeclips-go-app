package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"codeclips.tanvirrifat.io/internal/models"
	"github.com/go-playground/form/v4"
	_ "github.com/lib/pq"
)


type App struct{
	logger *slog.Logger
	templateCache map[string]*template.Template
	formDecoder *form.Decoder

	codeClips *models.CodeClipsModel

}

func main(){

	logger:= slog.New(slog.NewTextHandler(os.Stdout,nil))

	// cli arguments
	addr:=flag.String("addr",":4000","HTTP network address")

	dsn:= flag.String("dsn","postgres://codeclips@localhost/codeclips?sslmode=disable","Postgres connection string")

	db,err:= openDB(*dsn)

	if err!=nil{
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	flag.Parse()

	templateCache,err:= newTemplateCache()

	if err!=nil{
		logger.Error(err.Error())
		os.Exit(1)
	}

   formDecoder:= form.NewDecoder()
	 codeClips:= &models.CodeClipsModel{
			DB: db,
		}

	// dependency injection
	app:= &App{
		logger: logger,
		templateCache:templateCache,
		formDecoder: formDecoder,
		codeClips: codeClips,

	}

  srv:= &http.Server{
		Addr: *addr,
		Handler: app.routes(),
	}

	logger.Info(fmt.Sprintf("Starting server on %s",srv.Addr))

	err = srv.ListenAndServe()

	if err!=nil{
		logger.Error("Unable to start server", "error", err)
		os.Exit(1)
	}
}



func openDB(dsn string)(*sql.DB,error){
	db,err:=sql.Open("postgres",dsn)

	if err!=nil{
		return nil,err
	}

	if err:=db.Ping();err!=nil{
		return nil,err
	}

	return db,nil
}