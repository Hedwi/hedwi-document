package staticHandler

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRealPath(c *gin.Context) string {
	id := c.Param("id")

	realPath := id
	_, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		realPath = "index.html"
	}

	return realPath
}

func GetMeetPath(c *gin.Context) string {

	id := c.Param("id")
	realPath := id
	if len(id) == 10 {
		realPath = "index.html"
	}
	if id == "" {
		realPath = "index.html"
	}
	return realPath
}

func HandleStatics(c *gin.Context, realPath string, box fs.FS) {

	ext := filepath.Ext(realPath)

	contentType := "text/plain; charset=utf-8"
	switch ext {
	case ".html":
		contentType = "text/html; charset=utf-8"
	case ".svg":
		contentType = "text/html; charset=utf-8"
	case ".js":
		contentType = "text/javascript; charset=utf-8"
	case ".css":
		contentType = "text/css; charset=utf-8"
	default:
	}

	//c.FileFromFS(realPath, http.FS(docsBox))

	f, err := box.Open(realPath)
	if err != nil {
		c.String(404, "page not found")
		return
	}

	iconData, err := ioutil.ReadAll(f)
	if err != nil {
		c.String(404, "page not found")
		return
	}

	c.Data(http.StatusOK, contentType, []byte(iconData))

}
