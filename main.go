package main

import (
	"blog/database"
	"blog/server"
)

func main() {
	database.AddBlog("test", "./test.txt", "testComments")
	server.Greet()
}
