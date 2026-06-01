//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveRouterPrefixes_IncludeWithPrefix 테스트
package fastapi

import "testing"

func TestResolveRouterPrefixes_IncludeWithPrefix(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, APIRouter

app = FastAPI()
router = APIRouter(prefix="/v1")

app.include_router(router, prefix="/api")
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	if prefixes["router"] != "/api/v1" {
		t.Fatalf("expected /api/v1, got %q", prefixes["router"])
	}
}
