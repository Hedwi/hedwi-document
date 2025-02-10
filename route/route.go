package route

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"
	"strings"
	"time"

	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	//"github.com/markbates/pkger"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
	ginglog "github.com/szuecs/gin-glog"

	//"hedwi-document/render"

	"hedwi-document/internal/handlers/home"
	"hedwi-document/internal/handlers/mail"
	"hedwi-document/internal/handlers/meet"
	"hedwi-document/internal/handlers/send"
	"hedwi-document/internal/handlers/staticHandler"
	"hedwi-document/middlewares"
	"hedwi-document/render"
)

var StaticBox fs.FS
var TemplateBox embed.FS

func InitRouter() *gin.Engine {
	// 初始化路由
	flag.Parse()
	r := gin.Default()

	//store := cookie.NewStore([]byte("secret"))

	handler := gzip.NewHandler(gzip.Config{
		// gzip compression level to use
		CompressionLevel: 2,
		// minimum content length to trigger gzip, the unit is in byte.
		MinContentLength: 2048,
		// RequestFilter decide whether or not to compress response judging by request.
		// Filters are applied in the sequence here.
		RequestFilter: []gzip.RequestFilter{
			gzip.NewCommonRequestFilter(),
			gzip.DefaultExtensionFilter(),
		},
		// ResponseHeaderFilter decide whether or not to compress response
		// judging by response header
		ResponseHeaderFilter: []gzip.ResponseHeaderFilter{
			gzip.DefaultContentTypeFilter(),
		},
	})

	// 使用Sentry
	// r.Use(sentry.Recovery(raven.DefaultClient, false))
	// r.Use(middlewares.GlobalRecover)
	//r.Use(sentry.Recovery(raven.DefaultClient, false))

	r.Use(sentrygin.New(sentrygin.Options{}))
	r.Use(ginglog.Logger(10 * time.Second))
	//r.Use(middlewares.Cors)
	r.Use(middlewares.CacheControl)
	//r.Use(middlewares.JsonP())
	r.Use(handler.Gin)
	//r.Use(sessions.Sessions("mysession", store))

	//render.InitRender()
	//r.HTMLRender = &render.Render

	staticBox, _ := fs.Sub(StaticBox, "static")

	render.TemplateBox = TemplateBox
	render.InitRender()
	r.HTMLRender = &render.Render

	r.GET("/", home.Home)
	r.GET("/meet/", meet.Locale)
	r.GET("/meet", meet.Locale)
	r.GET("/document/meet", meet.Locale)
	r.GET("/document/meet/", meet.Locale)
	r.GET("/document/mail-suite", mail.Locale)
	r.GET("/document/mail-suite/", mail.Locale)
	r.GET("/document/send", send.Locale)
	r.GET("/document/send/", send.Locale)

	wellknownBox, _ := fs.Sub(StaticBox, "wellknown")
	r.StaticFS("/.well-known", http.FS(wellknownBox))

	locales := []string{"zh-hans", "en-us"}
	for _, locale := range locales {
		sendBox, _ := fs.Sub(StaticBox, "static/hedwi-api/"+locale)
		r.StaticFS("/document/api/"+locale, http.FS(sendBox))
		r.StaticFS("/api/"+locale, http.FS(sendBox))
	}

	docLocales := []string{"zh-hans", "en-us"}
	for _, locale := range docLocales {

		mailSuiteBox, _ := fs.Sub(StaticBox, "document/hedwi-mail-suite/"+locale)
		meetBox, _ := fs.Sub(StaticBox, "document/hedwi-meet/"+locale)
		apiBox, _ := fs.Sub(StaticBox, "document/hedwi-api/"+locale)

		r.GET("/document/mail-suite/"+locale+"/*path", func(c *gin.Context) {
			realPath := c.Param("path")
			realPath = strings.TrimPrefix(realPath, "/")
			realPath = strings.Replace(realPath, "_static", "static", 1)
			if realPath == "" {
				realPath = "index.html"
			}
			staticHandler.HandleStatics(c, realPath, mailSuiteBox)
		})

		r.GET("/document/meet/"+locale+"/*path", func(c *gin.Context) {
			realPath := c.Param("path")
			realPath = strings.TrimPrefix(realPath, "/")
			realPath = strings.Replace(realPath, "_static", "static", 1)
			if realPath == "" {
				realPath = "index.html"
			}
			staticHandler.HandleStatics(c, realPath, meetBox)
		})

		r.GET("/document/send/"+locale+"/*path", func(c *gin.Context) {
			realPath := c.Param("path")
			realPath = strings.TrimPrefix(realPath, "/")
			realPath = strings.Replace(realPath, "_static", "static", 1)
			if realPath == "" {
				realPath = "index.html"
			}
			staticHandler.HandleStatics(c, realPath, apiBox)
		})

	}

	r.StaticFS("/static", http.FS(staticBox))

	return r
}
