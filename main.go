package main

import (
	"fmt"
	"runtime"
	"strconv"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"hedwi-document/config"
)

func init() {
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := config.InitConfigInfo(); err != nil {
		fmt.Println(err)
		panic("init config fail")
	}

	port := ":" + strconv.Itoa(config.DefaultConfig.Port)
	print(port)

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./build/html", true)))
	r.Run(port)
}
