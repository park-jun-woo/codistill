//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what arguments 노드의 인자 중 __DIR__ . '/../Routes/X.php' 연결식을 Provider 파일 디렉터리 기준으로 정규화한 relPath를 반환한다
package laravel

import (
	"path"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// dirRelativeRouteFile scans the arguments of a call for a __DIR__-relative
// route file load of the form __DIR__ . '/../Routes/X.php' and resolves it
// against the provider file's directory (providerRelPath's dir), returning the
// normalized relPath that matches a parsedFiles key. Unlike base_path('X')
// (resolved from the project root), the base here is the provider's own
// directory, since __DIR__ is the directory of the file containing the call.
// It returns ok=false for any non-__DIR__ argument (variable/base_path/other),
// which the caller treats as unresolvable and skips.
func dirRelativeRouteFile(args *sitter.Node, providerRelPath string, src []byte) (string, bool) {
	for _, arg := range childrenOfType(args, "argument") {
		bin := findChildByType(arg, "binary_expression")
		if bin == nil {
			continue
		}
		name := findChildByType(bin, "name")
		if name == nil || nodeText(name, src) != "__DIR__" {
			continue
		}
		str := findChildByType(bin, "string")
		if str == nil {
			continue
		}
		suffix := extractStringContent(str, src)
		baseDir := path.Dir(strings.ReplaceAll(providerRelPath, "\\", "/"))
		return path.Clean(path.Join(baseDir, suffix)), true
	}
	return "", false
}
