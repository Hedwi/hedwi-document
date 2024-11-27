package mail

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hedwi-document/utils"
)

func Locale(c *gin.Context) {

	//now := time.Now()
	//ts := now.UnixMilli()

	v := c.Query("v")

	localeInCookie := utils.GetLocaleFromCookie(c)
	tag := utils.GetTag(localeInCookie, "")

	path := utils.GetLocalePath("mail-suite", tag, v)

	c.Redirect(http.StatusTemporaryRedirect, path)
	return
}
