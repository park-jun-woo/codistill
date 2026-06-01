//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Python 절대 import 경로를 파일 시스템 경로로 변환한다
package fastapi

import (
	"os"
	"path/filepath"
	"strings"
)

// resolveAbsoluteImportPath resolves an absolute Python import path to a file
// system path relative to absRoot. For example, "app.sneakers" resolves to
// "<absRoot>/app/sneakers/__init__.py" or "<absRoot>/app/sneakers.py".
// Returns "" if the module cannot be resolved.
func resolveAbsoluteImportPath(absRoot, module string) string {
	if module == "" || strings.HasPrefix(module, ".") {
		return ""
	}
	relPath := strings.ReplaceAll(module, ".", string(filepath.Separator))
	fullPath := filepath.Join(absRoot, relPath)

	pyFile := fullPath + ".py"
	if _, err := os.Stat(pyFile); err == nil {
		return pyFile
	}
	initFile := filepath.Join(fullPath, "__init__.py")
	if _, err := os.Stat(initFile); err == nil {
		return initFile
	}
	// 스캔 루트와 import(패키지) 루트가 다를 때(subpath 스캔) 조상 탐색으로 보강.
	return resolveAbsoluteImportAncestor(absRoot, module)
}
