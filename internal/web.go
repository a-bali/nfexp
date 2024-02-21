package nfexp

import (
	"github.com/a-bali/nfexp/web"
	"io/fs"
	"log"
	"net/http"
)

func StaticFileServer() http.Handler {
	var staticFS = fs.FS(&web.Web)
	htmlContent, err := fs.Sub(staticFS, "build")
	if err != nil {
		log.Fatal(err)
	}
	return http.FileServer(http.FS(htmlContent))
}
