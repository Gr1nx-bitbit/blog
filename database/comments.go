package database

import (
	"database/sql"
	"fmt"
)

type CommentTable struct {
}

// this function will be called whenever a new blog is made so
// it doesn't have to be exported. Just pass in a db instance
// and the name of the blog. remember, the blog name has to be
// UNIQUE since that will be the name of the table!
func createCommentTable(db *sql.DB, blogName string) {
	// sqlite doesn't have a boolean keyword. It instead sotres 0's and 1's as falses and trues, respectively. It can however recognize the
	// TRUE and FALSE keyword so it'll automatically make those 0's and 1's.
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + blogName + "CommentTable" + " (\nCommentID INTEGER PRIMARY KEY,\n OwnerReference INTEGER,\n Leaf INTEGER,\n HasChildren INTEGER,\n Comment TEXT\n)")
	if err != nil {
		panic(err)
	}
}

// if the comment is a root comment (at the highest level), the owner reference
// will be -1. "tableName" is the name of the comment table where the comment is
// to be put.
func addComment(tableName string, ownerRef int, comment string, db *sql.DB) {
	// we have to check whether the comment is a leaf
	rows, err := db.Query("SELECT HasChildren FROM " + tableName + " WHERE CommentID='" + string(ownerRef) + "'")
	if err != nil {
		panic(err)
	}

	var children int
	for rows.Next() {
		rows.Scan(&children)
	}

	if children == 1 {
		fmt.Println("Children Present")
	} else {
		fmt.Println("This comment is a first child and leaf.")
		_, err = db.Exec("UPDATE " + tableName + " SET HasChildren='TRUE', Leaf='FALSE' WHERE CommentID='" + string(ownerRef) + "'")
		if err != nil {
			panic(err)
		}
	}

	statement, err := db.Prepare("INSERT INTO " + tableName + "(OwnerReference, HasChildren, Leaf, Comment) VALUES(?,?,?,?)")
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(ownerRef, 0, 1, comment)
	if err != nil {
		panic(err)
	}

}

func getComments() CommentTable {
	return CommentTable{}
}
