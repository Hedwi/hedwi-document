package route

import (
	"embed"
	"flag"
	"io/fs"
	"net/http"
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

	wellknownBox, _ := fs.Sub(StaticBox, "wellknown")
	r.StaticFS("/.well-known", http.FS(wellknownBox))

	mailBox, _ := fs.Sub(StaticBox, "static/hedwi-mail")
	docsBox, _ := fs.Sub(StaticBox, "static/hedwi-docs")
	sendBox, _ := fs.Sub(StaticBox, "static/hedwi-api")
	r.StaticFS("/mail", http.FS(mailBox))
	r.StaticFS("/send", http.FS(sendBox))
	r.StaticFS("/docs", http.FS(docsBox))
	r.StaticFS("/static", http.FS(staticBox))

	//r.StaticFS("/mail", http.Dir("./mail"))
	//r.StaticFS("/mail", pkger.Dir("/mail"))
	//r.StaticFS("/mail", http.Dir("../frontend/dist/hedwi-inbox/"))
	return r
}
