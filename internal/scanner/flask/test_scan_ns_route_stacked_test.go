//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 한 클래스에 쌓인 @ns.route 두 개를 각각 path로 펼치는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NSRoute_StackedDecorators(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "x.py"), []byte(`from flask_restx import Resource

@ns.route("/a")
@ns.route("/b")
class X(Resource):
    def get(self):
        return {}
`), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, ep := range result.Endpoints {
		got[ep.Method+" "+ep.Path] = true
	}
	if !got["GET /a"] || !got["GET /b"] {
		t.Fatalf("expected GET /a and GET /b; endpoints: %+v", got)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d: %+v", len(result.Endpoints), got)
	}
}
