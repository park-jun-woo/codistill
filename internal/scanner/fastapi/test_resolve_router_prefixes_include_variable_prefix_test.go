//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveRouterPrefixes_IncludeWithVariablePrefix 테스트
package fastapi

import "testing"

func TestResolveRouterPrefixes_IncludeWithVariablePrefix(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, APIRouter

API_PREFIX = "/api"
app = FastAPI()
router = APIRouter(prefix="/v1")

app.include_router(router, prefix=API_PREFIX)
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
