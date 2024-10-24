package models

import (
	"database/sql"
	"html/template"
)


type CodeClips struct{
	Title string
	Language string
	Content template.HTML
}


type CodeClipsModel struct{
	DB *sql.DB
}

func (c *CodeClipsModel) Insert(title,language string,content template.HTML)error{
	stmt:=`INSERT INTO codeclips(title,language,content) VALUES($1,$2,$3)`
	_,err:=c.DB.Exec(stmt,title,language,content)
	if err!=nil{
		return err
	}
	return nil

}