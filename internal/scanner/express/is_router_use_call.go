//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression이 router.use() 호출인지 확인한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func isRouterUseCall(call *sitter.Node, src []byte, routers map[string]bool) bool {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return false
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil || !routers[nodeText(obj, src)] {
		return false
	}
	prop := mem.ChildByFieldName("property")
	return prop != nil && isMountMethod(nodeText(prop, src))
}
