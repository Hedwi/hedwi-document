package middlewares

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/amalfra/etag"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"hedwi-document/config"
)

const defaultCacheControl = "public, max-age=31536000"

var StaticBox fs.FS

//var etagMap map[string]string
var etagMap sync.Map

func Init() {
	initEtags()
}

func initEtags() {

	//etagMap = map[string]string{}
	box, _ := fs.Sub(StaticBox, ".")
	fs.WalkDir(box, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if path == "." {
			return nil
		}

		f, err := box.Open(path)
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(f)
		content := buf.String()

		fpath := fmt.Sprintf("/%s", strings.ToLower(path))
		if strings.HasSuffix(fpath, "index.html") {
			fpath = strings.Replace(fpath, "index.html", "", 1) // trim index.html
		}

		tag := etag.Generate(content, true)

		//etagMap[fpath] = tag
		etagMap.Store(fpath, tag)
		return nil
	})
}

func initTemplateEtags() {
	//fmt.Println("render.Render")
	//fmt.Println(render.Render)

	arr := [][]string{
		[]string{"/", "home"},
	}

	for _, item := range arr {
		path := item[0]
		value := item[1]
		content := fmt.Sprintf("%s-%s-%s", path, value, config.DefaultConfig.VersionTimestamp)
		tag := etag.Generate(content, true)
		etagMap.Store(path, tag)

	}

}

func StackTrace(all bool) string {
	buf := make([]byte, 20480)
	number := runtime.Stack(buf, all)
	stack := string(buf[:number])
	return string(stack)
}



func GlobalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if err := recover(); err != nil {
			stack := StackTrace(false)

			sentry.CaptureMessage(stack)
			//logger.Info("异常处理, msg:%s 出错日志 => %s", stack, err)
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  err,
			})
			c.Abort()
		}
	}(c)
	c.Next()
}

func Cors(c *gin.Context) {
	/*
	   c.Header("Access-Control-Allow-Origin", "*")
	   c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
	   c.Header("Access-Control-Allow-Credentials", "true")
	*/
	//c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains");
	c.Next()
}

func setCacheHeader(c *gin.Context, path string) {
	if _tag, ok := etagMap.Load(path); ok {
		tag := _tag.(string)
		c.Header("Cache-Control", defaultCacheControl)
		c.Header("Etag", tag)
	}
}

func CacheControl(c *gin.Context) {
	path := strings.ToLower(c.Request.URL.Path)
	if strings.HasSuffix(path, "/") {
		setCacheHeader(c, path)
	} else {
		if strings.HasPrefix(path, "/static") {
			setCacheHeader(c, path)
		} else {
			var extension = filepath.Ext(path)
			switch extension {
			case ".css", ".js", ".png", ".jpeg", ".jpg", ".gif", ".pdf", ".ico", ".woff", ".woff2", ".svg":
				setCacheHeader(c, path)
			default:
			}
		}
	}
	c.Next()
}
