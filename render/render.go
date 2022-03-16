package render

import (
	"embed"
	"fmt"
	"html/template"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

var Render multitemplate.Render
var TemplateBox embed.FS

func InitRender() {
	Render = render(TemplateBox)
}

func BuildTemplate(template_name string, templates []string) (*template.Template, error) {
	temp, err := template.New(template_name).Parse(
		strings.Join(templates, " "))
	return temp, err
}

func getContent(f embed.FS, filename string) string {

	fmt.Println(filename)
	data, err := f.ReadFile(filename)
	if err != nil {
		//glog.Error(err)
		fmt.Println(err)
		return ""
	} else {
		return string(data)
	}
}

func homeRender(f embed.FS, r *multitemplate.Render) error {
	home_templates := []string{
		getContent(f, "templates/base/base.html"),
		getContent(f, "templates/base/search.html"),
		getContent(f, "templates/base/navi.html"),
		getContent(f, "templates/base/footer.html"),
		getContent(f, "templates/base/track.html"),
		getContent(f, "templates/home/home.html"),
	}

	home_tmp, err := BuildTemplate("Home", home_templates)
	if err != nil {
		fmt.Println(err)
		return err
	}

	r.Add("Home", home_tmp)
	return nil
}

func errorRender(f embed.FS, r *multitemplate.Render) error {
	error_templates := []string{
		getContent(f, "templates/base/base.html"),
		getContent(f, "templates/base/search.html"),
		getContent(f, "templates/base/navi.html"),
		getContent(f, "templates/base/footer.html"),
		getContent(f, "templates/base/track.html"),
		getContent(f, "templates/error/error.html"),
	}
	error_tmp, err := BuildTemplate("error", error_templates)
	if err != nil {
		return err
	}
	r.Add("Error", error_tmp)
	return nil
}

func render(f embed.FS) multitemplate.Render {

	r := multitemplate.New()

	homeRender(f, &r)

	return r
}
