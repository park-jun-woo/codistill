//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what require/require_once 식의 인자(bare 'x.php' | __DIR__ . '/x.php' | base_path('routes/x.php'))를 요청 파일 디렉터리 기준 relPath로 해석한다
package laravel

import (
	"path"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// requireIncludePath resolves the route file loaded by a require_expression or
// require_once_expression. Bagisto-style modular route files load sibling files
// with a bare relative literal — require 'auth-routes.php' — which PHP resolves
// against the requiring file's own directory (include_path), so the literal is
// joined onto callerRelPath's dir. __DIR__ . '/x.php' concatenations resolve the
// same way, and base_path('routes/x.php') resolves from the project root. It
// returns ok=false for variable/other arguments, which the caller skips.
func requireIncludePath(node *sitter.Node, callerRelPath string, src []byte) (string, bool) {
	if call := findChildByType(node, "function_call_expression"); call != nil {
		if rel, ok := extractBasePathArg(call, src); ok {
			return rel, true
		}
	}
	baseDir := path.Dir(strings.ReplaceAll(callerRelPath, "\\", "/"))
	if bin := findChildByType(node, "binary_expression"); bin != nil {
		name := findChildByType(bin, "name")
		str := findChildByType(bin, "string")
		if name != nil && nodeText(name, src) == "__DIR__" && str != nil {
			return path.Clean(path.Join(baseDir, extractStringContent(str, src))), true
		}
		return "", false
	}
	if str := findChildByType(node, "string"); str != nil {
		return path.Clean(path.Join(baseDir, extractStringContent(str, src))), true
	}
	return "", false
}
