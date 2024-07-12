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

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (\nBlogID INTEGER PRIMARY KEY,\n BlogName TEXT,\n BlogContent TEXT,\n CommentTable TEXT\n)")
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

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (\nBlogID INTEGER PRIMARY KEY,\n BlogName TEXT,\n BlogContent TEXT,\n CommentTable TEXT\n)")
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

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (\nBlogID INTEGER PRIMARY KEY,\n BlogName TEXT,\n BlogContent TEXT,\n CommentTable TEXT\n)")
			if err != nil {
				panic(err)
			}

			return db
		} else {
			db, err := sql.Open("sqlite3", dbPath)
			if err != nil {
				panic(err)
			}

			_, err = db.Exec("CREATE TABLE IF NOT EXISTS Blogs (\nBlogID INTEGER PRIMARY KEY,\n BlogName TEXT,\n BlogContent TEXT,\n CommentTable TEXT\n)")
			if err != nil {
				panic(err)
			}

			return db
		}
	}
}

// this function will just add a new blog entry to the db. "blogName" is the name / title
// of the blog while "blogFilePath" is the path to the blog's content in the file system.
// db is an instance of the database since it will have already been created
func addBlog(blogName string, blogFilePath string, db *sql.DB) { // I want to return an err

	byteContent, err := os.ReadFile(blogFilePath)
	if err != nil {
		panic(err)
	}

	content := string(byteContent)

	statement, err := db.Prepare("INSERT INTO Blogs(BlogName, BlogContent, CommentTable) VALUES(?,?,?)")
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(blogName, content, blogName+"CommentTable")
	if err != nil {
		panic(err)
	}

	createCommentTable(db, blogName)

}

func getBlogs(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM Blogs")
	if err != nil {
		panic(err)
	}

	return rows
}
