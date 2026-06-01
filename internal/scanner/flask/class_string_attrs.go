//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 클래스 본문의 `name = "string"` 대입을 name→값 맵으로 수집한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// classStringAttrs collects top-level `name = "literal"` assignments in a class
// body into a name->value map (unquoted). Only string right-hand sides are kept,
// so Flask-AppBuilder declarations like base_url / route_base / resource_name are
// captured while datamodel = SQLAInterface(...) is ignored. Used by classBaseURL.
func classStringAttrs(classNode *sitter.Node, src []byte) map[string]string {
	attrs := make(map[string]string)
	body := findChildByType(classNode, "block")
	if body == nil {
		return attrs
	}
	for _, es := range childrenOfType(body, "expression_statement") {
		left, val, ok := stringAssignment(es, src)
		if ok {
			attrs[left] = val
		}
	}
	return attrs
}
