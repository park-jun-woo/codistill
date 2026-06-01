//ff:func feature=scan type=convert control=sequence topic=flask
//ff:what 파일 상대경로를 그 파일이 속한 Python 패키지(디렉터리 모듈)로 변환한다
package flask

import "strings"

// packageOf returns the package (directory module) a file belongs to. A package
// __init__.py is its own package (controllers/trigger/__init__.py -> controllers.trigger);
// a regular module drops its final name component
// (controllers/trigger/webhook.py -> controllers.trigger).
func packageOf(relPath string) string {
	slash := filepathToSlash(relPath)
	if strings.HasSuffix(slash, "/__init__.py") || slash == "__init__.py" {
		return modulePathOf(relPath)
	}
	mod := modulePathOf(relPath)
	if i := strings.LastIndex(mod, "."); i >= 0 {
		return mod[:i]
	}
	return ""
}
