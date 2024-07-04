package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func createBlogTable() {
	db, _ := sql.Open("sqlite3", "./blog.db")
	db.Exec("CREATE TABLE IF NOT EXISTS Blogs (BlogID INTEGER, BlogName TEXT, BlogContent TEXT, CommentTable TEXT, PRIMARY KEY (BlogID));")
}

func getBlogs() {
	db, _ := sql.Open("sqlite3", "./blog.db")
	result, _ := db.Exec(".tables")
	fmt.Printf("result: %v\n", result)
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
