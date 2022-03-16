package home

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"hedwi-document/config"
)

func Home(c *gin.Context) {

	_timestamp := config.DefaultConfig.VersionTimestamp
	now := time.Now()
	year := now.Format("2006")
	c.HTML(http.StatusOK, "Home", gin.H{"_timestamp": _timestamp, "year": year})
	return
}
