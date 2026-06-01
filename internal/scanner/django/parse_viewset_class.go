//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 클래스 노드를 ViewSet으로 파싱한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// parseViewSetClass parses a class node as a ViewSet if it is one. The class
// index lets it resolve custom intermediate base classes transitively.
func parseViewSetClass(classNode *sitter.Node, fi fileInfo, idx classIndex) *viewsetInfo {
	nameNode := findChildByType(classNode, "identifier")
	if nameNode == nil {
		return nil
	}
	parents := extractParentClasses(classNode, fi.src)
	if !isViewSetSubclass(parents, idx) {
		return nil
	}
	vs := &viewsetInfo{
		name:    nodeText(nameNode, fi.src),
		parents: parents,
		methods: resolveMethodsTransitive(parents, idx),
		file:    fi.relPath,
		line:    int(nameNode.StartPoint().Row) + 1,
	}
	body := findChildByType(classNode, "block")
	if body != nil {
		vs.serializerClass = extractClassAttribute(body, "serializer_class", fi.src)
		vs.actions = extractActions(body, fi.src, fi.relPath)
	}
	return vs
}
