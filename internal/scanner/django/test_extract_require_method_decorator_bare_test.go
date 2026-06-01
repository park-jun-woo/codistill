//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractRequireMethodDecorator_Bare — @require_POST가 POST 메서드를 도출
package django

import "testing"

func TestExtractRequireMethodDecorator_Bare(t *testing.T) {
	src := `
@require_POST
def webhook(request):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	funcDef := djFirst(t, root, "function_definition")
	got := extractRequireMethodDecorator(funcDef, []byte(src))
	if len(got) != 1 || got[0] != "POST" {
		t.Fatalf("expected [POST], got %v", got)
	}
}
