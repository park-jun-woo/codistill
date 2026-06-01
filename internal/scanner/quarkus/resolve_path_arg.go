//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what @Path 어노테이션 인자(리터럴/상수/결합/정규식)를 해석하여 OpenAPI 경로로 정규화한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func resolvePathArg(ann *sitter.Node, src []byte, imports map[string]string, referrerPath, projectRoot string) string {
	args := annotationArgs(ann, src)
	if args == nil {
		return ""
	}
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		switch child.Type() {
		case "(", ")", ",":
			continue
		case "element_value_pair":
			continue
		}
		val := evalPathExpr(child, src, imports, referrerPath, projectRoot)
		if val != "" {
			return normalizePathRegex(val)
		}
	}
	return ""
}
