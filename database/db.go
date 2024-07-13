package database

import (
	"database/sql"
	"fmt"
	"os"
)

// this function can be used to initialize a blog table in the database which
// there will ever only be ONE of. We can add some error handiling where it
// will check a specific path provided for a db instead of relying on the user.
// dbExists is self-explanatory as well as dbPath. pathReturn will return you
// the path the db has been set to so the rest of the program can use it
func BlogInit(dbExists bool, dbPath string) *sql.DB {
	return initializeBlogTable(dbExists, dbPath)
}

// assumes a db already exists and will return you an instance of a connection to it
func Connect(driver string, dbPath string) *sql.DB {
	db, err := sql.Open(driver, dbPath)
	if err != nil {
		panic(err)
	}

	return db
}

// this function will just add a new blog entry to the db. "blogName" is the name / title
// of the blog while "blogFilePath" is the path to the blog's content in the file system.
// db is an instance of the database since it will have already been created
func AddBlog(blogName string, file string, db *sql.DB) {
	addBlog(blogName, file, db)
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

// the owner ref is simply the ID of the parent comment however, how do we know
// if that comment is a root comment? We know we want it to have a value of -1
// if it doesn't have an owner, but how do we actually figure that out??
// Maybe the action of submitting a comment can call two different functions?
// If you submit a subcomment, it'll provide an owner-ref but if you create a root
// comment, it'll default the owner-ref to -1? It seems like we need help from the
// front end side, we can't just magically tell whether a comment is root or sub
// unless it's specified. Yeah, we just have to pass in that piece of data from the
// front end
func AddComment(tableName string, ownerRef int, commment string, db *sql.DB) {
	addComment(tableName, ownerRef, commment, db)
}

func IsDB(dbPath string) bool {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	} else {
		return false
	}
}
