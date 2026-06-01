//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveRouterPrefixes_Basic 테스트
package fastapi

import "testing"

func TestResolveRouterPrefixes_Basic(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, APIRouter

app = FastAPI()
router = APIRouter(prefix="/users")

app.include_router(router)
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	prefixes := resolveRouterPrefixes(root, src, nil)
	if prefixes["app"] != "" {
		t.Fatalf("expected empty prefix for app, got %q", prefixes["app"])
	}
	if prefixes["router"] != "/users" {
		t.Fatalf("expected /users for router, got %q", prefixes["router"])
	}
}
