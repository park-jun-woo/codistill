//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestUnknownHelperIgnored 테스트
package django

import "testing"

func TestUnknownHelperIgnored(t *testing.T) {
	// A helper named something other than path/re_path/rest_path is not a path call.
	src := []byte("urlpatterns = [magic_path('users/', UserView.as_view())]")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := djFirst(t, root, "call")
	if entry := parsePathCall(call, src); entry != nil {
		t.Fatalf("expected nil for unknown helper, got %+v", entry)
	}

	// A function with no register call in its body is not promoted to a wrapper.
	fi := newTestFileInfo(t, "def helper(a, b):\n    return a + b\n\nhelper('x', Y)\n")
	wrappers := collectRegisterWrappersFromFile(fi)
	if wrappers["helper"] {
		t.Fatalf("non-register helper wrongly promoted: %v", wrappers)
	}
	if regs := extractWrapperRegisterCalls([]fileInfo{fi}, wrappers); len(regs) != 0 {
		t.Fatalf("expected no wrapper registrations, got %d", len(regs))
	}
}
