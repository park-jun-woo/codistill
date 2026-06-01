//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what 한 파일 AST에서 add_namespace 호출의 ns 변수→prefix 매핑을 추출한다
package flask

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

// fileNamespacePrefixes scans one file AST for *.add_namespace(ns[, "/x"][, path="/x"])
// calls and returns a namespace-variable -> prefix map. The first positional
// identifier argument is the namespace variable; the prefix comes from a positional
// string argument when present, otherwise from the path= keyword. A namespace
// registered without any path maps to an empty prefix.
func fileNamespacePrefixes(root *sitter.Node, src []byte) namespacePrefix {
	prefixes := make(namespacePrefix)
	for _, call := range findAllByType(root, "call") {
		attr := findChildByType(call, "attribute")
		if attr == nil {
			continue
		}
		if !strings.HasSuffix(nodeText(attr, src), ".add_namespace") {
			continue
		}
		args := findChildByType(call, "argument_list")
		if args == nil {
			continue
		}
		nsVar := firstIdentArg(args, src)
		if nsVar == "" {
			continue
		}
		prefix := firstStringArg(args, src)
		if prefix == "" {
			prefix = extractKeywordArg(args, "path", src)
		}
		prefixes[nsVar] = prefix
	}
	return prefixes
}
