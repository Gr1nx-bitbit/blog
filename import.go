package blog

import (
	"github.com/Gr1nx-bitbit/blog/database"
	"github.com/Gr1nx-bitbit/blog/server"
)

func Cool() {
	database.AddBlog("test", "./test.txt", "testComments")
	server.Greet()
}

func Greet() {
	server.Greet()
}
