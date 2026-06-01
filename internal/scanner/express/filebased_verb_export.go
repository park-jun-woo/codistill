//ff:func feature=scan type=extract control=selection topic=express
//ff:what export_statement에서 HTTP VERB 식별자 export(const/function)를 식별해 메서드명과 핸들러 노드를 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// filebasedVerbExport inspects an export_statement node. If it exports a
// declaration whose identifier is an uppercase HTTP verb (GET/POST/PUT/PATCH/
// DELETE), it returns the uppercase method, the function/arrow node holding the
// handler body, and ok=true. Both `export const GET = (req,res) => {}` and
// `export async function GET(req,res) {}` forms are recognised. Non-verb
// exports (e.g. `export const config = {}`) return ok=false.
func filebasedVerbExport(exp *sitter.Node, src []byte) (method string, handler *sitter.Node, ok bool) {
	decl := exp.ChildByFieldName("declaration")
	if decl == nil {
		return "", nil, false
	}
	switch decl.Type() {
	case "lexical_declaration", "variable_declaration":
		vd := findChildByType(decl, "variable_declarator")
		if vd == nil {
			return "", nil, false
		}
		name := vd.ChildByFieldName("name")
		if name == nil || name.Type() != "identifier" {
			return "", nil, false
		}
		m, isVerb := filebasedVerbMethod(nodeText(name, src))
		if !isVerb {
			return "", nil, false
		}
		return m, vd.ChildByFieldName("value"), true
	case "function_declaration":
		name := decl.ChildByFieldName("name")
		if name == nil {
			return "", nil, false
		}
		m, isVerb := filebasedVerbMethod(nodeText(name, src))
		if !isVerb {
			return "", nil, false
		}
		return m, decl, true
	}
	return "", nil, false
}
