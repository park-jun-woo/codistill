//ff:func feature=scan type=extract control=sequence topic=express
//ff:what Express 경로의 :param을 OpenAPI {param}으로 변환한다
package express

import (
	"regexp"
)

var colonParamRe = regexp.MustCompile(`:([a-zA-Z_][a-zA-Z0-9_]*)(\([^)]*\))?\??`)

func expressPathToOpenAPI(path string) string {
	return colonParamRe.ReplaceAllString(path, "{$1}")
}
