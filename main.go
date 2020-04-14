package main

import (
	"fmt"
	"runtime"
	"strconv"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"hedwi-docs/config"
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

	router := gin.Default()
	router.Static("/", "./build/html")
	router.Run(port)
}
