//ff:func feature=scan type=test control=sequence topic=flask
//ff:what collectConfigureAPITuples가 (path, Resource) 튜플 리스트를 인식하는지 검증한다
package flask

import "testing"

func TestCollectConfigureAPITuples_Recognizes(t *testing.T) {
	src := []byte(`configure_api_from_blueprint(bp, [("/x", FooResource), ("/y", BarResource)])
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	regs := collectConfigureAPITuples(root, src)
	if len(regs) != 2 {
		t.Fatalf("expected 2 regs, got %d: %+v", len(regs), regs)
	}
	if regs[0].className != "FooResource" || regs[0].paths[0] != "/x" || regs[0].blueprintVar != "bp" {
		t.Errorf("reg 0 wrong: %+v", regs[0])
	}
	if regs[1].className != "BarResource" || regs[1].paths[0] != "/y" || regs[1].blueprintVar != "bp" {
		t.Errorf("reg 1 wrong: %+v", regs[1])
	}
}
