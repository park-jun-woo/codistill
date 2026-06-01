//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what Phase167: 3-hop(/api->/v1->leaf) + alias import 누적(LangFlow 패턴)
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanThreeHopAliasPrefix(t *testing.T) {
	root := t.TempDir()
	os.WriteFile(filepath.Join(root, "voice_mode.py"), []byte(
		"from fastapi import APIRouter\n"+
			"router = APIRouter()\n"+
			"@router.get('/listen')\n"+
			"def listen():\n    return {}\n"), 0o644)
	os.WriteFile(filepath.Join(root, "api_router.py"), []byte(
		"from fastapi import APIRouter\n"+
			"from .voice_mode import router as voice_router\n"+
			"router_v1 = APIRouter(prefix='/v1')\n"+
			"router_v1.include_router(voice_router)\n"+
			"router = APIRouter(prefix='/api')\n"+
			"router.include_router(router_v1)\n"), 0o644)
	os.WriteFile(filepath.Join(root, "main.py"), []byte(
		"from fastapi import FastAPI\n"+
			"from .api_router import router\n"+
			"app = FastAPI()\n"+
			"app.include_router(router)\n"), 0o644)

	paths := phase167PathSet(t, root)
	if !paths["/api/v1/listen"] {
		t.Fatalf("expected /api/v1/listen from 3-hop alias chain, got %v", paths)
	}
}
