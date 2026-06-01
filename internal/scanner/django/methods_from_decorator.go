//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 단일 decorator 노드에서 HTTP 메서드(@require_POST 류·@require_http_methods([...]))를 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// methodsFromDecorator returns the HTTP method(s) restricted by a single Django
// decorator node. It recognises bare shortcut decorators (@require_POST etc.)
// and the @require_http_methods([...]) call form. Unrelated/wrapper decorators
// yield nil.
func methodsFromDecorator(dec *sitter.Node, src []byte) []string {
	// Bare decorator: @require_POST
	if id := findChildByType(dec, "identifier"); id != nil {
		return requireMethodDecoratorNames[nodeText(id, src)]
	}
	// Call decorator: @require_http_methods([...])
	callNode := findChildByType(dec, "call")
	if callNode == nil {
		return nil
	}
	return httpMethodsCallList(callNode, src)
}
