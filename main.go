package main

import (
	"github.com/Gr1nx-bitbit/blog/database"
	"github.com/Gr1nx-bitbit/blog/server"
)

func main() {
	database.AddBlog("test", "./test.txt", "testComments")
}

func Greet() {
	server.Greet()
}
