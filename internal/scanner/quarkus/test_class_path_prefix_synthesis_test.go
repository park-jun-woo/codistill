//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestClassPathPrefixSynthesis -- 클래스 @Path 리터럴 프리픽스가 메서드 경로와 합성되는지 고정(Phase194)
package quarkus

import "testing"

func TestClassPathPrefixSynthesis(t *testing.T) {
	fi := qFileInfo(t, classPathPrefixSynthesisSource)
	resources := extractResources(fi)
	if len(resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(resources))
	}
	if resources[0].prefix != "/api" {
		t.Fatalf("class prefix: want /api, got %q", resources[0].prefix)
	}

	eps, _ := buildAllEndpoints(resources, "/abs")
	got := map[string]string{}
	for _, e := range eps {
		got[e.Handler] = e.Path
	}

	// 클래스 @Path("/api") + 메서드 @Path("/items") -> /api/items
	if got["listItems"] != "/api/items" {
		t.Errorf("listItems path: want /api/items, got %q", got["listItems"])
	}
	// 클래스 @Path("/api") + 메서드 @Path 없음 -> /api (메서드 단편 단독 붕괴 방지)
	if got["root"] != "/api" {
		t.Errorf("root path: want /api, got %q", got["root"])
	}
}
