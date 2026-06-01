//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 ns.add_resource 등록을 ns prefix와 합성해 엔드포인트를 만드는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_NSAddResource_WithPrefix(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "res.py"), []byte(`from flask_restx import Resource

class TrialResource(Resource):
    def get(self):
        return {}
`), 0o644)
	os.WriteFile(filepath.Join(dir, "trial.py"), []byte(`from flask_restx import Api
api = Api()
api.add_namespace(console_ns, path="/console/api")
console_ns.add_resource(TrialResource, "/trial")
`), 0o644)

	result, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, ep := range result.Endpoints {
		got[ep.Method+" "+ep.Path] = true
	}
	if !got["GET /console/api/trial"] {
		t.Fatalf("missing GET /console/api/trial; endpoints: %+v", got)
	}
}
