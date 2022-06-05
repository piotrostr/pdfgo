package server

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

func FilterPdfs(files []*multipart.FileHeader) []io.ReadSeeker {
	pdfs := []io.ReadSeeker{}
	for _, file := range files {
		if filepath.Ext(file.Filename) == ".pdf" {
			f, err := file.Open()
			if err != nil {
				continue
			}
			defer f.Close()

			buf := bytes.NewBuffer(nil)
			_, err = io.Copy(buf, f)
			if err != nil {
				continue
			}
			rs := bytes.NewReader(buf.Bytes())
			if err == nil {
				pdfs = append(pdfs, rs)
			}
		}
	}
	return pdfs
}

func SetupRouter() (r *gin.Engine, err error) {
	gin.SetMode(gin.ReleaseMode)

	r = gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK\n")
	})

	r.POST("/merge", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]

		pdfs := FilterPdfs(files)
		if len(pdfs) == 0 {
			ctx.String(http.StatusBadRequest, "no pdf files uploaded\n")
			return
		}
		var merged bytes.Buffer
		pdf.Merge(pdfs, &merged, nil)
		ctx.Data(http.StatusOK, "application/pdf", merged.Bytes())
	})

	return
}
