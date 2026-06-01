//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRoutesFromCall_ArrayPath 배열형 경로 전개 테스트
package express

import "testing"

func TestExtractRoutesFromCall_ArrayPath(t *testing.T) {
	fi := mustParse(t, []byte(`r.get(['/a', '/a/*'], h);`))
	routes := extractRoutesFromCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true})
	if len(routes) != 2 {
		t.Fatalf("want 2 routes, got %d: %+v", len(routes), routes)
	}
	if routes[0].Path != "/a" || routes[1].Path != "/a/*" {
		t.Fatalf("paths: %q %q", routes[0].Path, routes[1].Path)
	}
	if routes[0].Method != "GET" || routes[0].Router != "r" || routes[1].Method != "GET" || routes[1].Router != "r" {
		t.Fatalf("route meta: %+v", routes)
	}
}
