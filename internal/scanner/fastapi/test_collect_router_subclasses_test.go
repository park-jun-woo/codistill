//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what Phase168: collectRouterSubclasses 직접/전이 상속 + 비라우터 제외
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestCollectRouterSubclasses(t *testing.T) {
	src := []byte(
		"from fastapi import APIRouter\n" +
			"class UserAPIRouter(APIRouter):\n    pass\n" +
			"class AdminRouter(UserAPIRouter):\n    pass\n" +
			"class Svc:\n    pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	subs := collectRouterSubclasses([]*sitter.Node{root}, [][]byte{src})
	if !subs["UserAPIRouter"] {
		t.Fatalf("expected UserAPIRouter recognized, got %v", subs)
	}
	if !subs["AdminRouter"] {
		t.Fatalf("expected transitive AdminRouter recognized, got %v", subs)
	}
	if subs["Svc"] {
		t.Fatalf("non-router class Svc must not be recognized, got %v", subs)
	}
}
