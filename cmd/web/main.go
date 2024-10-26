package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"codeclips.tanvirrifat.io/internal/models"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/lib/pq"
)


type App struct{
	logger *slog.Logger
	templateCache map[string]*template.Template
	formDecoder *form.Decoder

	codeClips *models.CodeClipsModel

	users *models.UserModel

	sessionManager *scs.SessionManager


}

func main(){

	logger:= slog.New(slog.NewTextHandler(os.Stdout,nil))

    port := os.Getenv("PORT")
    if port == "" {
        port = "4000"
    }


	// cli arguments
addr := flag.String("addr", ":" + port, "HTTP network address") 
  
	  DB_URL := os.Getenv("DB_URL")
    if DB_URL == "" {
        logger.Error("DB_URL environment variable is not set")
				
        os.Exit(1)
    }


	dsn:= flag.String("dsn",DB_URL,"Database URL")

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
		users:= &models.UserModel{
			DB: db,
		}



		sessionManager:= scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 *time.Hour

	// dependency injection
	app:= &App{
		logger: logger,
		templateCache:templateCache,
		formDecoder: formDecoder,
		codeClips: codeClips,
		users:users,
		sessionManager:sessionManager ,


		


	}



  srv:= &http.Server{
		Addr:*addr,
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