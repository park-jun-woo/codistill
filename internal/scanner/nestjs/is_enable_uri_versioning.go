//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what call_expression이 enableVersioning(URI)인지 확인하고 defaultVersion을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// isEnableURIVersioning checks if a call_expression is
// app.enableVersioning({ type: VersioningType.URI, defaultVersion: '1' }).
// It returns (true, defaultVersion) when URI versioning is enabled; the
// defaultVersion is "" when absent or not a simple string literal.
func isEnableURIVersioning(call *sitter.Node, src []byte) (bool, string) {
	member := findChildByType(call, "member_expression")
	if member == nil {
		return false, ""
	}
	prop := findChildByType(member, "property_identifier")
	if prop == nil || nodeText(prop, src) != "enableVersioning" {
		return false, ""
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return false, ""
	}
	obj := findChildByType(args, "object")
	if obj == nil {
		// enableVersioning() with no args defaults to URI
		return true, ""
	}
	if !objectHasURIType(obj, src) {
		return false, ""
	}
	return true, extractURIDefaultVersion(obj, src)
}
