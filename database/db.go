package database

import (
	"database/sql"
	"fmt"
	"os"
)

func BlogInit(dbExists bool, dbPath string) *sql.DB {
	return initializeBlogTable(dbExists, "")
}

// assumes a db already exists and will return you an instance of a connection to it
func Connect(driver string, dbPath string) *sql.DB {
	db, err := sql.Open(driver, dbPath)
	if err != nil {
		panic(err)
	}

	return db
}

func AddBlog(blogName string, file string, commentTable string, db *sql.DB) {
	addBlog(blogName, file, commentTable, db)
}

func PrintBlogs(db *sql.DB) {
	rows := getBlogs(db)

	var id int
	var name string
	var content string
	var commentTable string

	for rows.Next() {
		rows.Scan(&id, &name, &content, &commentTable)
		fmt.Println(content)
		fmt.Println()
		fmt.Println(commentTable)
		fmt.Println()
	}
}

func IsDB(dbPath string) bool {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	} else {
		return false
	}
}
