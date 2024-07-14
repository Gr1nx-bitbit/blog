package main

import (
	"github.com/Gr1nx-bitbit/blog/database"
	"github.com/Gr1nx-bitbit/blog/server"
)

func cool() {
	db := database.BlogInit(database.IsDB("./blog.db"), "")
	database.AddBlog("test", "./test.txt", db)
	database.AddComment("testCommentTable", -1, "Does this work?", db)
	database.PrintBlogs(db)
}

func main() {
	server.Serve()
}
