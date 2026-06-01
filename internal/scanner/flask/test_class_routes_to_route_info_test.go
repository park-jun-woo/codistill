//ff:func feature=scan type=test control=sequence topic=flask
//ff:what classRoutesToRouteInfo가 prefix+path 합성과 param 추출을 정확히 수행하는지 검증한다
package flask

import "testing"

func TestClassRoutesToRouteInfo_PathAndParams(t *testing.T) {
	methods := []classMethod{
		{name: "GET", line: 3},
		{name: "POST", line: 6},
	}
	routes := classRoutesToRouteInfo(methods, "/<int:user_id>", "/api/users", "UserAPI.", "api.py")
	if len(routes) != 2 {
		t.Fatalf("expected 2 routes, got %d", len(routes))
	}
	if routes[0].path != "/api/users/{user_id}" {
		t.Errorf("expected /api/users/{user_id}, got %s", routes[0].path)
	}
	if routes[0].method != "GET" {
		t.Errorf("expected GET, got %s", routes[0].method)
	}
	if routes[0].handler != "UserAPI.GET" {
		t.Errorf("expected UserAPI.GET, got %s", routes[0].handler)
	}
	if routes[0].line != 3 {
		t.Errorf("expected line 3, got %d", routes[0].line)
	}
	if len(routes[0].params) != 1 || routes[0].params[0].name != "user_id" || routes[0].params[0].converter != "int" {
		t.Errorf("expected param user_id:int, got %+v", routes[0].params)
	}
	if routes[1].method != "POST" || routes[1].handler != "UserAPI.POST" {
		t.Errorf("expected POST UserAPI.POST, got %s %s", routes[1].method, routes[1].handler)
	}
}
