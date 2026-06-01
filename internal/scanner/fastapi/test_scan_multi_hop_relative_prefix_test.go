//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what Phase167: 다단계 상대 include_router(__init__ /api -> sub /app -> leaf /about) 누적(Mealie 패턴)
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

// TestScanMultiHopRelativePrefix reproduces the Mealie pattern: a top-level
// package __init__ defines APIRouter(prefix="/api") and includes a sub-package
// via a parenthesized relative import (`from . import (app,)`); the sub-package
// __init__ defines APIRouter(prefix="/app") and includes a leaf module; the leaf
// defines APIRouter(prefix="/about") with decorator routes. Every router is the
// generic name "router". The full chain must accumulate to /api/app/about.
func TestScanMultiHopRelativePrefix(t *testing.T) {
	root := t.TempDir()
	pkg := filepath.Join(root, "routes")
	app := filepath.Join(pkg, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}

	// routes/__init__.py: /api router, parenthesized relative include of sub-pkg.
	os.WriteFile(filepath.Join(pkg, "__init__.py"), []byte(
		"from fastapi import APIRouter\n"+
			"from . import (\n    app,\n)\n"+
			"router = APIRouter(prefix=\"/api\")\n"+
			"router.include_router(app.router)\n"), 0o644)

	// routes/app/__init__.py: /app router, relative include of leaf module.
	os.WriteFile(filepath.Join(app, "__init__.py"), []byte(
		"from fastapi import APIRouter\n"+
			"from . import app_about\n"+
			"router = APIRouter(prefix=\"/app\")\n"+
			"router.include_router(app_about.router)\n"), 0o644)

	// routes/app/app_about.py: leaf /about router with a decorator route.
	os.WriteFile(filepath.Join(app, "app_about.py"), []byte(
		"from fastapi import APIRouter\n"+
			"router = APIRouter(prefix=\"/about\")\n"+
			"@router.get(\"\")\n"+
			"def get_about():\n    return {}\n"), 0o644)

	// main.py: mount the top-level router on the app.
	os.WriteFile(filepath.Join(root, "main.py"), []byte(
		"from fastapi import FastAPI\n"+
			"from routes import router\n"+
			"app = FastAPI()\n"+
			"app.include_router(router)\n"), 0o644)

	paths := phase167PathSet(t, root)
	if !paths["/api/app/about"] {
		t.Fatalf("expected /api/app/about from multi-hop relative chain, got %v", paths)
	}
}
