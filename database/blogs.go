package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// this function can be used to initialize a blog table in the database which
// there will ever only be ONE of. We can add some error handiling where it
// will check a specific path provided for a db instead of relying on the user.
// dbExists is self-explanatory as well as dbPath. pathReturn will return you
// the path the db has been set to so the rest of the program can use it
func initializeBlogTable(dbExists bool, dbPath string) (databaseInstance *sql.DB) {
	if !dbExists {
		if dbPath == "" {
			os.Create("blog.db")
			db, err := sql.Open("sqlite3", "./blog.db")
			if err != nil {
				panic(err)
			}

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (BlogID INTEGER PRIMARY KEY, BlogName TEXT, BlogContent TEXT, CommentTable TEXT)")
			if err != nil {
				panic(err)
			}

			return db
		} else {
			os.Create(dbPath)
			db, err := sql.Open("sqlite3", dbPath)
			if err != nil {
				panic(err)
			}

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (BlogID INTEGER PRIMARY KEY, BlogName TEXT, BlogContent TEXT, CommentTable TEXT)")
			if err != nil {
				panic(err)
			}

			return db
		}

	} else {
		if dbPath == "" {
			db, err := sql.Open("sqlite3", "./blog.db")
			if err != nil {
				panic(err)
			}

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (BlogID INTEGER PRIMARY KEY, BlogName TEXT, BlogContent TEXT, CommentTable TEXT)")
			if err != nil {
				panic(err)
			}

			return db
		} else {
			db, err := sql.Open("sqlite3", dbPath)
			if err != nil {
				panic(err)
			}

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (BlogID INTEGER PRIMARY KEY, BlogName TEXT, BlogContent TEXT, CommentTable TEXT)")
			if err != nil {
				panic(err)
			}

			return db
		}
	}
}

// this function will just add a new blog entry to the db. "blogName" is the name / title
// of the blog while "blogFilePath" is the path to the blog's content in the file system
// and "commentTable" is the unique name of the specific blog's table for comments. db is
// an instance of the database since it will have already been created
func addBlog(blogName string, blogFilePath string, commentTable string, db *sql.DB) { // I want to return an err

	byteContent, err := os.ReadFile(blogFilePath)
	if err != nil {
		panic(err)
	}

	content := string(byteContent)

	statement, err := db.Prepare("INSERT INTO Blogs(BlogName, BlogContent, CommentTable) VALUES(?,?,?)")
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(blogName, content, commentTable)
	if err != nil {
		panic(err)
	}

}

func getBlogs(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM Blogs")
	if err != nil {
		panic(err)
	}

	return rows
}
