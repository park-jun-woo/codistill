//ff:func feature=scan type=extract control=iteration dimension=2 topic=quarkus
//ff:what 클래스/인터페이스 노드의 상위 타입명(extends superclass + implements super_interfaces)을 수집한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func collectSupertypeNames(cls *sitter.Node, src []byte) []string {
	var names []string
	if superclass := findChildByType(cls, "superclass"); superclass != nil {
		if name := extractSuperclassName(superclass, src); name != "" {
			names = append(names, name)
		}
	}
	for _, key := range []string{"super_interfaces", "extends_interfaces"} {
		si := findChildByType(cls, key)
		if si == nil {
			continue
		}
		typeList := findChildByType(si, "type_list")
		if typeList == nil {
			continue
		}
		for i := 0; i < int(typeList.ChildCount()); i++ {
			child := typeList.Child(i)
			if child.Type() == "type_identifier" {
				names = append(names, nodeText(child, src))
			}
		}
	}
	return names
}
