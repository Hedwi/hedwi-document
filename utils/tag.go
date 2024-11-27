package utils

import (
	"log/slog"

	"github.com/Xuanwo/go-locale"
	"golang.org/x/text/language"
)

func GetTag(localeInCookie string, localeName string) language.Tag {

	var tag language.Tag

	if localeName == "" {
		if localeInCookie != "" {
			localeName = localeInCookie
		}
	}

	if localeName != "" {
		t, err := language.Parse(localeName)
		if err == nil {
			tag = t
		}
	} else {
		//
	}

	if tag.String() == "" {
		t, err := locale.Detect()
		if err != nil {
			slog.Error("detect locale error", "err", err.Error())
		} else {
			tag = t
		}
	}
	return tag
}
