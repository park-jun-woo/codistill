//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 함수 정의를 @api_view 또는 @require_* 메서드 데코레이터 함수 뷰로 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseFuncView parses a function definition into a funcViewInfo when its
// decorators declare HTTP methods. DRF's @api_view(["GET", ...]) takes priority;
// otherwise plain Django views restricted by @require_POST / @require_GET /
// @require_http_methods([...]) are captured so they are emitted with their real
// methods instead of the GET-only plain-view fallback. Returns nil when no
// HTTP-method-declaring decorator is present.
func parseFuncView(funcDef *sitter.Node, fi fileInfo) *funcViewInfo {
	nameNode := findChildByType(funcDef, "identifier")
	if nameNode == nil {
		return nil
	}
	methods := extractAPIViewDecorator(funcDef, fi.src)
	if methods == nil {
		methods = extractRequireMethodDecorator(funcDef, fi.src)
	}
	if methods == nil {
		return nil
	}
	return &funcViewInfo{
		name:    nodeText(nameNode, fi.src),
		methods: methods,
		file:    fi.relPath,
		line:    int(nameNode.StartPoint().Row) + 1,
	}
}
