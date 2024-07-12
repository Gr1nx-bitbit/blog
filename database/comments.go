package database

import (
	"database/sql"
)

// this function will be called whenever a new blog is made so
// it doesn't have to be exported. Just pass in a db instance
// and the name of the blog. remember, the blog name has to be
// UNIQUE since that will be the name of the table!
func createCommentTable(db *sql.DB, blogName string) {
	// sqlite doesn't have a boolean keyword. It instead sotres 0's and 1's as falses and trues, respectively. It can however recognize the
	// TRUE and FALSE keyword so it'll automatically make those 0's and 1's.
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + blogName + "CommentTable" + " (\nCommentID INTEGER PRIMARY KEY,\n OwnerReference INTEGER,\n Leaf INTEGER,\n Comment TEXT\n)")
	if err != nil {
		panic(err)
	}
}
