//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 객체 리터럴에서 defaultVersion 문자열 리터럴 값을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// extractURIDefaultVersion returns the defaultVersion value from
// enableVersioning({ type: VersioningType.URI, defaultVersion: '1' }).
// Only simple string literals are returned; VERSION_NEUTRAL, arrays, or other
// non-literal forms yield "" (conservative — no version prefix applied).
func extractURIDefaultVersion(obj *sitter.Node, src []byte) string {
	for i := 0; i < int(obj.ChildCount()); i++ {
		child := obj.Child(i)
		if child.Type() != "pair" {
			continue
		}
		key := findChildByType(child, "property_identifier")
		if key == nil || nodeText(key, src) != "defaultVersion" {
			continue
		}
		return findPairStringValue(child, src)
	}
	return ""
}
