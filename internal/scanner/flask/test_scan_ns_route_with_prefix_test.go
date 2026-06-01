//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 @ns.route 클래스 라우트를 ns prefix와 합성하는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NSRoute_WithPrefix(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "plugin.py"), []byte(`from flask_restx import Resource

@inner_api_ns.route("/invoke/llm")
class PluginInvokeLLMApi(Resource):
    def post(self):
        return {}
`), 0o644)
	os.WriteFile(filepath.Join(dir, "app.py"), []byte(`from flask_restx import Api
api = Api()
api.add_namespace(inner_api_ns, path="/inner/api")
`), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, ep := range result.Endpoints {
		got[ep.Method+" "+ep.Path] = true
	}
	if !got["POST /inner/api/invoke/llm"] {
		t.Fatalf("missing POST /inner/api/invoke/llm; endpoints: %+v", got)
	}
}
