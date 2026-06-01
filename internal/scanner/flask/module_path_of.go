//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 파일 상대경로를 Python 모듈 경로로 변환한다
package flask

import "strings"

// modulePathOf converts a file's relative path into the Python module path other
// files would import it as. A package __init__.py maps to its directory's dotted
// name (controllers/trigger/__init__.py -> controllers.trigger); a regular module
// drops the .py suffix (controllers/trigger/webhook.py -> controllers.trigger.webhook).
func modulePathOf(relPath string) string {
	p := strings.TrimSuffix(filepathToSlash(relPath), ".py")
	p = strings.TrimPrefix(p, "./")
	p = strings.TrimPrefix(p, "/")
	p = strings.TrimSuffix(p, "/__init__")
	if p == "__init__" {
		return ""
	}
	return strings.ReplaceAll(p, "/", ".")
}
