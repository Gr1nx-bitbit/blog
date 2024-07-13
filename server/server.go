package server

import (
	"net/http"

	"github.com/Gr1nx-bitbit/blog/database"
	"github.com/gin-gonic/gin"
)

/*
well, I called the package server but that doesn't mean much
for the blog I still have to come up with the design, how I
want it to look and what it'll be. honestly, the blog should be
an add on to just what my website is! My website should just
be a collection of the projects I've done and am trying to do.
This blog thing will just be something I can let thoughts out on
*/

/*
now that we have some basic functionality with the database, let's try and connect
it to a frontend. We're just looking for something extremely simple. A container to
hold the blog content and then a container to hold the comments. Each comment can
have a child, I'm guessing that's a div with a paragraph in there, eh, that should work
*/

/*
For creating dynamic routes (we'll be uploading new blogs which will each have its own route –
I should add that as a field in the Blogs table – but I don't want to restart the service every
time we upload something new) I think we just use a for loop or a range and just do a GET or POST
on that route specified in the DB – seems like that's another thing I have to add to the Blogs table.
*/

type Article struct {
	Blob     database.Blog
	Comments []database.Comment
}

func test(c *gin.Context) {
	content := database.Blog{
		Title:   "test",
		Content: getBlogContent("test"),
	}
	c.HTML(http.StatusOK, "blog-layout.html", content)
}

func getBlogContent(blogName string) string {
	db := database.Connect("sqlite3", "./blog.db")
	rows := database.GetBlog(blogName, db)

	var content string
	if rows != nil {
		for rows.Next() {
			rows.Scan(&content)
		}
	}

	return content
}

func Serve() {
	router := gin.Default()
	router.LoadHTMLGlob("front-end/templates/*")
	router.GET("/", test)
	router.Run(":4242")
}
