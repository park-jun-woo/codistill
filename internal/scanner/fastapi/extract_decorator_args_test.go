//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractDecoratorArgs 테스트
package fastapi

import "testing"

func TestExtractDecoratorArgs(t *testing.T) {
	src := []byte("@router.post('/users', status_code=201, response_model=UserOut)\ndef f(): pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	decs := findAllByType(root, "decorator")
	if len(decs) == 0 {
		t.Fatal("no decorator")
	}
	callNode := findChildByType(decs[0], "call")
	da := extractDecoratorArgs(callNode, src)
	if da.path != "/users" {
		t.Fatalf("path: got %q", da.path)
	}
	if da.statusCode != 201 {
		t.Fatalf("status: got %d", da.statusCode)
	}
	if da.responseModel != "UserOut" {
		t.Fatalf("respModel: got %q", da.responseModel)
	}
	if da.responseClass != "" {
		t.Fatalf("respClass: got %q", da.responseClass)
	}
	if !da.includeInSchema {
		t.Fatal("includeInSchema: expected true by default")
	}

	// nil callNode
	da0 := extractDecoratorArgs(nil, src)
	if da0.path != "" || da0.statusCode != 0 || da0.responseModel != "" || da0.responseClass != "" {
		t.Fatal("expected empty for nil")
	}
	if !da0.includeInSchema {
		t.Fatal("includeInSchema: expected true for nil")
	}

	// decorator without argument_list (bare attribute access)
	src2 := []byte("@router.get\ndef f(): pass\n")
	root2, err := parsePython(src2)
	if err != nil {
		t.Fatal(err)
	}
	decs2 := findAllByType(root2, "decorator")
	if len(decs2) == 0 {
		t.Fatal("no decorator")
	}
	// The decorator's child is an attribute, not a call node; pass it directly
	attr := findChildByType(decs2[0], "attribute")
	da2 := extractDecoratorArgs(attr, src2)
	if da2.path != "" || da2.statusCode != 0 || da2.responseModel != "" || da2.responseClass != "" {
		t.Fatal("expected empty for non-call decorator")
	}
}
