//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 ns prefix 미발견 시 route path만으로 엔드포인트를 만드는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NSRoute_NoPrefix(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "plugin.py"), []byte(`from flask_restx import Resource

@my_ns.route("/invoke/llm")
class X(Resource):
    def post(self):
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
	if !got["POST /invoke/llm"] {
		t.Fatalf("missing POST /invoke/llm; endpoints: %+v", got)
	}
}
