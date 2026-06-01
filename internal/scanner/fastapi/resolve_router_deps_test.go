//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveRouterDeps 테스트
package fastapi

import "testing"

func TestResolveRouterDeps(t *testing.T) {
	src := []byte(`
from fastapi import FastAPI, APIRouter, Depends

app = FastAPI()
router = APIRouter(prefix="/sneakers", dependencies=[Depends(verify_token)])
public = APIRouter(prefix="/public")
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	deps := resolveRouterDeps(root, src, nil)

	// router should have verify_token
	if len(deps["router"]) != 1 {
		t.Fatalf("expected 1 dep for router, got %d", len(deps["router"]))
	}
	if deps["router"][0] != "verify_token" {
		t.Fatalf("expected verify_token, got %s", deps["router"][0])
	}

	// app has no dependencies
	if len(deps["app"]) != 0 {
		t.Fatalf("expected 0 deps for app, got %d", len(deps["app"]))
	}

	// public has no dependencies
	if len(deps["public"]) != 0 {
		t.Fatalf("expected 0 deps for public, got %d", len(deps["public"]))
	}
}
