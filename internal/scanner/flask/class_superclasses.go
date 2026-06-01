//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what class_definition의 상위클래스(식별자/attribute) 목록을 반환한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// classSuperclasses returns the superclass names of a class_definition as
// source-text strings (e.g. "Resource", "flask_restful.Resource", "ModelRestApi").
// Keyword arguments (metaclass=...) are skipped. The returned strings preserve
// dotted attribute paths so callers can resolve import aliases via
// collectImportAliases (Phase102 infra).
func classSuperclasses(classNode *sitter.Node, src []byte) []string {
	args := findChildByType(classNode, "argument_list")
	if args == nil {
		return nil
	}
	var supers []string
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		switch child.Type() {
		case "identifier", "attribute":
			supers = append(supers, nodeText(child, src))
		}
	}
	return supers
}
