//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what extractAliasImport + extractImports: aliased_import 수집과 origName 분리
package fastapi

import "testing"

func TestExtractAliasImport(t *testing.T) {
	src := []byte("from m import router as foo, files_router as files_v2\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	imports := extractImports(root, src)

	got := map[string]string{}
	for _, imp := range imports {
		got[imp.name] = imp.origName
	}
	if got["foo"] != "router" {
		t.Fatalf("foo origName = %q, want router", got["foo"])
	}
	if got["files_v2"] != "files_router" {
		t.Fatalf("files_v2 origName = %q, want files_router", got["files_v2"])
	}

	// 비별칭 import는 origName == name
	src2 := []byte("from m import APIRouter\n")
	root2, _ := parsePython(src2)
	for _, imp := range extractImports(root2, src2) {
		if imp.name == "APIRouter" && imp.origName != "APIRouter" {
			t.Fatalf("non-alias origName = %q, want APIRouter", imp.origName)
		}
	}
}
