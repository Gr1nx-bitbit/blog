package database

import (
	"database/sql"
	"fmt"
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

func isDB(dbPath string) bool {
	if _, err := os.Stat(dbPath); err == nil {
		return true
	} else {
		return false
	}
}

func addBlog(blogName string, file string, commentTable string) { // I want to return an err
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		panic(err)
	}

	byteContent, err := os.ReadFile(file)
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

func dbPing() {
	db, _ := sql.Open("sqlite3", "./blog.db")
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	result, _ := db.Exec("CREATE TABLE IF NOT EXISTS Hello (id INTEGER PRIMARY KEY, First TEXT)")
	fmt.Printf("create table result: %v\n", result)

	statement, _ := db.Prepare("INSERT INTO Hello (First) VALUES (?)")
	result, _ = statement.Exec("Tommy")
	fmt.Println(result.RowsAffected())

	rows, _ := db.Query("SELECT * FROM Hello")
	var id int
	var name string
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println("id:", id, "name", name)
	}
}
