//ff:func feature=scan type=test control=sequence topic=flask
//ff:what collectAddResource가 클래스/단일·다중 path 인자를 추출하는지 검증한다
package flask

import "testing"

func TestCollectAddResource_SingleAndMulti(t *testing.T) {
	src := []byte(`api.add_resource(EventsResource, "/data/events/last")
api.add_resource(UserResource, "/users", "/users/<int:id>")
notcall.other(Foo, "/x")
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	regs := collectAddResource(root, src)
	if len(regs) != 2 {
		t.Fatalf("expected 2 regs, got %d: %+v", len(regs), regs)
	}
	if regs[0].className != "EventsResource" || len(regs[0].paths) != 1 || regs[0].paths[0] != "/data/events/last" {
		t.Errorf("reg 0 wrong: %+v", regs[0])
	}
	if regs[1].className != "UserResource" || len(regs[1].paths) != 2 {
		t.Fatalf("reg 1 wrong: %+v", regs[1])
	}
	if regs[1].paths[0] != "/users" || regs[1].paths[1] != "/users/<int:id>" {
		t.Errorf("reg 1 paths wrong: %+v", regs[1].paths)
	}
}
