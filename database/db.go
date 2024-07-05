package database

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func BlogInit(dbExists bool, dbPath string) *sql.DB {
	return initializeBlogTable(false, "")
}

// assumes a db already exists and will return you an instance of a connection to it
func Connect(driver string, dbPath string) *sql.DB {
	db, err := sql.Open(driver, dbPath)
	if err != nil {
		panic(err)
	}

	return db
}

func AddBlog(blogName string, file string, commentTable string) {
	// addBlog(blogName, file, commentTable)
	// createBlogTable()
	// getBlogs()
	dbPing()
}

func Rand() {
	id := uuid.New()
	fmt.Println(id.String())
}
