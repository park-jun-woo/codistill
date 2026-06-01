//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Scan이 Flask-AppBuilder @expose 라우트를 base_url과 합성하는지 검증한다
package flask

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScan_Appbuilder_Expose(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "api.py"), []byte(`from flask_appbuilder.api import BaseApi, expose

class FooApi(BaseApi):
    @expose("/login", methods=["POST"])
    def login(self):
        return {}

    @expose("/refresh", methods=["POST"])
    def refresh(self):
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
	if !got["POST /api/v1/foo/login"] {
		t.Fatalf("missing POST /api/v1/foo/login; endpoints: %+v", got)
	}
	if !got["POST /api/v1/foo/refresh"] {
		t.Fatalf("missing POST /api/v1/foo/refresh; endpoints: %+v", got)
	}
}
