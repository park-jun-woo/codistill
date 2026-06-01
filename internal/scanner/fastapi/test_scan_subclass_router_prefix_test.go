//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what Phase168: 커스텀 APIRouter 서브클래스(분리 파일)의 prefix 인식 + cross-file 합성(Mealie 패턴)
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanSubclassRouterPrefix(t *testing.T) {
	root := t.TempDir()
	base := filepath.Join(root, "routers.py")
	os.WriteFile(base, []byte(
		"from fastapi import APIRouter\n"+
			"class UserAPIRouter(APIRouter):\n    pass\n"), 0o644)
	os.WriteFile(filepath.Join(root, "api_tokens.py"), []byte(
		"from .routers import UserAPIRouter\n"+
			"router = UserAPIRouter(prefix='/users')\n"+
			"@router.post('/api-tokens')\n"+
			"def create():\n    return {}\n"), 0o644)
	os.WriteFile(filepath.Join(root, "main.py"), []byte(
		"from fastapi import FastAPI\n"+
			"from .api_tokens import router\n"+
			"app = FastAPI()\n"+
			"app.include_router(router, prefix='/api')\n"), 0o644)

	paths := phase167PathSet(t, root)
	if !paths["/api/users/api-tokens"] {
		t.Fatalf("expected /api/users/api-tokens from subclass router prefix chain, got %v", paths)
	}
}
