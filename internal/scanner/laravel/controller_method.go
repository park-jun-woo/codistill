//ff:type feature=scan type=model topic=laravel
//ff:what 컨트롤러 메서드 정보(이름/파라미터/FormRequest 참조/return 노드)
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// controllerMethod holds information about a controller method.
type controllerMethod struct {
	name           string
	params         []methodParam
	formRequestRef string         // FormRequest type hint if any
	returnNodes    []*sitter.Node // return statement nodes
	methodNode     *sitter.Node   // method_declaration node for body-level scans
	src            []byte
	srcFI          *fileInfo // controller file (carries use-imports for FormRequest/Resource PSR-4 resolution)
}
