package database

import (
	"fmt"

	"github.com/google/uuid"
)

func AddBlog(blogName string, file string, commentTable string) {
	// addBlog(blogName, file, commentTable)
	createBlogTable()
	getBlogs()
}

func Rand() {
	id := uuid.New()
	fmt.Println(id.String())
}
