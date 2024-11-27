package home

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"hedwi-document/config"
	"hedwi-document/utils"
)

func Home(c *gin.Context) {

	_timestamp := config.DefaultConfig.VersionTimestamp
	now := time.Now()
	year := now.Format("2006")
	locale := utils.GetLocale(utils.GetTag(utils.GetLocaleFromCookie(c), ""))

	c.HTML(http.StatusOK, "Home", gin.H{"_timestamp": _timestamp, "year": year, "locale": locale})
	return
}
