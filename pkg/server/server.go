package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (r *gin.Engine, err error) {
	gin.SetMode(gin.ReleaseMode)

	r = gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK\n")
	})

	r.POST("/", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			fmt.Println(file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded\n", len(files)))
	})

	return
}
