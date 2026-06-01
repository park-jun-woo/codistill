//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveRouterPrefixes_VariableRef 테스트
package fastapi

import "testing"

func TestResolveRouterPrefixes_VariableRef(t *testing.T) {
	src := []byte(`
from fastapi import APIRouter

API_STR = "/api"
sneakers_router = APIRouter(prefix=API_STR)
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	if prefixes["sneakers_router"] != "/api" {
		t.Fatalf("expected /api for sneakers_router, got %q", prefixes["sneakers_router"])
	}
}
