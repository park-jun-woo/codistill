//ff:func feature=scan type=test control=sequence topic=flask
//ff:what mustParse 테스트 헬퍼 — Python 소스를 파싱해 루트 노드와 바이트를 반환한다
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func mustParse(t *testing.T, code string) (*sitter.Node, []byte) {
	t.Helper()
	src := []byte(code)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	return root, src
}
