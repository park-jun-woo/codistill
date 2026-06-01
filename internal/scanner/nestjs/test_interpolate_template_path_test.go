//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what interpolateTemplatePath의 ${const} 보간·미해석조각 원형유지 테스트
package nestjs

import "testing"

func TestInterpolateTemplatePath(t *testing.T) {
	src := []byte(`const PREFIX_APIV3_DATA = '/api/v3/data/:baseName';`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	pc := enumPathCtx{root: root, src: src, absFile: "f.ts"}

	// resolvable ${CONST}
	got, ok := pc.interpolateTemplatePath("${PREFIX_APIV3_DATA}/:modelId/records")
	if !ok {
		t.Fatalf("expected interpolation to fire")
	}
	if got != "/api/v3/data/:baseName/:modelId/records" {
		t.Errorf("got %q", got)
	}

	// no template fragment -> not fired
	if v, ok := pc.interpolateTemplatePath("/plain/:id"); ok || v != "/plain/:id" {
		t.Errorf("unexpected fire: %q %v", v, ok)
	}

	// unresolvable fragment (function call) -> kept verbatim
	got2, ok2 := pc.interpolateTemplatePath("${fn()}/x")
	if !ok2 {
		t.Fatalf("expected fire (fragment present)")
	}
	if got2 != "${fn()}/x" {
		t.Errorf("expected verbatim, got %q", got2)
	}
}
