//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what array 노드의 직접 자식 string 리터럴들을 unquote하여 path 슬라이스로 수집한다(비-string 요소는 스킵)
package express

import sitter "github.com/smacker/go-tree-sitter"

func collectArrayStringPaths(arr *sitter.Node, src []byte) []string {
	var paths []string
	for _, s := range childrenOfType(arr, "string") {
		paths = append(paths, unquoteTS(nodeText(s, src)))
	}
	return paths
}
