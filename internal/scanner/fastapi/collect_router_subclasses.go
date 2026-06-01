//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 전 파일의 class_definition에서 APIRouter(직·간접) 상속 클래스명을 고정점 수집한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// collectRouterSubclasses scans class definitions across all parsed files and
// returns the set of class names that (directly or transitively) inherit from
// APIRouter. It iterates to a fixed point so that chains like
// class A(APIRouter) / class B(A) are all recognized.
//
// roots[i] is the AST root for srcs[i]; entries with a nil root are skipped.
func collectRouterSubclasses(roots []*sitter.Node, srcs [][]byte) map[string]bool {
	subclasses := make(map[string]bool)
	for changed := true; changed; {
		changed = collectRouterSubclassesSweep(roots, srcs, subclasses)
	}
	return subclasses
}
