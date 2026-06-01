//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what Phase167: subpath 스캔 루트 + 키워드 prefix include_router 누적(OpenWebUI 패턴)
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanKeywordPrefixAncestor(t *testing.T) {
	root := t.TempDir()
	pkg := filepath.Join(root, "backend", "open_webui")
	routers := filepath.Join(pkg, "routers")
	os.MkdirAll(routers, 0o755)

	os.WriteFile(filepath.Join(routers, "__init__.py"), []byte(""), 0o644)
	os.WriteFile(filepath.Join(routers, "users.py"), []byte(
		"from fastapi import APIRouter\n"+
			"router = APIRouter()\n"+
			"@router.get('/')\n"+
			"def list_users():\n    return []\n"), 0o644)
	os.WriteFile(filepath.Join(pkg, "main.py"), []byte(
		"from fastapi import FastAPI\n"+
			"from open_webui.routers import users\n"+
			"app = FastAPI()\n"+
			"app.include_router(users.router, prefix='/api/v1/users')\n"), 0o644)

	// 스캔 루트 = subpath backend/open_webui (import 루트 = backend)
	paths := phase167PathSet(t, pkg)
	if !paths["/api/v1/users/"] && !paths["/api/v1/users"] {
		t.Fatalf("expected /api/v1/users prefix on leaf route, got %v", paths)
	}
}
