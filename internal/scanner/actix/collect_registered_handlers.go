//ff:func feature=scan type=extract control=iteration dimension=2 topic=actix
//ff:what 전 파일의 .service(<식별자>) bare-ident 인자명을 실등록 핸들러 집합으로 수집한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// collectRegisteredHandlers walks every file's AST and gathers the bare
// identifier arguments of all .service(<ident>) calls into a set. Macro routes
// (#[get]) are only emitted when their handler appears here, filtering out
// dead routes whose .service() registration is commented out (a line_comment
// is not a call_expression, so appendServiceCallHandlers skips it naturally).
func collectRegisteredHandlers(files []*fileInfo) map[string]bool {
	registered := make(map[string]bool)
	for _, fi := range files {
		walkNodes(fi.root, func(n *sitter.Node) {
			for _, h := range appendServiceCallHandlers(n, fi.src, nil) {
				registered[h] = true
			}
		})
	}
	return registered
}
