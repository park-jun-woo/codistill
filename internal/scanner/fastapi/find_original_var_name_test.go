//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findOriginalVarName: prefix에 존재 / 미존재 분기
package fastapi

import "testing"

func TestFindOriginalVarName(t *testing.T) {
	fi := &fileInfo{prefixes: map[string]string{"router": "/api"}}
	if got := findOriginalVarName("router", nil, fi); got != "router" {
		t.Fatalf("got %q", got)
	}
	if got := findOriginalVarName("missing", nil, fi); got != "" {
		t.Fatalf("got %q", got)
	}
	// alias: from m import router as foo -> orig "router"
	refImports := []importInfo{{name: "foo", module: "m", origName: "router"}}
	if got := findOriginalVarName("foo", refImports, fi); got != "router" {
		t.Fatalf("alias lookup got %q, want router", got)
	}
}
