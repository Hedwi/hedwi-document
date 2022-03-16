package main

import (
	"embed"
	"fmt"
	"strconv"

	//"github.com/gin-contrib/pprof"
	//"github.com/gin-gonic/gin"

	"hedwi-document/config"
	"hedwi-document/route"
)

//go:embed static
//go:embed templates
var static embed.FS

func init() {
}

func main() {

	if err := config.InitConfigInfo(); err != nil {
		fmt.Println(err)
		panic("init config fail")
	}

	port := ":" + strconv.Itoa(config.DefaultConfig.Port)
	print(port)

	//route.StaticBox = static
	//r.Use(static.Serve("/", static.LocalFile("./build/html", true)))

	route.StaticBox = static
	route.TemplateBox = static

	r := route.InitRouter()

	r.Run(port)

}
