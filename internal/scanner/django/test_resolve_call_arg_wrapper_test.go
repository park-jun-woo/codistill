//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestResolveCallArg_ViewWrapperUnwrapping 테스트 (Phase165)
package django

import "testing"

func TestResolveCallArg_ViewWrapperUnwrapping(t *testing.T) {
	// Single wrapper: staff_member_required(my_view) -> my_view
	src := []byte("path('x/', staff_member_required(my_view))\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := djFirst(t, root, "argument_list")
	pos := positionalArgs(args)
	var e urlEntry
	resolveCallArg(&e, pos[1], src)
	if e.viewName != "my_view" {
		t.Fatalf("single wrapper viewName = %q, want my_view", e.viewName)
	}

	// Nested wrappers: login_required(staff_member_required(v)) -> v
	src2 := []byte("path('y/', login_required(staff_member_required(v)))\n")
	root2, _ := parsePython(src2)
	args2 := djFirst(t, root2, "argument_list")
	pos2 := positionalArgs(args2)
	var e2 urlEntry
	resolveCallArg(&e2, pos2[1], src2)
	if e2.viewName != "v" {
		t.Fatalf("nested wrapper viewName = %q, want v", e2.viewName)
	}

	// Wrapper around as_view dict: csrf_exempt(MyView.as_view({"get": "list"}))
	// must preserve viewName and methodActions.
	src3 := []byte("path('z/', csrf_exempt(MyView.as_view({\"get\": \"list\"})))\n")
	root3, _ := parsePython(src3)
	args3 := djFirst(t, root3, "argument_list")
	pos3 := positionalArgs(args3)
	var e3 urlEntry
	resolveCallArg(&e3, pos3[1], src3)
	if e3.viewName != "MyView" {
		t.Fatalf("as_view wrapper viewName = %q, want MyView", e3.viewName)
	}
	if e3.methodActions["get"] != "list" {
		t.Fatalf("as_view wrapper methodActions = %+v, want get->list", e3.methodActions)
	}

	// require_* prefix family: require_GET(listing) -> listing
	src4 := []byte("path('g/', require_GET(listing))\n")
	root4, _ := parsePython(src4)
	args4 := djFirst(t, root4, "argument_list")
	pos4 := positionalArgs(args4)
	var e4 urlEntry
	resolveCallArg(&e4, pos4[1], src4)
	if e4.viewName != "listing" {
		t.Fatalf("require_* wrapper viewName = %q, want listing", e4.viewName)
	}

	// Unknown call: preserve legacy behavior (wrapper name as view name).
	src5 := []byte("path('u/', some_unknown_helper(inner_view))\n")
	root5, _ := parsePython(src5)
	args5 := djFirst(t, root5, "argument_list")
	pos5 := positionalArgs(args5)
	var e5 urlEntry
	resolveCallArg(&e5, pos5[1], src5)
	if e5.viewName != "some_unknown_helper" {
		t.Fatalf("unknown call viewName = %q, want some_unknown_helper (legacy)", e5.viewName)
	}
}
