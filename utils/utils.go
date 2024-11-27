package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"

	"hedwi-document/internal/handlers/consts"
)

func GetLocaleFromCookie(c *gin.Context) string {

	locale := ""

	localeCookie, err := c.Cookie(consts.LocaleCookieKey)
	if err == nil {
		locale = localeCookie
		//return fmt.Sprintf("%s%s", name, locale)
	}
	return locale
}

func GetLocale(tag language.Tag) string {

	locale := "zh-hans"
	switch strings.ToLower(tag.String()) {
	case "zh_cn", "zh-tw", "zh-hk", "zh-hans":
		locale = "zh-hans"
	case "en-us":
		locale = "en-us"
	default:
	}
	return locale
}

func GetLocalePath(product string, tag language.Tag, v string) string {
	locale := "zh-hans"
	path := fmt.Sprintf("/%s/%s/?v=%s", product, locale, v)
	switch strings.ToLower(tag.String()) {
	case "zh_cn", "zh-tw", "zh-hk", "zh-hans":
		locale = "zh-hans"
		path = fmt.Sprintf("/%s/%s/?v=%s", product, locale, v)
	case "en-us":
		locale = "en-us"
		path = fmt.Sprintf("/%s/%s/?v=%s", product, locale, v)
	default:
	}
	return path
}

func GetLocalePathWithId(product string, tag language.Tag, v string, id string) string {
	locale := "zh-hans"
	path := fmt.Sprintf("/%s/%s/%s?v=%s", product, locale, id, v)
	switch strings.ToLower(tag.String()) {
	case "zh_cn", "zh-tw", "zh-hk", "zh-hans":
		locale = "zh-hans"
		path = fmt.Sprintf("/%s/%s/%s?v=%s", product, locale, id, v)
	case "en-us":
		locale = "en-us"
		path = fmt.Sprintf("/%s/%s/%s?v=%s", product, locale, id, v)
	default:
	}
	return path
}
