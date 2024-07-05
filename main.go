package main

import (
	"github.com/Gr1nx-bitbit/blog/database"
	"github.com/Gr1nx-bitbit/blog/server"
)

func main() {
	db := database.BlogInit(true, "") // is this not working because I need to use ADD instead of INSERT?
	// database.AddBlog("test", "./test.txt", "testComments", db)
	database.PrintBlogs(db)
}

func Greet() {
	server.Greet()
}
