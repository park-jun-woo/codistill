//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRoutesFromCall_ArrayParamPath 배열형 + param 경로 전개 테스트
package express

import "testing"

func TestExtractRoutesFromCall_ArrayParamPath(t *testing.T) {
	fi := mustParse(t, []byte(`r.get(['/i', '/i/:id'], h);`))
	routes := extractRoutesFromCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true})
	if len(routes) != 2 || routes[0].Path != "/i" || routes[1].Path != "/i/:id" {
		t.Fatalf("got %+v", routes)
	}
}
