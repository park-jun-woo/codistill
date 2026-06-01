//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractRequireMethodDecorator_None — 메서드 데코레이터 없는 plain 뷰는 nil(=GET 폴백 유지)
package django

import "testing"

func TestExtractRequireMethodDecorator_None(t *testing.T) {
	src := `
@login_required
def view(request):
    pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	funcDef := djFirst(t, root, "function_definition")
	if got := extractRequireMethodDecorator(funcDef, []byte(src)); got != nil {
		t.Fatalf("expected nil (GET fallback), got %v", got)
	}
}
