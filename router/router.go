package router

import (
	"regexp"
)

type Route struct {
	urlPattern string
	keys []string
	source string
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