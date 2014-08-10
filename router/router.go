package router

import (
	"regexp"
	"github.com/seanmcgary/minweb/server"
)

type Route struct {
	UrlPattern string
	Keys []string
	Source string
}

func (r Route) Match(url string)(bool){
	reg, _ := regexp.Compile(r.UrlPattern)

	matches := reg.FindAllStringSubmatch(url, -1)
	return len(matches) > 0
}

func CreateRoute(url string) (r Route){
	reg, _ := regexp.Compile(`\/`)

	//url := "/foo/:test/:balls"
	keys := make([]string, 0, 0)
	source := url

	url = reg.ReplaceAllString(url, "\\/")

	reg, _ = regexp.Compile(`\.`)
	url = reg.ReplaceAllString(url, `\\.?`)

	reg, _ = regexp.Compile(`\*`)
	url = reg.ReplaceAllString(url, `.+`)

	reg, _ = regexp.Compile(`:(\w+)(?:\(([^\)]+)\))?(\?)?`)

	url = reg.ReplaceAllStringFunc(url, func(str string) string{
		keys = append(keys, str[1:])
		return `([^\/]+)`
	})

	reg, _ = regexp.Compile(`\\\/\(\[\^\/\]\*\)`)
	url = reg.ReplaceAllString(url, `(?:\\/(\\w*))?`)

	url = "^" + url + `\/?$`

	return Route{url, keys, source}
}