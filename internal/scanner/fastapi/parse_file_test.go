//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what parseFile 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "main.py")
	os.WriteFile(f, []byte("from .models import User\nx = 1\n"), 0o644)

	fi, err := parseFile(dir, f, nil)
	if err != nil {
		t.Fatal(err)
	}
	if fi.absPath != f {
		t.Fatalf("absPath: got %s", fi.absPath)
	}
	if fi.root == nil {
		t.Fatal("root is nil")
	}

	// non-existent file
	_, err = parseFile(dir, filepath.Join(dir, "nonexistent.py"), nil)
	if err == nil {
		t.Fatal("expected error for non-existent file")
	}
	// the parsePython error branch is unreachable: tree-sitter parses any bytes.
}
