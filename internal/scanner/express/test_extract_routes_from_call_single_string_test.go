//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractRoutesFromCall_SingleStringRegression 단일 string 경로 회귀 테스트
package express

import "testing"

func TestExtractRoutesFromCall_SingleStringRegression(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', h);`))
	routes := extractRoutesFromCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true})
	if len(routes) != 1 || routes[0].Path != "/x" || routes[0].Method != "GET" || routes[0].Router != "r" {
		t.Fatalf("got %+v", routes)
	}
}
