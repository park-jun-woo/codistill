//ff:func feature=scan type=test control=sequence topic=flask
//ff:what parseExposeDecorator가 @expose(path, methods=...)를 파싱하고 비대상은 거르는지 검증한다
package flask

import (
	"reflect"
	"testing"
)

func TestParseExposeDecorator(t *testing.T) {
	dec, src := firstDecorator(t, "@expose(\"/login\", methods=[\"POST\"])\ndef f(): pass\n")
	path, methods, ok := parseExposeDecorator(dec, src, importAlias{})
	if !ok || path != "/login" || !reflect.DeepEqual(methods, []string{"POST"}) {
		t.Fatalf("expose list: got path=%q methods=%v ok=%v", path, methods, ok)
	}

	dec, src = firstDecorator(t, "@expose(\"/\", methods=(\"POST\",))\ndef f(): pass\n")
	path, methods, ok = parseExposeDecorator(dec, src, importAlias{})
	if !ok || path != "/" || !reflect.DeepEqual(methods, []string{"POST"}) {
		t.Fatalf("expose tuple: got path=%q methods=%v ok=%v", path, methods, ok)
	}

	dec, src = firstDecorator(t, "@expose(\"/items\")\ndef f(): pass\n")
	path, methods, ok = parseExposeDecorator(dec, src, importAlias{})
	if !ok || path != "/items" || !reflect.DeepEqual(methods, []string{"GET"}) {
		t.Fatalf("expose default GET: got path=%q methods=%v ok=%v", path, methods, ok)
	}

	// Non-expose decorator (attribute call) must be rejected.
	dec, src = firstDecorator(t, "@app.route(\"/x\")\ndef f(): pass\n")
	if _, _, ok := parseExposeDecorator(dec, src, importAlias{}); ok {
		t.Fatal("expected ok=false for @app.route")
	}

	// Aliased import: @ex resolves to expose.
	dec, src = firstDecorator(t, "@ex(\"/y\", methods=[\"GET\"])\ndef f(): pass\n")
	if _, _, ok := parseExposeDecorator(dec, src, importAlias{"ex": "expose"}); !ok {
		t.Fatal("expected ok=true for aliased expose")
	}
}
