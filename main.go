package main

import (
	"github.com/Gr1nx-bitbit/blog/database"
	// "github.com/Gr1nx-bitbit/blog/server"
)

func main() {
	db := database.BlogInit(database.IsDB("./blog"), "")
	database.AddBlog("test", "./test.txt", db)
	database.AddComment("testCommentTable", -1, "Does this work?", db)
	database.PrintBlogs(db)
}
