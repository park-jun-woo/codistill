//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what TestExtractRequireMethodDecorator_HTTPMethods — @require_http_methods(["POST","PUT"])가 두 메서드를 도출
package django

import "testing"

func TestExtractRequireMethodDecorator_HTTPMethods(t *testing.T) {
	src := `
@csrf_exempt
@require_http_methods(["POST", "PUT"])
def view(request):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	funcDef := djFirst(t, root, "function_definition")
	got := extractRequireMethodDecorator(funcDef, []byte(src))
	if len(got) != 2 {
		t.Fatalf("expected 2 methods, got %v", got)
	}
	want := map[string]bool{"POST": true, "PUT": true}
	for _, m := range got {
		if !want[m] {
			t.Errorf("unexpected method %q in %v", m, got)
		}
	}
}
