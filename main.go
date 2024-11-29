package main

import (
	"embed"
	"fmt"

	//"github.com/gin-contrib/pprof"
	//"github.com/gin-gonic/gin"

	"hedwi-document/config"
	"hedwi-document/middlewares"
	"hedwi-document/route"
)

//go:embed static
//go:embed templates
//go:embed document
var static embed.FS

func init() {
}

func main() {

	if err := config.InitWithEnv(); err != nil {
		fmt.Println(err)
		panic("init config fail")
	}

	addr := config.DefaultConfig.Addr

	//route.StaticBox = static
	//r.Use(static.Serve("/", static.LocalFile("./build/html", true)))

	middlewares.StaticBox = static
	middlewares.Init()

	route.StaticBox = static
	route.TemplateBox = static

	r := route.InitRouter()

	r.Run(addr)

}
