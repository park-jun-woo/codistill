//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveAbsoluteImportAncestor: subpath 스캔 시 패키지 루트 조상 탐색 해석
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveAbsoluteImportAncestor(t *testing.T) {
	root := t.TempDir()
	// /root/backend/open_webui/routers/__init__.py  (패키지 루트 = backend)
	pkgRouters := filepath.Join(root, "backend", "open_webui", "routers")
	os.MkdirAll(pkgRouters, 0o755)
	initFile := filepath.Join(pkgRouters, "__init__.py")
	os.WriteFile(initFile, []byte(""), 0o644)

	// 스캔 루트는 backend/open_webui (subpath), import는 open_webui.routers
	scanRoot := filepath.Join(root, "backend", "open_webui")

	// 직접 매칭(기존)은 실패해야 한다(open_webui/open_webui/routers 없음)
	if got := resolveAbsoluteImportPath(scanRoot, "open_webui.routers"); got != initFile {
		t.Fatalf("ancestor resolution = %q, want %q", got, initFile)
	}

	// 회귀: 루트 직접 매칭 케이스
	os.MkdirAll(filepath.Join(scanRoot, "app"), 0o755)
	appPy := filepath.Join(scanRoot, "app", "x.py")
	os.WriteFile(appPy, []byte(""), 0o644)
	if got := resolveAbsoluteImportPath(scanRoot, "app.x"); got != appPy {
		t.Fatalf("direct resolution = %q, want %q", got, appPy)
	}

	// 미존재 모듈은 빈 문자열
	if got := resolveAbsoluteImportPath(scanRoot, "does.not.exist"); got != "" {
		t.Fatalf("unresolvable = %q, want empty", got)
	}
}
