package models

import (
	"database/sql"
	"html/template"
)


type CodeClips struct{
	Id int
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

func (c *CodeClipsModel) GetAll() ([]CodeClips, error) {
    // Prepare the SQL query to get all records
    stmt := `SELECT title, language, content FROM codeclips`

    rows, err := c.DB.Query(stmt)
    if err != nil {
        return nil, err
    }
    
    defer rows.Close()

    var codeClips []CodeClips

   
    for rows.Next() {
        var clip CodeClips
        err := rows.Scan(&clip.Title, &clip.Language, &clip.Content)
        if err != nil {
            return nil, err
        }
        // Append the result to the slice
        codeClips = append(codeClips, clip)
    }

    // Check if there were any errors during row iteration
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return codeClips, nil
}
