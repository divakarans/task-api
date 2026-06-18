package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)




func InitDB(path string)(*sql.DB,error) {
	
	database,err:=sql.Open("sqlite",path)
	if err!=nil{
		return nil,err
	}

	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT NOT NULL,
			created_at DATETIME NOT NULL
		);`

	_, err = database.Exec(query)
	if err != nil {
		return nil, err
	}

	return database,nil
	
}