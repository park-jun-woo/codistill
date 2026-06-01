//ff:func feature=scan type=test control=sequence topic=flask
//ff:what collectAddURLRule가 path·핸들러·메서드·앱루트(!) 플래그를 추출하는지 검증한다
package flask

import "testing"

func TestCollectAddURLRule_Basic(t *testing.T) {
	src := []byte(`_bp.add_url_rule('!/admin/admins', 'admins', RHAdmins, methods=('GET', 'POST'))
_bp.add_url_rule('/users/<int:id>', 'user', RHUser)
notcall.other(Foo, '/x')
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	regs := collectAddURLRule(root, src, "blueprint.py")
	if len(regs) != 2 {
		t.Fatalf("expected 2 regs, got %d: %+v", len(regs), regs)
	}
	r0 := regs[0]
	if r0.rawPath != "/admin/admins" || !r0.appRoot {
		t.Errorf("reg 0 path/appRoot wrong: %+v", r0)
	}
	if r0.handler != "RHAdmins" || r0.blueprintVar != "_bp" || r0.file != "blueprint.py" {
		t.Errorf("reg 0 handler/base/file wrong: %+v", r0)
	}
	if len(r0.methods) != 2 || r0.methods[0] != "GET" || r0.methods[1] != "POST" {
		t.Errorf("reg 0 methods wrong: %+v", r0.methods)
	}
	r1 := regs[1]
	if r1.rawPath != "/users/<int:id>" || r1.appRoot {
		t.Errorf("reg 1 path/appRoot wrong: %+v", r1)
	}
	if len(r1.methods) != 0 {
		t.Errorf("reg 1 expected no methods, got %+v", r1.methods)
	}
}
